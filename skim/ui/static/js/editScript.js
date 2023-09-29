function editScript() {
    const fsVidProp = editDialog.getAttribute('fsVid-prop')
    const fsVid = JSON.parse(fsVidProp)
    fsVid.script = document.getElementById("new-script").value

    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/editFSVidScript", true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify(fsVid));

    editDialog.close()
}
