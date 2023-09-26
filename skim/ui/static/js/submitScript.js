function getVidSrc() {
    return document.querySelector("video.dialog-preview-vid")
        .querySelector('source')
        .getAttribute('src');
}

function submitScript(scriptTextNode) {
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/writeScriptToFile", true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({
        script: scriptTextNode.value, vidPath: getVidSrc(),
    }));

    scriptTextNode.value = ""
    scriptingDialog.close()
    combineFSVidWithScriptTTSAudio()
}


function combineFSVidWithScriptTTSAudio() {
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/fsVids", true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({
        combineLastVideoWithScriptAudio: true
    }));
}