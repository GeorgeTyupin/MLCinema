from flask import Flask, request, jsonify
import numpy as np
import pickle
from movie_utils import preprocess_text, extract_named_entities, encode_text
from sklearn.metrics.pairwise import cosine_similarity
from natasha import Doc, NewsNERTagger, NewsEmbedding, Segmenter, MorphVocab
import json
import requests
import os
import subprocess
import sys

app = Flask(__name__)

# Загружаем данные при старте
movies = []
movie_embeddings = []

def load_movies_from_go():
    """Загружает фильмы из Go API и сохраняет в movies.json"""
    try:
        print("Получаем фильмы из Go сервиса...")
        
        # Запрос к Go API
        response = requests.post('http://localhost:8000/api/get-films', timeout=10)
        
        if response.status_code == 200:
            go_movies = response.json()
            print(f"Получено {len(go_movies)} фильмов из Go API")
            
            # Преобразуем в нужный формат и сохраняем
            os.makedirs('ml_service/data', exist_ok=True)
            with open('ml_service/data/movies.json', 'w', encoding='utf-8') as f:
                json.dump(go_movies, f, ensure_ascii=False, indent=2)
            
            print("Фильмы сохранены в ml_service/data/movies.json")
            
            # Автоматически запускаем preprocess_embeddings.py
            print("Запускаем обработку эмбеддингов...")
            run_preprocessing()
            
            return go_movies
        else:
            print(f"Ошибка получения данных из Go API: {response.status_code}")
            return None
            
    except requests.exceptions.ConnectionError:
        print("Go сервис недоступен на localhost:8000")
        return None
    except Exception as e:
        print(f"Ошибка при загрузке из Go API: {e}")
        return None

def run_preprocessing():
    """Запускает обработку эмбеддингов"""
    try:
        import subprocess
        import sys
        
        print("Начинаем создание эмбеддингов...")
        
        # Запускаем preprocess_embeddings.py
        result = subprocess.run([
            sys.executable, 'preprocess_embeddings.py'
        ], cwd='ml_service', capture_output=True, text=True)
        
        if result.returncode == 0:
            print("Эмбеддинги успешно созданы!")
            print(result.stdout)
        else:
            print("Ошибка при создании эмбеддингов:")
            print(result.stderr)
            
    except Exception as e:
        print(f"Ошибка запуска preprocess_embeddings.py: {e}")

def load_movies_data():
    """Загружает данные о фильмах и эмбеддинги"""
    global movies, movie_embeddings
    
    movies_file = 'ml_service/data/movies.json'
    embeddings_file = 'ml_service/model/embedder.pkl'
    
    # Проверяем наличие файлов
    if not os.path.exists(movies_file):
        print("Файл movies.json не найден, загружаем из Go API...")
        movies = load_movies_from_go()
        if not movies:
            print("ПРЕДУПРЕЖДЕНИЕ: Не удалось загрузить фильмы!")
            movies = []
            return
    else:
        # Загружаем существующий файл
        with open(movies_file, 'r', encoding='utf-8') as f:
            movies = json.load(f)
    
    # Загружаем эмбеддинги если есть
    if os.path.exists(embeddings_file):
        with open(embeddings_file, 'rb') as f:
            movie_embeddings = pickle.load(f)
        print(f"Загружено {len(movie_embeddings)} эмбеддингов")
    else:
        print("ПРЕДУПРЕЖДЕНИЕ: Файл embeddings не найден!")
        # Если есть фильмы, но нет эмбеддингов - создаем их
        if movies:
            print("Создаем эмбеддинги для загруженных фильмов...")
            run_preprocessing()
            # Пытаемся загрузить созданные эмбеддинги
            if os.path.exists(embeddings_file):
                with open(embeddings_file, 'rb') as f:
                    movie_embeddings = pickle.load(f)
                print(f"Загружено {len(movie_embeddings)} эмбеддингов")
            else:
                movie_embeddings = []
        else:
            movie_embeddings = []

# Загружаем данные при импорте
load_movies_data()

def determine_alpha(user_input, user_entities):
    """
    Учитывает плотность именованных сущностей в запросе
    """
    if not user_entities:
        return 1.0

    segmenter = Segmenter()
    emb = NewsEmbedding()
    morph_vocab = MorphVocab()
    ner_tagger = NewsNERTagger(emb)

    doc = Doc(user_input)
    doc.segment(segmenter)
    doc.tag_ner(ner_tagger)

    total_tokens = len(doc.tokens)
    important_weight = 0
    entity_tokens = 0
    for span in doc.spans:
        span.normalize(morph_vocab)

        if span.type == 'PER':
            important_weight += 2
        elif span.type in ('LOC', 'ORG'):
            important_weight += 1
        else:
            important_weight += 0.5

        entity_tokens += len(span.tokens)

    if total_tokens == 0:
        density = 0.0
    else:
        density = entity_tokens / total_tokens

    if len(user_entities) == 0:
        normalized_importance = 0.0
    else:
        normalized_importance = important_weight / len(user_entities)

    ner_strength = 0.5 * density + 0.5 * (normalized_importance / 2)
    alpha = 1 - ner_strength
    alpha = min(max(alpha, 0.3), 1.0)

    return alpha

