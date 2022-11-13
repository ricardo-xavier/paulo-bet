var bet;
var league;
var login;
var token;

function updateBet(pBet, pLeague, pLogin, pToken) {
    var dict = {
        "ARG": "Argentina",
        "AUS": "Austrália",
        "GAL": "Gales",
        "DEN": "Dinamarca",
        "ECU": "Ecuador",
        "ENG": "Inglaterra",
        "FRA": "França",
        "HOL": "Holanda",
        "IRN": "Irã",
        "KSA": "Arábia Saudita",
        "MEX": "México",
        "POL": "Polônia",
        "QAT": "Catar",
        "SEN": "Senegal",
        "TUN": "Tunísia",
        "USA": "Estados Unidos"
    };

    console.log("bets")
    bet = pBet;
    league = pLeague;
    login = pLogin;
    token = pToken;
    console.log(bet)
    console.log(league)
    console.log(login)
    console.log(token)
    var betsDiv = document.getElementById("bets_div");
    betsDiv.style = "display: none";
    var betDiv = document.getElementById("bet_div");
    betDiv.style = "display: block";

    var homeLbl = document.getElementById("home_lbl");
    homeLbl.innerHTML = bet.matchId.split("-")[0] + "-" + dict[bet.matchId.split("-")[0]];

    var visitorsLbl = document.getElementById("visitors_lbl");
    visitorsLbl.innerHTML = bet.matchId.split("-")[1] + "-" + dict[bet.matchId.split("-")[1]];

    var home = document.getElementById("home");
    home.value = bet.home;

    var visitors = document.getElementById("visitors");
    visitors.value = bet.visitors;

}

function confirm() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        console.log("bet response");
        console.log(this.readyState);
        console.log(this.status);
        console.log(this.responseText);
        if (this.readyState == 4) {
            if (this.status == 204) {
                alert("Aposta executada com sucesso!");
            } else {
                alert("Erro " + this.status + " " + this.responseText)
            }
            var betDiv = document.getElementById("bet_div");
            betDiv.style = "display: none";
            bets(login, league, login, token);
        }
    }
    url = "https://bocykia5x4.execute-api.sa-east-1.amazonaws.com/production/bets/"
            + league + "/" + login;
    console.log(url);
    xhttp.open("POST", url);
    xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    var home = parseInt(document.getElementById("home").value);
    var visitors = parseInt(document.getElementById("visitors").value);
    request = JSON.stringify({ 
        "token": token,
        "matchId": bet.matchId,
        "home": home,
        "visitors": visitors });
    xhttp.send(request)
}

function cancel() {
    var betDiv = document.getElementById("bet_div");
    betDiv.style = "display: none";
    bets(login, league, login, token);
}
