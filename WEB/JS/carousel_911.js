const image = ["../../assets/911(1).jpg", "../../assets/911(2).jpg", "../../assets/911(3).jpg", "../../assets/911(4).jpg", "../../assets/911(5).jpg"];
var nombre = 0;

function suivant(sens) {
    nombre = nombre + sens;
    if (nombre > image.length - 1)
        nombre = 0;
    if (nombre < 0)
    nombre = image.length - 1;
    document.getElementById("image").src = image[nombre];
}