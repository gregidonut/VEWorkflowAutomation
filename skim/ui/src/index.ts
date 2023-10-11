const form = <HTMLFormElement>document.getElementById("upload-form");
const uploadProgressElement = <HTMLProgressElement>document.getElementById("upload-progress");
const copyProgressElement = <HTMLProgressElement>document.getElementById("copy-progress");

let intervalId: NodeJS.Timeout;
form.addEventListener("submit", (e) => {
    e.preventDefault();

    const xhr = new XMLHttpRequest();

    xhr.upload.addEventListener("progress", (event) => {
        if (event.lengthComputable) {
            uploadProgressElement.value = (event.loaded / event.total) * 100;
        }
    });

    xhr.open("POST", "/upload", true);

    const formData = new FormData(form);
    xhr.send(formData);
    intervalId = setInterval(getCopyProgressPercentage, 1000);
});

async function getCopyProgressPercentage() {
    const response = await fetch("/copyProgress");
    if (!response.ok) {
        console.error(`Failed to fetch copy progress: ${response.status}`);
        return;
    }

    const percentage = await response.json();
    copyProgressElement.value = percentage;

    if (percentage === 100) {
        clearInterval(intervalId); // Stop the interval when the percentage reaches 100
    }
}
