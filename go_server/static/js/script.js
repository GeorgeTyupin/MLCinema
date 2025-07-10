var films = [];

function customHeightTextarea() {
    const textarea = document.querySelector(".search-wrapper textarea");
    textarea.addEventListener("input", function () {
        this.style.height = "auto";
        this.style.height = (this.scrollHeight) + "px";
    });
}

function renderFilmPosters() {
    if (!films.length) {
        console.warn("Нет фильмов для отображения");
        return;
    }

    $(".film-card").each(function(index, row) {
        if (films[index]) {
            $(row).removeClass("skeleton");
            let cardImg = $(row).children()[0];
            cardImg.src = `..${films[index].imagePath}`;
            $(row).attr("id", films[index].id);
        }
    });
}

function renderCategories(categories) {
    categories.forEach(category => {
        $(".menu").append(`<li class="menu-item" id="${category.id}"><span>${category.name}</span></li>`);
    });

    let maxWidth = 0;

    $('.menu-item').each(function () {
        let width = $(this).width();
        if (width > maxWidth) {
            maxWidth = width;
        }
    });
    
    console.log(maxWidth)

    $('.menu-item').each(function () {
        $(this).width(maxWidth);
    });
}

function redirectToFilm() {
    window.location.href = `/film?film_id=${this.id}`;
}

function getFilms() {
    $.post("/api/get-films", function(response) {
        films = response;       
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
}

main();
