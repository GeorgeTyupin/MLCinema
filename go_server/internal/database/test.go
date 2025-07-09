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

func SeedTestData() {
	var count int64
	DB.Model(&models.Film{}).Count(&count)
	if count > 0 {
		return
	}

	films := []struct {
		Title       string
		Year        uint
		Country     string
		Genre       string
		ImagePath   string
		Description string
		Actors      []string
	}{
		{"Интерстеллар", 2014, "США", "Фантастика", "/static/img/interstellar.jpg", "Будущее Земли под угрозой. Астронавты ищут новую планету.", []string{"Мэттью МакКонахи", "Энн Хэтэуэй"}},
		{"Начало", 2010, "США", "Фантастика", "/static/img/inception.jpg", "Проникновение в сны и внедрение идей.", []string{"Леонардо Ди Каприо", "Том Харди"}},
		{"Матрица", 1999, "США", "Фантастика", "/static/img/matrix.jpg", "Реальность — это симуляция.", []string{"Киану Ривз", "Кэрри-Энн Мосс"}},
		{"Бойцовский клуб", 1999, "США", "Драма", "/static/img/fight_club.jpg", "Мужчина создаёт подпольный бойцовский клуб.", []string{"Брэд Питт", "Эдвард Нортон"}},
		{"Форрест Гамп", 1994, "США", "Драма", "/static/img/forrest_gump.jpg", "История жизни Форреста.", []string{"Том Хэнкс"}},
		{"Побег из Шоушенка", 1994, "США", "Драма", "/static/img/shawshank.jpg", "Заключённый планирует побег.", []string{"Тим Роббинс", "Морган Фримен"}},
		{"Темный рыцарь", 2008, "США", "Боевик", "/static/img/dark_knight.jpg", "Бэтмен против Джокера.", []string{"Кристиан Бейл", "Хит Леджер"}},
		{"Гладиатор", 2000, "США", "Исторический", "/static/img/gladiator.jpg", "Генерал становится гладиатором.", []string{"Рассел Кроу"}},
		{"Зеленая миля", 1999, "США", "Драма", "/static/img/green_mile.jpg", "Сверхъестественные способности заключённого.", []string{"Том Хэнкс", "Майкл Кларк Дункан"}},
		{"1+1", 2011, "Франция", "Комедия, драма", "/static/img/intouchables.jpg", "Инвалид и сиделка становятся друзьями.", []string{"Франсуа Клюзе", "Омар Си"}},
		{"Остров проклятых", 2010, "США", "Триллер", "/static/img/shutter_island.jpg", "Маршалы США расследуют исчезновение пациентки.", []string{"Леонардо Ди Каприо"}},
		{"Леон", 1994, "Франция", "Боевик, драма", "/static/img/leon.jpg", "Киллер берёт под опеку девочку.", []string{"Жан Рено", "Натали Портман"}},
		{"Титаник", 1997, "США", "Мелодрама", "/static/img/titanic.jpg", "История любви на лайнере.", []string{"Леонардо Ди Каприо", "Кейт Уинслет"}},
		{"Король Лев", 1994, "США", "Мультфильм", "/static/img/lion_king.jpg", "Молодой лев Симба ищет своё место.", []string{"Джонатан Тейлор Томас"}},
		{"Властелин колец: Братство кольца", 2001, "Новая Зеландия", "Фэнтези", "/static/img/lotr_fellowship.jpg", "Фродо отправляется уничтожить Кольцо.", []string{"Элайджа Вуд", "Иэн Маккеллен"}},
		{"Криминальное чтиво", 1994, "США", "Криминал", "/static/img/pulp_fiction.jpg", "Истории мафии в Лос-Анджелесе.", []string{"Джон Траволта", "Сэмюэл Л. Джексон"}},
		{"Престиж", 2006, "США", "Драма, триллер", "/static/img/prestige.jpg", "Противостояние двух фокусников.", []string{"Кристиан Бейл", "Хью Джекман"}},
		{"Социальная сеть", 2010, "США", "Драма", "/static/img/social_network.jpg", "Создание Facebook и конфликты вокруг него.", []string{"Джесси Айзенберг", "Эндрю Гарфилд"}},
		{"Джокер", 2019, "США", "Триллер", "/static/img/joker.jpg", "История становления Джокера.", []string{"Хоакин Феникс"}},
		{"Бегущий по лезвию 2049", 2017, "США", "Фантастика", "/static/img/blade_runner_2049.jpg", "Офицер Кей раскрывает тайну, угрожающую обществу.", []string{"Райан Гослинг", "Харрисон Форд"}},
	}

	for _, f := range films {
		var actors []*models.Actor
		for _, name := range f.Actors {
			actors = append(actors, getOrCreateActor(name))
		}

		film := models.Film{
			Title:       f.Title,
			Year:        f.Year,
			Country:     f.Country,
			Genre:       f.Genre,
			ImagePath:   f.ImagePath,
			Description: f.Description,
			Actors:      actors,
		}

		if err := DB.Create(&film).Error; err != nil {
			log.Printf("Ошибка при создании фильма \"%s\": %v", film.Title, err)
		}
	}
}
