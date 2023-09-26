let currentRawCommitVids = []

async function fetchStaticFiles() {
    const response = await fetch("/listCommittedFiles");
    const data = await response.json();

    const fileList = document.querySelector("section.committed-timeline ul");
    if (data === null) {
        fileList.innerHTML = '';
        fileList.innerHTML = '<li>No committed files yet.</li>';
        return;
    }

    if (currentRawCommitVids.length === data.length) {
        return;
    }

    currentRawCommitVids = data;

    fileList.innerHTML = '';
    currentRawCommitVids.forEach((file, index) => {

        const listItem = document.createElement('li');

        const divWrapper = document.createElement('div');
        divWrapper.classList.add('commit-vid-wrapper');

        const video = document.createElement('video');
        video.width = 128;
        video.controls = true;

        const source = document.createElement('source');
        source.src = `/static/uploads/workspace/rawCommitVids/${file}`;
        source.type = 'video/mp4';

        const fallbackText = document.createTextNode('Your browser does not support the video tag.');

        video.appendChild(source);
        video.appendChild(fallbackText);

        divWrapper.appendChild(video);

        listItem.appendChild(divWrapper);

        fileList.appendChild(listItem);

        if (index === currentRawCommitVids.length - 1) {
            const existingVideos = scriptingDialog.querySelectorAll('video');
            existingVideos.forEach((existingVideo) => {
                existingVideo.remove();
            });

            const lastVideo = video.cloneNode(true);
            lastVideo.className = "dialog-preview-vid"
            lastVideo.width = 384;
            scriptingDialog.insertBefore(lastVideo, scriptingDialog.firstChild);
        }
    });
}

fetchStaticFiles()
setInterval(fetchStaticFiles, 1000)