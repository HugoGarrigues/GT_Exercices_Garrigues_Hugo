const image = ["../assets/carousel1.webp", "../assets/carousel2.webp", "../assets/carousel3.jpg", "../assets/carousel4.webp"];
var nombre = 0;

function suivant(sens) {
    nombre = nombre + sens;
    if (nombre > image.length - 1)
        nombre = 0;
    if (nombre < 0)
    nombre = image.length - 1;
    document.getElementById("image").src = image[nombre];
}