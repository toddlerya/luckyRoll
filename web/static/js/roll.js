var lucky;

function randomChoice(arr) {
    return arr[Math.floor(Math.random() * arr.length)]
}

function startRoll(nameArray) {
    var displayName = document.getElementById("displayName");
    lucky = setInterval(function() {
        displayName.innerHTML = randomChoice(nameArray)["Name"];
    }, 100);
}

function stopRoll() {
    clearInterval(lucky);
}