function deleteScript() {
    console.log("running delete script...")

    const fsVidProp = editDialog.getAttribute('fsVid-prop')
    const fsVid = JSON.parse(fsVidProp)

    fetch("/deleteFSVid", {
        method: "POST", headers: {
            "Content-Type": "applicationOld/json"
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

    deleteDialog.close()
    console.log("delete script ran successfully")
    return false;
}
