let currentCommitVids = []

async function fetchStaticFiles() {
    const response = await fetch("/listCommittedFiles");
    const data = await response.json();

    const fileList = document.querySelector("section.committed-timeline ul");
    if (data === null) {
        fileList.innerHTML = '';
        fileList.innerHTML = '<li>No committed files yet.</li>';
        return;
    }

    if (currentCommitVids.length === data.length) {
        return;
    }

    currentCommitVids = data;

    fileList.innerHTML = '';
    currentCommitVids.forEach((file) => {

        const listItem = document.createElement('li');

        const divWrapper = document.createElement('div');
        divWrapper.classList.add('commit-vid-wrapper');

        const video = document.createElement('video');
        video.width = 128;
        video.controls = true;

        const source = document.createElement('source');
        source.src = `/static/uploads/workspace/commitVids/${file}`;
        source.type = 'video/mp4';

        const fallbackText = document.createTextNode('Your browser does not support the video tag.');

        video.appendChild(source);
        video.appendChild(fallbackText);

        divWrapper.appendChild(video);

        listItem.appendChild(divWrapper);

        fileList.appendChild(listItem);
    });
}

fetchStaticFiles()
setInterval(fetchStaticFiles, 1000)