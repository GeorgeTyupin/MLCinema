from flask import Flask, request, jsonify
import numpy as np
import pickle
from movie_utils import preprocess_text, extract_named_entities, encode_text, load_movies
from sklearn.metrics.pairwise import cosine_similarity
from natasha import Doc, NewsNERTagger, NewsEmbedding, Segmenter, MorphVocab
import json

app = Flask(__name__)

# Загружаем данные
movies = load_movies('ml_service/data/movies.json')
with open('ml_service/model/embedder.pkl', 'rb') as f:
    movie_embeddings = pickle.load(f)

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

@app.route('/test', methods=['GET', 'POST'])
def test():
    """Тестовый эндпоинт для отладки"""
    if request.method == 'GET':
        return jsonify({
            'message': 'ML сервис работает',
            'movies_count': len(movies),
            'sample_movie': movies[0]['title'] if movies else 'Нет фильмов'
        })
    else:
        # POST запрос для тестирования
        query = request.form.get('query', 'тест')
        return recommend()

if __name__ == '__main__':
    print(f"Загружено {len(movies)} фильмов")
    print(f"Загружено {len(movie_embeddings)} эмбеддингов")
    app.run(host='0.0.0.0', port=5000, debug=True)