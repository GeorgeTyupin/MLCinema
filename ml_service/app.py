from flask import Flask, request, jsonify
import numpy as np
import pickle
from movie_utils import preprocess_text, extract_named_entities, encode_text, load_movies
from sklearn.metrics.pairwise import cosine_similarity
from natasha import Doc, NewsNERTagger, NewsEmbedding, Segmenter, MorphVocab

app = Flask(__name__)

# Загружаем данные
movies = load_movies('ml_service/data/movies.json')
with open('ml_service/model/embedder.pkl', 'rb') as f:
    movie_embeddings = pickle.load(f)

def determine_alpha(user_input, user_entities):
    """
    Учитывает плотность именованных сущностей в запросе
    Учитывает важность типов сущностей
    """
    if not user_entities:
        return 1.0  # если нет сущностей работаем только с семантикой

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

        if span.type == 'PER':  # Имя актёра, режиссёра и т.п.
            important_weight += 2
        elif span.type in ('LOC', 'ORG'):  # Локации, организации
            important_weight += 1
        else:
            important_weight += 0.5  # Прочее (даты, числа и т.п.)

        entity_tokens += len(span.tokens)

    if total_tokens == 0:
        density = 0.0
    else:
        density = entity_tokens / total_tokens

    # Нормализованная доля весов (к сущностям пользователя)
    if len(user_entities) == 0:
        normalized_importance = 0.0
    else:
        normalized_importance = important_weight / len(user_entities)

    # Комбинированная значимость NER
    ner_strength = 0.5 * density + 0.5 * (normalized_importance / 2)

    # Инвертируем, чтобы получить alpha (вес cosine similarity)
    alpha = 1 - ner_strength

    # Ограничим диапазон значений alpha
    alpha = min(max(alpha, 0.3), 1.0)

    return alpha

@app.route('/recommend', methods=['POST'])
def recommend():
    user_input = request.json.get('query')
    print(user_input)
    if not user_input:
        return jsonify({'error': 'No query provided'}), 400

    preprocessed = preprocess_text(user_input) #очищает, лемматизирует, убирает стоп-слова
    user_vector = encode_text(preprocessed) #делает вектор
    user_entities = extract_named_entities(user_input) # вычисляет локации, имена через Natasha

    # Вычисляем косинусную близость
    similarities = cosine_similarity([user_vector], movie_embeddings)[0] # по сути про скалярное произведение

    # Подсчёт совпадений по именованным сущностям
    entity_scores = []
    for movie in movies:
        movie_entities = movie.get('named_entities', [])
        user_set = set(user_entities)
        movie_set = set(movie_entities)

        union = user_set | movie_set
        if not union:
            score = 0.0  # Невозможно оценить пересечение - нет данных
        else:
            match_count = len(user_set & movie_set)
            score = match_count / len(union)

        entity_scores.append(score)


    alpha = determine_alpha(user_input, user_entities)
    # Комбинированный скор (веса можно подобрать эмпирически)
    combined_scores = alpha * similarities + (1 - alpha) * np.array(entity_scores)

    # Топ-10
    top_indices = np.argsort(combined_scores)[::-1][:10]
    top_movies = [movies[i] for i in top_indices]

    return jsonify(top_movies)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8000)

