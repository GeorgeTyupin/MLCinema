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
            console.log(films[index])
            $(row).attr("id", films[index].id);
        }
    });
}

function getFilms() {
    $.post("/api/get-films", function(response) {
        films = response;          
        renderFilmPosters();       
    });
}

function main() {
    getFilms();
    customHeightTextarea();
}

main();
