function getVidSrc() {
    return document.querySelector("video.dialog-preview-vid")
        .querySelector('source')
        .getAttribute('src');
}

function submitScript() {
    const scriptTextNode = document.getElementById('script')
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/writeScriptToFile", true);
    xhr.setRequestHeader('Content-Type', 'applicationOld/json');
    xhr.send(JSON.stringify({
        script: scriptTextNode.value, vidPath: getVidSrc(),
    }));

    scriptTextNode.value = ""
    scriptingDialog.close()
    combineFSVidWithScriptTTSAudio()
}


function combineFSVidWithScriptTTSAudio() {
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/generateFSVids", true);
    xhr.setRequestHeader('Content-Type', 'applicationOld/json');
    xhr.send(JSON.stringify({
        combineLastVideoWithScriptAudio: true
    }));
}