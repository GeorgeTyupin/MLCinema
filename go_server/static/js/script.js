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

function main() {
    getFilms();
    getCategories();
    customHeightTextarea();
    $(".film-card").click(redirectToFilm);
    $(".menu").on("click", ".menu-item", activeCategory);
}

main();
