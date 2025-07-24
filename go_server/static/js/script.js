var films = [];

function customHeightTextarea() {
    const textarea = document.querySelector(".search-wrapper textarea");
    textarea.addEventListener("input", function () {
        this.style.height = "auto";
        this.style.height = (this.scrollHeight) + "px";
    });
}

function activeCategory() {
    let width = $(this).width();
    let height = $(this).height();

    if ($(this).hasClass('active-category')) {
        $(this).removeClass('active-category');
        $(this).width(width - 50);
        $(this).height(height - 10);
        $(".section-title").eq(0).text("Специальная подборка для вас");
        $(".film-row").eq(0).children().remove();
        $(".section-title").eq(1).removeClass('hide');
        $(".film-row").eq(1).removeClass('hide');
        $(".film-row").eq(0).append(`                    
            <div class="film-card skeleton" id=""><img src="" alt=""></div>
            <div class="film-card skeleton" id=""><img src="" alt=""></div>
            <div class="film-card skeleton" id=""><img src="" alt=""></div>
            <div class="film-card skeleton" id=""><img src="" alt=""></div>
            <div class="film-card skeleton" id=""><img src="" alt=""></div>
        `);
        renderFilmPosters();
        $(".film-card").click(redirectToFilm);
    } else {
        $('.menu-item').each(function () {
            if ($(this).hasClass('active-category')) {
                console.log(this)
                $(this).removeClass('active-category');
                $(this).width(width);
                $(this).height(height);
            }
        });
        $(this).addClass('active-category');
        $(this).width(width + 50);
        $(this).height(height + 10);
        renderFilmByCategory($(this).children().eq(0).text(), $(this).attr("id").slice(7));
        $(".film-card").click(redirectToFilm);
    }
}

function renderFilmByCategory(cat_text, cat_id) {
    $(".section-title").eq(1).addClass('hide');
    $(".film-row").eq(1).addClass('hide');
    $(".section-title").eq(0).text(cat_text)
    $(".film-row").eq(0).children().remove();
    films.forEach( (film) => {
        film.categories.forEach((cat) => {
            if (cat.id == cat_id) {
                $(".film-row").eq(0).append(`<div class="film-card" id="film_id-${film.id}"><img src="${film.imagePath}" alt=""></div>`);
            }
        });
    });
}

function renderFilmPosters() {
    if (!films.length) {
        console.warn("Нет фильмов для отображения");
    }

    $(".film-card").each(function(index, row) {
        if (films[index]) {
            $(row).removeClass("skeleton");
            let cardImg = $(row).children()[0];
            cardImg.src = `${films[index].imagePath}`;
            $(row).attr("id", `film_id-${films[index].id}`);
        }
    });
}

function renderCategories(categories) {
    categories.forEach(category => {
        $(".menu").append(`<li class="menu-item" id="cat_id-${category.id}"><span>${category.name}</span></li>`);
    });

    let maxWidth = 0;

    $('.menu-item').each(function () {
        let width = $(this).width();
        if (width > maxWidth) {
            maxWidth = width;
        }
    });

    $('.menu-item').each(function () {
        $(this).width(maxWidth);
    });
}

function redirectToFilm() {
    window.location.href = `/film?film_id=${this.id.slice(8)}`;
}

function getFilms() {
    $.post("/api/get-films", function(response) {
        films = response;   
        console.log(films)    
        renderFilmPosters();       
    });
}

function getCategories() {
    $.post("/api/get-categories", function(response) {    
        renderCategories(response);
    });
}

// Функция ML поиска
function performMLSearch() {
    const query = $(".search-wrapper textarea").val().trim();
    
    if (!query) {
        showError("Введите поисковый запрос", 1);
        return;
    }

    // Показываем индикатор загрузки
    showLoading();

    $.post("/", {"query": query})
        .done(function(response) {
            handleSearchSuccess(response, query);
        })
        .fail(function(xhr) {
            handleSearchError(xhr, query);
        });
}

// Обработка успешного ответа
function handleSearchSuccess(response, query) {
    hideLoading();

    // Если это массив фильмов (успешный поиск)
    if (Array.isArray(response)) {
        displaySearchResults(response, query);
    }
    // Если это объект с результатами
    else if (response.results && Array.isArray(response.results)) {
        displaySearchResults(response.results, query);
    }
    // Если нет результатов
    else {
        showNoResults(query);
    }
}

