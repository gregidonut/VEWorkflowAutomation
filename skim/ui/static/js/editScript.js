function editScript() {
    console.log("logging fsVid-prop:")
    console.log(editDialog.getAttribute('fsVid-prop'))

    const fsVidProp = editDialog.getAttribute('fsVid-prop')
    const fsVid = JSON.parse(fsVidProp)
    fsVid.script = document.getElementById("new-script").value

    console.log("logging fsvid.script:")
    console.log(fsVid.script)

    editDialog.close()
}