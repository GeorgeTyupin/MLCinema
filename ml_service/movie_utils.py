import json
import pymorphy2
from natasha import Doc, NewsNERTagger, NewsEmbedding, Segmenter, MorphVocab
from sentence_transformers import SentenceTransformer
from sklearn.metrics.pairwise import cosine_similarity
import numpy as np
import re

morph = pymorphy2.MorphAnalyzer()          # для лемматизации
segmenter = Segmenter()                    # сегментация текста
emb = NewsEmbedding()                      # эмбеддинги для NER
ner_tagger = NewsNERTagger(emb)            # модель NER
morph_vocab = MorphVocab()                 # нормализация сущностей
model = SentenceTransformer('distiluse-base-multilingual-cased-v1') # мультиязычный трансформер

model = SentenceTransformer('distiluse-base-multilingual-cased-v1')

STOP_WORDS = set(['в', 'на', 'и', 'не', 'что', 'это', 'как', 'из', 'с', 'по', 'а', 'о', 'для', 'но'])

def preprocess_text(text):
    text = text.lower()
    text = re.sub(r'[^а-яa-z0-9\s]', ' ', text)
    tokens = text.split()
    lemmas = [morph.parse(token)[0].normal_form for token in tokens if token not in STOP_WORDS]
    return ' '.join(lemmas)

def extract_named_entities(text):
    doc = Doc(text)
    doc.segment(segmenter)
    doc.tag_ner(ner_tagger)
    return [span.text.lower() for span in doc.spans]

def encode_text(text):
    return model.encode([text])[0]

def load_movies(path='ml_servicedata/movies.json'):
    with open(path, 'r', encoding='utf-8') as f:
        return json.load(f)
