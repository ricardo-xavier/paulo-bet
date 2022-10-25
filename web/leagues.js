function leagues(user, token) {
    console.log("leagues")
    console.log(user)
    console.log(token)
    var loginDiv = document.getElementById("login_div");
    loginDiv.style = "display: none";
    var leaguesDiv = document.getElementById("leagues_div");
    leaguesDiv.style = "display: block";

    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        console.log("leagues response");
        console.log(this.readyState);
        console.log(this.status);
        console.log(this.responseText);
        if (this.readyState == 4) {
            if (this.status == 200) {
                var response = JSON.parse(this.responseText)
                response.leagues.forEach(function(league) {
                    addLeague(user, league, token);
                });
            } else {
                alert(this.status + " " + this.responseText)
            }
        }
    }
    xhttp.open("GET", "https://bocykia5x4.execute-api.sa-east-1.amazonaws.com/production/leagues/"
            + user + "?token=" + token);
    xhttp.send();
}

function addLeague(user, league, token) {
    const row = document.createElement("div");
    row.classList.add("row");
    const button = document.createElement("button");
    button.classList.add("btn");
    button.classList.add("btn-primary");
    button.classList.add("btn-block");
    button.innerHTML = league.leagueId;
    button.onclick = function() {
        ranking(user, league.leagueId, token);
    };
    row.appendChild(button);
    var leagues = document.getElementById("leagues");
    leagues.appendChild(row);

    var separator = document.createElement("div");
    separator.classList.add("row");
    var br = document.createElement("br");
    row.appendChild(br);
    leagues.appendChild(separator);
}
