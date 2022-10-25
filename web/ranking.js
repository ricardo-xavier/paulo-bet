function ranking(user, league, token) {
    console.log("ranking")
    console.log(user)
    console.log(league)
    console.log(token)
    var leaguesDiv = document.getElementById("leagues_div");
    leaguesDiv.style = "display: none";
    var rankingDiv = document.getElementById("ranking_div");
    rankingDiv.style = "display: block";

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
                alert(this.status + " " + this.responseText)
            }
        }
    }
    url = "https://bocykia5x4.execute-api.sa-east-1.amazonaws.com/production/ranking/"
            + league + "?token=" + token + "&login=" + user;
    console.log("ranking url");
    console.log("[" + league + "]");
    console.log("[" + token + "]");
    console.log("[" + user + "]");
    console.log(url);
    xhttp.open("GET", url);
    xhttp.send();
}

function addRanking(ranking) {
    var row = document.createElement("div");
    row.classList.add("row");

    var userCol = document.createElement("div");
    userCol.classList.add("col-xs-6");
    var userBtn = document.createElement("p");
    userBtn.innerHTML = ranking.userId;
    userCol.appendChild(userBtn);
    row.appendChild(userCol);

    var scoreCol = document.createElement("div");
    scoreCol.classList.add("col-xs");
    var scoreBtn = document.createElement("button");
    scoreBtn.classList.add("btn");
    scoreBtn.classList.add("btn-primary");
    scoreBtn.innerHTML = ranking.score;
    scoreBtn.onclick = function() {
console.log(ranking);
        alert(ranking.userId);
    };
    scoreCol.appendChild(scoreBtn);
    row.appendChild(scoreCol);

    var rankingDiv = document.getElementById("ranking");
    rankingDiv.appendChild(row);

    var separator = document.createElement("div");
    separator.classList.add("row");
    var br = document.createElement("br");
    separator.appendChild(br);
    rankingDiv.appendChild(separator);
}
