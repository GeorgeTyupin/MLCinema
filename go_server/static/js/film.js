var film = {};

function renderCurrentFilm() {
    console.log(film);   
    console.log($(".poster img"))
    $(".poster img").attr("src", film.imagePath);
    $(".film-title").text(film.title);
    $(".film-meta").text(film.country + "," + film.year + "·" + film.genre)
    $(".film-description").text(film.description)
}

function getFilmData(film_id) {
    console.log(film_id)
    $.post("/api/get-current-film", {"film_id" : film_id}, function(response) {
        film = response;  
        renderCurrentFilm();
    });
}

function main() {
    let params = new URLSearchParams(window.location.search);
    film_id = params.get("film_id");
    getFilmData(film_id);
}

main();
