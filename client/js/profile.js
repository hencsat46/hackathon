if (localStorage.getItem("token") == "") {
    window.location.replace("http://localhost:5000/auth/")
}