@app.route('/recommend', methods=['POST'])
def recommend():
    """Обрабатываем поисковый запрос от Go сервера"""
    
    # Проверяем что данные загружены
    if not movies:
        return jsonify({
            'status': 'error',
            'error': 'Данные о фильмах не загружены. Попробуйте /sync-movies'
        }), 500
    
    if not movie_embeddings:
        return jsonify({
            'status': 'error', 
            'error': 'Эмбеддинги не загружены. Нужно пересчитать модель.'
        }), 500
    
    # 1. Получаем данные из запроса
    try:
        # Попробуем получить JSON данные
        if request.is_json:
            data = request.get_json()
            user_input = data.get('query', '')
            limit = data.get('limit', 10)
        else:
            # Или данные из формы
            user_input = request.form.get('query', '')
            limit = int(request.form.get('limit', 10))
            
        # Проверяем, что запрос не пустой
        if not user_input or user_input.strip() == '':
            return jsonify({
                'status': 'error',
                'error': 'Пустой поисковый запрос'
            }), 400
            
    except Exception as e:
        return jsonify({
            'status': 'error',
            'error': 'Неверный формат запроса'
        }), 400

    # 2. Обрабатываем запрос вашим алгоритмом
    try:
        # Предобработка текста
        preprocessed = preprocess_text(user_input)
        
        # Получаем вектор запроса
        user_vector = encode_text(preprocessed)
        
        # Извлекаем именованные сущности
        user_entities = extract_named_entities(user_input)

        # Вычисляем косинусную близость
        similarities = cosine_similarity([user_vector], movie_embeddings)[0]

        # Подсчёт совпадений по именованным сущностям
        entity_scores = []
        for movie in movies:
            movie_entities = movie.get('named_entities', [])
            user_set = set(user_entities)
            movie_set = set(movie_entities)

            union = user_set | movie_set
            if not union:
                score = 0.0
            else:
                match_count = len(user_set & movie_set)
                score = match_count / len(union)

            entity_scores.append(score)

        # Определяем коэффициент alpha
        alpha = determine_alpha(user_input, user_entities)
        
        # Комбинированный скор
        combined_scores = alpha * similarities + (1 - alpha) * np.array(entity_scores)

        # Получаем все результаты, отсортированные по релевантности
        sorted_indices = np.argsort(combined_scores)[::-1]
        
        # Фильтруем только релевантные результаты
        relevant_movies = []
        min_threshold = 0.1  # Минимальный порог релевантности
        
        for idx in sorted_indices:
            score = combined_scores[idx]
            if score > min_threshold and len(relevant_movies) < limit:
                movie = movies[idx].copy()
                movie['relevance_score'] = float(score)  # Добавляем скор для отладки
                relevant_movies.append(movie)
        
        top_movies = relevant_movies
        
        # 3. Возвращаем результат
        return jsonify({
            'status': 'success',
            'query': user_input,
            'found': len(top_movies),
            'alpha_coefficient': float(alpha),
            'min_threshold': min_threshold,
            'results': top_movies
        })
        
    except Exception as e:
        return jsonify({
            'status': 'error',
            'error': f'Ошибка обработки: {str(e)}'
        }), 500

@app.route('/health', methods=['GET'])
def health():
    """Проверка работоспособности сервиса"""
    return jsonify({
        'status': 'healthy',
        'service': 'ml-recommendation',
        'movies_loaded': len(movies),
        'embeddings_loaded': len(movie_embeddings)
    })

@app.route('/sync-movies', methods=['POST'])
def sync_movies():
    """Принудительно обновить список фильмов из Go API"""
    global movies, movie_embeddings
    
    movies = load_movies_from_go()  # Это автоматически запустит preprocessing
    if movies:
        # Перезагружаем эмбеддинги после обработки
        embeddings_file = 'ml_service/model/embedder.pkl'
        if os.path.exists(embeddings_file):
            with open(embeddings_file, 'rb') as f:
                movie_embeddings = pickle.load(f)
        
        return jsonify({
            'status': 'success',
            'message': f'Синхронизировано {len(movies)} фильмов и созданы эмбеддинги',
            'movies_count': len(movies),
            'embeddings_count': len(movie_embeddings)
        })
    else:
        return jsonify({
            'status': 'error',
            'error': 'Не удалось загрузить фильмы из Go API'
        }), 500

if __name__ == '__main__':
    print(f"Загружено {len(movies)} фильмов")
    print(f"Загружено {len(movie_embeddings)} эмбеддингов")
    app.run(host='0.0.0.0', port=5000, debug=True)