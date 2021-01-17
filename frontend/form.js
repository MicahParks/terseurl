async function submitForm(e) {

    e.preventDefault();

    const submitButton = document.getElementById("submitButton");
    submitButton.disabled = true;
    setTimeout(() => submitButton.disabled = false, 1000); // TODO

    let terse = new Terse();
    terse.originalURL = document.getElementById("originalURL").value;
    terse.shortenedURL = document.getElementById("shortenedURL").value;

    console.log(JSON.stringify(terse));
}