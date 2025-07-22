import json
import pickle
from movie_utils import preprocess_text, encode_text, extract_named_entities

# Загрузка фильмов в новой структуре
with open('ml_service/data/movies.json', 'r', encoding='utf-8') as f:
    movies = json.load(f)

embeddings = []

for movie in movies:
    # Собираем весь доступный текст в одну строку
    full_text_parts = [
        movie.get('title', ''),
        movie.get('description', ''),
        movie.get('genre', ''),
        ' '.join([cat['name'] for cat in movie.get('categories', [])]),
        ' '.join([actor['name'] for actor in movie.get('actors', [])])
    ]
    full_text = ' '.join(full_text_parts)

    # Предобработка
    preprocessed = preprocess_text(full_text)

    # Извлечение именованных сущностей
    movie['named_entities'] = extract_named_entities(full_text)

    # Получение вектора
    embeddings.append(encode_text(preprocessed))

# Сохраняем эмбеддинги
with open('ml_service/model/embedder.pkl', 'wb') as f:
    pickle.dump(embeddings, f)

# Обновляем movies.json с named_entities
with open('ml_service/data/movies.json', 'w', encoding='utf-8') as f:
    json.dump(movies, f, ensure_ascii=False, indent=2)

print("Готово: эмбеддинги сохранены, named_entities добавлены в movies.json")
