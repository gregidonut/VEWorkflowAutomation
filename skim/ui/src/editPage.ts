export default function EditPage() {
    // Get the target element to observe
    const targetElement = <HTMLElement>document.querySelector("figure.spinner");
    targetElement.style.backgroundColor = "#b99a9a"
    //
    // // Create an Intersection Observer
    // const observer = new IntersectionObserver((entries) => {
    //     // Check if the target element is in the viewport
    //     if (entries[0].isIntersecting) {
    //         console.log("The spinner is now in the viewport");
    //     }
    // });
    //
    // // Start observing the target element
    // observer.observe(targetElement);
}
