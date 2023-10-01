let currentRawCommitVids = []

async function fetchRawCommitVidFiles() {
    const response = await fetch("/listCommittedFiles");
    const data = await response.json();

    const fileList = document.querySelector("section.committed-timeline ul");
    if (data === null) {
        return;
    }

    if (currentRawCommitVids.length === data.length) {
        return;
    }

    currentRawCommitVids = data;

    fileList.innerHTML = '';
    currentRawCommitVids.forEach((file, index) => {
        const video = document.createElement('video');
        video.width = 128;
        video.controls = true;

        const source = document.createElement('source');
        source.src = `/static/uploads/workspace/rawCommitVids/${file}`;
        source.type = 'video/mp4';

        const fallbackText = document.createTextNode('Your browser does not support the video tag.');

        video.appendChild(source);
        video.appendChild(fallbackText);

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

fetchRawCommitVidFiles()
setInterval(fetchRawCommitVidFiles, 1000)