// Обработка ошибок
function handleSearchError(xhr, query) {
    hideLoading();
    
    try {
        const error = JSON.parse(xhr.responseText);
        
        if (error.code === 1) {
            showError("Введите поисковый запрос", 1);
        } else if (error.code === 2) {
            showError("Сервис поиска временно недоступен", 2);
        } else {
            showError("Произошла ошибка при поиске", 0);
        }
    } catch (e) {
        showError("Произошла ошибка при поиске", 0);
    }
}

// Отображение результатов поиска
function displaySearchResults(movies, query) {
    // Очищаем текущие результаты
    $(".film-row").eq(0).empty();
    $(".section-title").eq(0).text(`Результаты поиска: "${query}"`);
    
    // Скрываем вторую секцию
    $(".section-title").eq(1).addClass('hide');
    $(".film-row").eq(1).addClass('hide');

    if (movies.length === 0) {
        showNoResults(query);
        return;
    }

    // Отображаем найденные фильмы
    movies.forEach((movie, index) => {
        if (index < 20) { // Показываем максимум 20 результатов
            $(".film-row").eq(0).append(`
                <div class="film-card" id="film_id-${movie.id}">
                    <img src="${movie.imagePath}" alt="${movie.title}" 
                         onerror="this.src='/static/img/no-image.jpg'">
                </div>
            `);
        }
    });
    
    // Обновляем обработчики клика
    $(".film-card").off('click').on('click', redirectToFilm);
}

// Показ сообщения "нет результатов"
function showNoResults(query) {
    $(".film-row").eq(0).html(`
        <div class="no-results">
            <p>По запросу "${query}" ничего не найдено</p>
            <p>Попробуйте изменить поисковый запрос</p>
        </div>
    `);
}

// Показ ошибки
function showError(message, code) {
    $(".film-row").eq(0).html(`
        <div class="error-message">
            <p>❌ ${message}</p>
            ${code === 2 ? '<p>Попробуйте повторить поиск позже</p>' : ''}
        </div>
    `);
}

// Показ индикатора загрузки
function showLoading() {
    $(".section-title").eq(0).text("Поиск...");
    $(".film-row").eq(0).html(`
        <div class="loading">
            <div class="loading-spinner"></div>
            <p>Ищем подходящие фильмы...</p>
        </div>
    `);
    
    // Скрываем вторую секцию
    $(".section-title").eq(1).addClass('hide');
    $(".film-row").eq(1).addClass('hide');
}

// Скрытие индикатора загрузки
function hideLoading() {
    // Загрузка скрывается автоматически при отображении результатов
}

// Сброс поиска
function resetSearch() {
    $(".search-wrapper textarea").val('');
    $(".section-title").eq(0).text("Специальная подборка для вас");
    $(".section-title").eq(1).removeClass('hide');
    $(".film-row").eq(1).removeClass('hide');
    
    // Восстанавливаем изначальные фильмы
    $(".film-row").eq(0).html(`
        <div class="film-card skeleton"><img src="" alt=""></div>
        <div class="film-card skeleton"><img src="" alt=""></div>
        <div class="film-card skeleton"><img src="" alt=""></div>
        <div class="film-card skeleton"><img src="" alt=""></div>
        <div class="film-card skeleton"><img src="" alt=""></div>
    `);
    
    renderFilmPosters();
    $(".film-card").click(redirectToFilm);
}

// Обновляем обработчики событий
$(document).ready(function() {
    // Поиск по Enter
    $(".search-wrapper textarea").on("keypress", function(e) {
        if (e.which === 13 && !e.shiftKey) {
            e.preventDefault();
            performMLSearch();
        }
    });

    // Поиск по клику на иконку
    $(".search-wrapper i").on("click", function() {
        performMLSearch();
    });

    // Очистка поиска по Escape
    $(".search-wrapper textarea").on("keydown", function(e) {
        if (e.which === 27) { // Escape
            resetSearch();
        }
    });
});

function main() {
    getFilms();
    getCategories();
    customHeightTextarea();
    $(".film-card").click(redirectToFilm);
    $(".menu").on("click", ".menu-item", activeCategory);
}

main();
