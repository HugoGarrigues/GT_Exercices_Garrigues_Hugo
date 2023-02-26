const image = ["../../assets/718(1).jpg", "../../assets/718(2).jpg", "../../assets/718(3).jpg", "../../assets/718(4).jpg", "../../assets/718(5).jpg"];
var nombre = 0;

function suivant(sens) {
    nombre = nombre + sens;
    if (nombre > image.length - 1)
        nombre = 0;
    if (nombre < 0)
    nombre = image.length - 1;
    document.getElementById("image").src = image[nombre];
}