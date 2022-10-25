function login() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        console.log("login response");
        console.log(this.readyState);
        console.log(this.status);
        console.log(this.responseText);
        if (this.readyState == 4) {
            if (this.status == 200) {
                var login = document.getElementById("user").value;
                document.cookie = "user=" + login +  "; expires=Fri, 31 Dec 9999 23:59:59 GMT; path=/";
                var response = JSON.parse(this.responseText)
                leagues(login, response.token);
            } else {
                alert("Erro " + this.status + " " + this.responseText)
            }
        }
    }
    xhttp.open("POST", "https://bocykia5x4.execute-api.sa-east-1.amazonaws.com/production/login");
    xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    request = JSON.stringify({ 
        "login": document.getElementById("user").value, 
        "password": document.getElementById("password").value });
    xhttp.send(request)
}
