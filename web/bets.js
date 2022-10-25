var login;
var league;
var user;
var token;

function bets(pLogin, pLeague, pUser, pToken) {
    console.log("bets")
    login = pLogin;
    league = pLeague;
    user = pUser;
    token = pToken;
    console.log(login)
    console.log(league)
    console.log(user)
    console.log(token)
    var rankingDiv = document.getElementById("ranking_div");
    rankingDiv.style = "display: none";
    var betsDiv = document.getElementById("bets_div");
    betsDiv.style = "display: block";

    removeChildren("bets");

    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        console.log("bets response");
        console.log(this.readyState);
        console.log(this.status);
        console.log(this.responseText);
        if (this.readyState == 4) {
            if (this.status == 200) {
                var response = JSON.parse(this.responseText)
                response.bets.forEach(function(bet) {
                    addBet(bet);
                });
            } else {
                alert("Erro " + this.status + " " + this.responseText)
            }
        }
    }
    url = "https://bocykia5x4.execute-api.sa-east-1.amazonaws.com/production/bets/"
            + league + "/" + user + "?token=" + token + "&login=" + login;
    console.log(url);
    xhttp.open("GET", url);
    xhttp.send();
}

function addBet(bet) {
    var row = document.createElement("div");
    row.classList.add("row");

    var matchCol = document.createElement("div");
    matchCol.classList.add("col-xs-4");
    var matchP = document.createElement("p");
    matchP.innerHTML = bet.date + " " + bet.matchId;
    matchCol.appendChild(matchP);
    row.appendChild(matchCol);

    var resultCol = document.createElement("div");
    resultCol.classList.add("col-xs-2");
    var resultBtn = document.createElement("button");
    resultBtn.classList.add("btn");
    resultBtn.innerHTML = bet.home + "x" + bet.visitors;
    if (bet.editable) {
        resultBtn.classList.add("btn-primary");
        resultBtn.onclick = function() {
            updateBet(bet, league, login, token);
        };
    } else {
        resultBtn.classList.add("btn-danger");
    }
    resultCol.appendChild(resultBtn);
    row.appendChild(resultCol);

    var scoreCol = document.createElement("div");
    scoreCol.classList.add("col-xs-2");
    var scoreBtn = document.createElement("button");
    scoreBtn.classList.add("btn");
    scoreBtn.innerHTML = bet.score;
    scoreBtn.classList.add("btn-success");
    scoreCol.appendChild(scoreBtn);
    row.appendChild(scoreCol);

    var betsDiv = document.getElementById("bets");
    betsDiv.appendChild(row);

    var separator = document.createElement("div");
    separator.classList.add("row");
    var hr = document.createElement("hr");
    separator.appendChild(hr);
    betsDiv.appendChild(separator);
}

function backToRanking() {
    var betsDiv = document.getElementById("bets_div");
    betsDiv.style = "display: none";
    ranking(login, league, token);
}
