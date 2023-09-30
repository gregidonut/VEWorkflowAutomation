const editDialog = document.querySelector("#edit-dialog")

let currentActualCommitVids = []

async function fetchActualCommitVidFiles() {
    console.log("calling fetch on getFSVids endpoint...")
    const response = await fetch("/getFSVids");
    const data = await response.json();

    const fileList = document.querySelector("section.committed-timeline ul");
    if (data === null) {
        fileList.innerHTML = '';
        fileList.innerHTML = '<li>No committed files yet.</li>';
        return;
    }

    if (arraysAreEqualByLastModified(currentActualCommitVids, data)) {
        return;
    }

    currentActualCommitVids = data;

    fileList.innerHTML = '';
    currentActualCommitVids.forEach((fsVid) => {

        const listItem = document.createElement('li');

        const itemDivWrapper = document.createElement('div');
        itemDivWrapper.classList.add('commit-vid-wrapper');

        const video = document.createElement('video');
        video.width = 256;
        video.controls = true;

        const source = document.createElement('source');
        source.src = `/static/uploads/workspace/actualCommitVids/${fsVid.vBasePath}?timestamp=${fsVid.lastModified}`;
        source.type = 'video/mp4';

        const fallbackText = document.createTextNode('Your browser does not support the video tag.');

        video.appendChild(source);
        video.appendChild(fallbackText);

        itemDivWrapper.appendChild(video);
        const scriptSection = document.createElement('div')
        itemDivWrapper.appendChild(scriptSection)

        const scriptSectionCtrls = document.createElement("div")

        const editBtn = document.createElement("button")
        editBtn.className = "script-edit-btn"
        editBtn.innerHTML = `<img src="/static/assets/editIcon.svg" alt="edit-icon" width="20px"/>`
        editBtn.querySelector("img").addEventListener("click", function () {
            const existingVideos = editDialog.querySelectorAll('video');
            existingVideos.forEach((existingVideo) => {
                existingVideo.remove();
            });

            const videoToEditScriptOn = document.createElement("video")
            videoToEditScriptOn.width = 384;
            videoToEditScriptOn.controls = true;

            const source = document.createElement('source');
            source.src = `/static/uploads/workspace/rawCommitVids/${fsVid.vBasePath}`;
            source.type = 'video/mp4';

            const fallbackText = document.createTextNode('Your browser does not support the video tag.');

            videoToEditScriptOn.appendChild(source);
            videoToEditScriptOn.appendChild(fallbackText);


            editDialog.insertBefore(videoToEditScriptOn, editDialog.firstChild);
            editDialog.querySelector("textarea").value = ""
            editDialog.querySelector("textarea").value = fsVid.script

            editDialog.setAttribute("fsVid-prop", JSON.stringify(fsVid))
            editDialog.showModal()
        })

        const delBtn = document.createElement("button")
        delBtn.className = "commit-vid-delete-btn"
        delBtn.innerHTML = `<img src="/static/assets/delIcon.svg" alt="edit-icon" width="20px"/>`
        delBtn.querySelector("img").addEventListener("click", function () {
            console.log(fsVid.vBasePath)
        })

        scriptSectionCtrls.appendChild(editBtn)
        scriptSectionCtrls.appendChild(delBtn)

        const scriptText = document.createElement('p')
        scriptText.className = "script-text"
        scriptText.innerHTML = fsVid.script

        scriptSection.appendChild(scriptSectionCtrls)
        scriptSection.appendChild(scriptText)

        listItem.appendChild(itemDivWrapper);
        fileList.appendChild(listItem);
    });
}

fetchActualCommitVidFiles()
setInterval(fetchActualCommitVidFiles, 1000)

const editDialogCloseBtn = document.querySelector("button.close-edit-dialog-btn")
editDialogCloseBtn.addEventListener("click", () => {
    editDialog.close()
})

function arraysAreEqualByLastModified(array1, array2) {
    if (array1.length !== array2.length) {
        return false;
    }

    for (let i = 0; i < array1.length; i++) {
        if (array1[i].lastModified !== array2[i].lastModified) {
            return false;
        }
    }

    return true;
}
