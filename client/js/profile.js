if (localStorage.getItem("token") == "") {
    window.location.replace("http://localhost:5000/auth/")
}

var usernameEdited = false
var passwordEdited = false
var emailEdited = false

function changeUsername() {
    if (usernameEdited) {
        
    } else {

    }
}

function setUsername() {
    document.getElementById("username").innerText = localStorage.getItem("username")
}

setUsername()

async function updateUsername() {
    let newUsername = document.getElementById("newUsername").value
    data = {
        guid: localStorage.getItem("guid"),
        username: newUsername
    }

    const request = new Request('http://localhost:3000/user/updateUsername', {
        method: 'PUT',
        mode: 'cors',
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
        body: JSON.stringify(data)
    })

    const response = await (await fetch(request)).json()
    if (response.error == "") {
        localStorage.setItem("username", newUsername)
        window.location.reload()
    }
}

async function updatePassword() {
    let oldPassword = document.getElementById("oldPassword").value
    let newPassword = document.getElementById("newPassword").value

    data = {
        guid: localStorage.getItem("guid"),
        old_password: oldPassword,
        password: newPassword
    }

    const request = new Request('http://localhost:3000/user/updatePassword', {
        method: 'PUT',
        mode: 'cors',
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
        body: JSON.stringify(data)
    })

    const response = await (await fetch(request)).json()
    console.log(response)
    if (response.error == "") {
        window.location.reload()
    } else {
        alert('provided old password is incorrect')
    }
}

async function logout() {
    localStorage.clear()
    window.location.replace('http://localhost:5000/auth/')
}

async function deleteAcc() {
    const request = new Request(`http://localhost:3000/user/delete/${localStorage.getItem("guid")}`, {
        method: 'DELETE',
        mode: 'cors',
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
    })

    const response = await (await fetch(request)).json()
    console.log(response)
    if (response.error == "") {
        await logout()
    } else {
        alert(response)
    }   
}