const image = ["../../assets/panamera(1).jpg", "../../assets/panamera(2).jpg", "../../assets/panamera(3).jpg", "../../assets/panamera(4).jpg", "../../assets/panamera(5).jpg"];
var nombre = 0;

function suivant(sens) {
    nombre = nombre + sens;
    if (nombre > image.length - 1)
        nombre = 0;
    if (nombre < 0)
    nombre = image.length - 1;
    document.getElementById("image").src = image[nombre];
}