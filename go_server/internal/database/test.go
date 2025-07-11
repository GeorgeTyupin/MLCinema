package database

import (
	"log"

	"github.com/GeorgeTyupin/MLCinema/go_server/internal/models"
)

func getOrCreateActor(name string) *models.Actor {
	var actor models.Actor
	if err := DB.Where("name = ?", name).First(&actor).Error; err != nil {
		actor = models.Actor{Name: name}
		if err := DB.Create(&actor).Error; err != nil {
			log.Printf("Ошибка при создании актёра %s: %v", name, err)
		}
	}
	return &actor
}

func getOrCreateCategory(name string) *models.Category {
	var category models.Category
	if err := DB.Where("name = ?", name).First(&category).Error; err != nil {
		category = models.Category{Name: name}
		if err := DB.Create(&category).Error; err != nil {
			log.Printf("Ошибка при создании категории %s: %v", name, err)
		}
	}
	return &category
}

func SeedTestData() {
	var count int64
	DB.Model(&models.Film{}).Count(&count)
	if count > 0 {
		return
	}

	// Список фильмов (оригинальный)
	films := []struct {
		Title       string
		Year        uint
		Country     string
		ImagePath   string
		Description string
		Actors      []string
		Categories  []string
	}{
		{"Интерстеллар", 2014, "США", "/static/img/interstellar.jpg", "Будущее Земли под угрозой. Астронавты ищут новую планету.", []string{"Мэттью МакКонахи", "Энн Хэтэуэй"}, []string{"Фантастика"}},
		{"Начало", 2010, "США", "/static/img/inception.jpg", "Проникновение в сны и внедрение идей.", []string{"Леонардо Ди Каприо", "Том Харди"}, []string{"Фантастика", "Триллер"}},
		{"Матрица", 1999, "США", "/static/img/matrix.jpg", "Реальность — это симуляция.", []string{"Киану Ривз", "Кэрри-Энн Мосс"}, []string{"Фантастика"}},
		{"Бойцовский клуб", 1999, "США", "/static/img/fight_club.jpg", "Мужчина создаёт подпольный бойцовский клуб.", []string{"Брэд Питт", "Эдвард Нортон"}, []string{"Драма"}},
		{"Форрест Гамп", 1994, "США", "/static/img/forrest_gump.jpg", "История жизни Форреста.", []string{"Том Хэнкс"}, []string{"Драма", "Биография"}},
		{"Побег из Шоушенка", 1994, "США", "/static/img/shawshank.jpg", "Заключённый планирует побег.", []string{"Тим Роббинс", "Морган Фримен"}, []string{"Драма"}},
		{"Темный рыцарь", 2008, "США", "/static/img/dark_knight.jpg", "Бэтмен против Джокера.", []string{"Кристиан Бейл", "Хит Леджер"}, []string{"Боевик"}},
		{"Гладиатор", 2000, "США", "/static/img/gladiator.jpg", "Генерал становится гладиатором.", []string{"Рассел Кроу"}, []string{"История", "Боевик"}},
		{"Зеленая миля", 1999, "США", "/static/img/green_mile.jpg", "Сверхъестественные способности заключённого.", []string{"Том Хэнкс", "Майкл Кларк Дункан"}, []string{"Драма", "Фэнтези"}},
		{"1+1", 2011, "Франция", "/static/img/intouchables.jpg", "Инвалид и сиделка становятся друзьями.", []string{"Франсуа Клюзе", "Омар Си"}, []string{"Драма", "Комедия"}},
		{"Остров проклятых", 2010, "США", "/static/img/shutter_island.jpg", "Маршалы США расследуют исчезновение пациентки.", []string{"Леонардо Ди Каприо"}, []string{"Триллер"}},
		{"Леон", 1994, "Франция", "/static/img/leon.jpg", "Киллер берёт под опеку девочку.", []string{"Жан Рено", "Натали Портман"}, []string{"Боевик", "Драма"}},
		{"Титаник", 1997, "США", "/static/img/titanic.jpg", "История любви на лайнере.", []string{"Леонардо Ди Каприо", "Кейт Уинслет"}, []string{"Мелодрама"}},
		{"Король Лев", 1994, "США", "/static/img/lion_king.jpg", "Молодой лев Симба ищет своё место.", []string{"Джонатан Тейлор Томас"}, []string{"Мультфильм"}},
		{"Властелин колец: Братство кольца", 2001, "Новая Зеландия", "/static/img/lotr_fellowship.jpg", "Фродо отправляется уничтожить Кольцо.", []string{"Элайджа Вуд", "Иэн Маккеллен"}, []string{"Фэнтези"}},
		{"Криминальное чтиво", 1994, "США", "/static/img/pulp_fiction.jpg", "Истории мафии в Лос-Анджелесе.", []string{"Джон Траволта", "Сэмюэл Л. Джексон"}, []string{"Криминал"}},
		{"Престиж", 2006, "США", "/static/img/prestige.jpg", "Противостояние двух фокусников.", []string{"Кристиан Бейл", "Хью Джекман"}, []string{"Драма", "Триллер"}},
		{"Социальная сеть", 2010, "США", "/static/img/social_network.jpg", "Создание Facebook и конфликты вокруг него.", []string{"Джесси Айзенберг", "Эндрю Гарфилд"}, []string{"Драма", "Биография"}},
		{"Джокер", 2019, "США", "/static/img/joker.jpg", "История становления Джокера.", []string{"Хоакин Феникс"}, []string{"Триллер"}},
		{"Бегущий по лезвию 2049", 2017, "США", "/static/img/blade_runner_2049.jpg", "Офицер Кей раскрывает тайну, угрожающую обществу.", []string{"Райан Гослинг", "Харрисон Форд"}, []string{"Фантастика"}},
	}

	for _, f := range films {
		var actors []*models.Actor
		for _, name := range f.Actors {
			actors = append(actors, getOrCreateActor(name))
		}

		var categories []*models.Category
		for _, cname := range f.Categories {
			categories = append(categories, getOrCreateCategory(cname))
		}

		film := models.Film{
			Title:       f.Title,
			Year:        f.Year,
			Country:     f.Country,
			ImagePath:   f.ImagePath,
			Description: f.Description,
			Actors:      actors,
			Categories:  categories,
		}

		if err := DB.Create(&film).Error; err != nil {
			log.Printf("Ошибка при создании фильма \"%s\": %v", film.Title, err)
		}
	}
}
