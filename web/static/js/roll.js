var lucky;

function randomChoice(arr) {
    return arr[Math.floor(Math.random() * arr.length)]
}

function startRoll(nameArray) {
    var displayName = document.getElementById("content");
    lucky = setInterval(function() {
        displayName.innerHTML = randomChoice(nameArray)["Name"];
    }, 50);
}

function stopRoll() {
    clearInterval(lucky);
}