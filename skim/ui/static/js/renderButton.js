const renderButton = document.querySelector("section.scripting-space section.main-ui-options button")
renderButton.addEventListener("click", async function () {
    console.log("rendering video...")


    const response = await fetch("/commitFinalVid")
    const finalVideoData = await response.json();

    const finalDivWrapper = document.createElement('div');
    finalDivWrapper.classList.add("final-vid-wrapper");

    const finalVid = document.createElement("video")
    finalVid.width = 640
    finalVid.controls = true;

    const source = document.createElement('source');
    source.src = `/static/uploads/workspace/actualCommitVids/${finalVideoData.vBasePath}?timestamp=${finalVideoData.lastModified}`;
    source.type = 'video/mp4';

    const fallbackText = document.createTextNode('Your browser does not support the video tag.');

    finalVid.appendChild(source);
    finalVid.appendChild(fallbackText);
})
