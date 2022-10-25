var login;
var league;
var token;

function ranking(pLogin, pLeague, pToken) {
    console.log("ranking")
    login = pLogin;
    league = pLeague;
    token = pToken;
    console.log(login)
    console.log(league)
    console.log(token)
    var leaguesDiv = document.getElementById("leagues_div");
    leaguesDiv.style = "display: none";
    var rankingDiv = document.getElementById("ranking_div");
    rankingDiv.style = "display: block";

    removeChildren("ranking");

    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        console.log("ranking response");
        console.log(this.readyState);
        console.log(this.status);
        console.log(this.responseText);
        if (this.readyState == 4) {
            if (this.status == 200) {
                var response = JSON.parse(this.responseText)
                response.ranking.forEach(function(ranking) {
                    addRanking(ranking);
                });
            } else {
                alert("Erro " + this.status + " " + this.responseText)
            }
        }
    }
    url = "https://bocykia5x4.execute-api.sa-east-1.amazonaws.com/production/ranking/"
            + league + "?token=" + token + "&login=" + login;
    xhttp.open("GET", url);
    xhttp.send();
}

function addRanking(ranking) {
    var row = document.createElement("div");
    row.classList.add("row");

    var userCol = document.createElement("div");
    userCol.classList.add("col-xs-6");
    var userP = document.createElement("p");
    userP.innerHTML = ranking.userId;
    userCol.appendChild(userP);
    row.appendChild(userCol);

    var scoreCol = document.createElement("div");
    scoreCol.classList.add("col-xs");
    var scoreBtn = document.createElement("button");
    scoreBtn.classList.add("btn");
    scoreBtn.classList.add("btn-primary");
    scoreBtn.innerHTML = ranking.score;
    scoreBtn.onclick = function() {
        bets(login, league, ranking.userId, token);
    };
    scoreCol.appendChild(scoreBtn);
    row.appendChild(scoreCol);

    var rankingDiv = document.getElementById("ranking");
    rankingDiv.appendChild(row);

    var separator = document.createElement("div");
    separator.classList.add("row");
    var hr = document.createElement("hr");
    separator.appendChild(hr);
    rankingDiv.appendChild(separator);
}

function backToLeagues() {
    var rankingDiv = document.getElementById("ranking_div");
    rankingDiv.style = "display: none";
    leagues(login, token);
}
