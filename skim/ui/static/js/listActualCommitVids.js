let currentActualCommitVids = []

async function fetchActualCommitVidFiles() {
    const response = await fetch("/getFSVids");
    const data = await response.json();

    const fileList = document.querySelector("section.committed-timeline ul");
    if (data === null) {
        fileList.innerHTML = '';
        fileList.innerHTML = '<li>No committed files yet.</li>';
        return;
    }

    if (currentActualCommitVids.length === data.length) {
        return;
    }

    currentActualCommitVids = data;

    fileList.innerHTML = '';
    currentActualCommitVids.forEach((file, index) => {

        const listItem = document.createElement('li');

        const divWrapper = document.createElement('div');
        divWrapper.classList.add('commit-vid-wrapper');

        const video = document.createElement('video');
        video.width = 256;
        video.controls = true;

        const source = document.createElement('source');
        source.src = `/static/uploads/workspace/actualCommitVids/${file}`;
        source.type = 'video/mp4';

        const fallbackText = document.createTextNode('Your browser does not support the video tag.');

        video.appendChild(source);
        video.appendChild(fallbackText);

        divWrapper.appendChild(video);

        listItem.appendChild(divWrapper);

        fileList.appendChild(listItem);
    });
}

fetchActualCommitVidFiles()
setInterval(fetchActualCommitVidFiles, 1000)