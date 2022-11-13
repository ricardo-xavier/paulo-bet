function leagues(login, token) {
    console.log("leagues")
    console.log(login)
    console.log(token)
    var loginDiv = document.getElementById("login_div");
    loginDiv.style = "display: none";
    var leaguesDiv = document.getElementById("leagues_div");
    leaguesDiv.style = "display: block; margin: 25px 25px";

    removeChildren("leagues");

    var userHidden = document.getElementById("user_hidden");
    userHidden.value = login;
    var tokenHidden = document.getElementById("token_hidden");
    tokenHidden.value = token;

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
                    addLeague(login, league, token);
                });
            } else {
                alert("Erro " + this.status + " " + this.responseText)
            }
        }
    }
    xhttp.open("GET", "https://bocykia5x4.execute-api.sa-east-1.amazonaws.com/production/leagues/"
            + login + "?token=" + token);
    xhttp.send();
}

function addLeague(login, league, token) {
    const row = document.createElement("div");
    row.classList.add("row");
    const button = document.createElement("button");
    button.classList.add("btn");
    button.classList.add("btn-primary");
    button.classList.add("btn-block");
    button.innerHTML = league.leagueId;
    button.onclick = function() {
        ranking(login, league.leagueId, token);
    };
    row.appendChild(button);
    var leagues = document.getElementById("leagues");
    leagues.appendChild(row);

    var separator = document.createElement("div");
    separator.classList.add("row");
    var hr = document.createElement("hr");
    row.appendChild(hr);
    leagues.appendChild(separator);
}
