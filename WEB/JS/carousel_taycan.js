const image = ["../../assets/taycan(1).jpg", "../../assets/taycan(2).jpg", "../../assets/taycan(3).jpg", "../../assets/taycan(4).jpg", "../../assets/taycan(5).jpg"];
var nombre = 0;

function suivant(sens) {
    nombre = nombre + sens;
    if (nombre > image.length - 1)
        nombre = 0;
    if (nombre < 0)
    nombre = image.length - 1;
    document.getElementById("image").src = image[nombre];
}