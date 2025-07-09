function customHeightTextarea() {
    const textarea = document.querySelector(".search-wrapper textarea");
    textarea.addEventListener("input", function () {
        this.style.height = "auto";
        this.style.height = (this.scrollHeight) + "px";
    });
}

function getFilms() {
    $.post("/api/get-films", 'hello', success = function(response) {
        console.log(response);
    });
}

function main() {
    console.log('Привет')
    getFilms();
    customHeightTextarea();
}

main();