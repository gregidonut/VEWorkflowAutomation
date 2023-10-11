const form = <HTMLFormElement>document.getElementById("upload-form");
const uploadProgressElement = <HTMLProgressElement>document.getElementById("upload-progress");

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
});

