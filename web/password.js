function changePassword(login, token) {
    var leaguesDiv = document.getElementById("leagues_div");
    leaguesDiv.style = "display: none";
    var passwordDiv = document.getElementById("password_div");
    passwordDiv.style = "display: block; margin: 25px 25px";
}

function confirmPassword() {
    var password1 = document.getElementById("password_1").value;
    var password2 = document.getElementById("password_2").value;
    if (!password1 || !password2) {
        alert("senha nao informada");
        if (!password1) {
            document.getElementById("password_1").focus();
        } else {
            document.getElementById("password_2").focus();
        }
        return;
    }
    if (password1 != password2) {
        alert("as senhas estao diferentes");
        document.getElementById("password_1").focus();
        return;
    }

    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        console.log("password response");
        console.log(this.readyState);
        console.log(this.status);
        console.log(this.responseText);
        if (this.readyState == 4) {
            if (this.status == 204) {
                alert("Senha alterada com sucesso!");
                hide();
            } else {
                alert("Erro " + this.status + " " + this.responseText)
            }
        }
    }
    xhttp.open("PATCH", "https://bocykia5x4.execute-api.sa-east-1.amazonaws.com/production/login");
    xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    request = JSON.stringify({ 
        "token": document.getElementById("token_hidden").value,
        "login": document.getElementById("user_hidden").value,
        "password": document.getElementById("password_1").value });
console.log(request);
    xhttp.send(request)
}

function cancelPassword() {
    hide();
}

function hide() {
    var leaguesDiv = document.getElementById("leagues_div");
    leaguesDiv.style = "display: block; margin: 25px 25px";
    var passwordDiv = document.getElementById("password_div");
    passwordDiv.style = "display: none";
}
