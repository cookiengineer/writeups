// Call a Go function from JavaScript. The function printSomething()
// in main.go is bound to the JavaScript symbol "ps".
function sayHello() {
    let name = document.getElementById("yourName").value
    ps(name)
}

// This is a JavaScript function that we will call from Go
// in main.go.
function setDivContent(divId, content) {
    document.getElementById(divId).innerHTML = content
}
