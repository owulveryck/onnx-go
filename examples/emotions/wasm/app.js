const canvas = document.getElementById('canvas');
const context = canvas.getContext('2d');
const video = document.getElementById('video');

if (navigator.mediaDevices && navigator.mediaDevices.getUserMedia) {
    navigator.mediaDevices.getUserMedia({ video: true }).then(stream => {
        video.srcObject = stream;
        video.play();
    });
}

document.getElementById('smile').addEventListener('click', () => {
    context.drawImage(video, 0, 0, 64, 64);
});