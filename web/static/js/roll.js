var lucky;
var goStatus = true;

function randomChoice(arr) {
    return arr[Math.floor(Math.random() * arr.length)]
}

function focusGo() {
    document.getElementById('go').focus();
}

function enteyPress() {
    var goButton = document.getElementById("go");
    if (event.keyCode == 13) {
        goButton.click();
    }
}

function roll(nameArray) {
    if (goStatus) {
        startRoll(nameArray);
    } else {
        stopRoll();
    }
    goStatus = !goStatus;
}

function startRoll(nameArray) {
    var displayName = document.getElementById("content");
    var goButton = document.getElementById("go");
    goButton.innerHTML = "停止";
    lucky = setInterval(function() {
        displayName.innerHTML = randomChoice(nameArray)["Name"];
    }, 50);

}

function stopRoll() {
    var goButton = document.getElementById("go");
    goButton.innerHTML = "开始";
    clearInterval(lucky);
}

function addOnloadEvent(func) {
    var oldonload = window.onload;
    if (typeof window.onload != 'function') {
        window.onload = func;
    } else {
        window.onload = function() {
            oldonload();
            func();
        }
    }
}

addOnloadEvent(focusGo);