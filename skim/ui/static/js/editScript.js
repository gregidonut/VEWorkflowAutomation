function editScript() {
    const fsVidProp = editDialog.getAttribute('fsVid-prop')
    const fsVid = JSON.parse(fsVidProp)
    fsVid.script = document.getElementById("new-script").value

    fetch("/editFSVidScript", {
        method: "POST", headers: {
            "Content-Type": "application/json"
        }, body: JSON.stringify(fsVid)
    })
        .then(response => {
            if (!response.ok) {
                console.error("Error: " + response.status);
            }
        })
        .catch(error => {
            // Handle network or fetch error here
            console.error("Fetch Error: " + error);
        });

    console.log("script successfully edited")
    editDialog.close()
}
