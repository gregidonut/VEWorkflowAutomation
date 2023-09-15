const COLOR_BLUE = "rgb(0, 0, 255)"
const COLOR_AQUA = "rgb(0, 255, 255)"
const COLOR_GREY = "rgb(97, 97, 97)"

let isMouseDown = false;
let selectedElements = [];

let vidElements = document.querySelectorAll("div.vid-wrapper");

vidElements.forEach(function (e) {
    e.addEventListener("mousedown", function () {
        isMouseDown = true;
        toggleSelection(e);
    });

    e.addEventListener("mouseover", function () {
        if (isMouseDown) {
            toggleSelection(e);
        }
    });
});

document.addEventListener("mouseup", function () {
    isMouseDown = false;
});

function toggleSelection(element) {
    if (element.style.backgroundColor === COLOR_BLUE) {
        element.style.backgroundColor = COLOR_AQUA;
        const index = selectedElements.indexOf(element);
        if (index > -1) {
            selectedElements.splice(index, 1);
        }
        return
    }

    if (selectedElements.length >= 4) {
        return
    }

    element.style.backgroundColor = COLOR_BLUE;
    selectedElements.push(element);
}

let button = document.querySelector("section.initial-timeline button.initial-timeline-select-btn");
button.addEventListener("click", function () {
    selectedElements.forEach(function (e) {
        e.style.backgroundColor = COLOR_GREY;
    });
    selectedElements.sort(function (a, b) {
        const srcA = a.querySelector("video source").getAttribute("src");
        const srcB = b.querySelector("video source").getAttribute("src");
        return srcA.localeCompare(srcB);
    });
    selectedElements.forEach(function (e) {
        const srcAttr = e.querySelector("video source").getAttribute("src")
        console.log(srcAttr)
    })

    let srcFilePaths = []
    selectedElements.forEach(function (e) {
        srcFilePaths.push(e.querySelector("video source").getAttribute("src"))
    })

    if (srcFilePaths.length < 1) {
        selectedElements = []
        console.log("==========")
        return
    }

    stitchOneSecondVids(srcFilePaths)

    selectedElements = []
    console.log("==========")

});

document.addEventListener("click", function (event) {
    const timelineSection = document.querySelector("section.initial-timeline");
    if (!timelineSection.contains(event.target)) {
        selectedElements.forEach(function (element) {
            element.style.backgroundColor = COLOR_AQUA;
        });
        selectedElements = [];
    }
});

function stitchOneSecondVids(arr) {
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/stitchOneSecondVideos", true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify(arr));
}