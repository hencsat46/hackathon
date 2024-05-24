async function send() {
    const username = document.querySelector('.username-text').value
    const password = document.querySelector('.password-text').value

    dataObject = {
        username: username,
        password: password,
    }
    
    const request = new Request('http://localhost:3000/user/login', {
        method: 'POST',
        mode: 'cors',
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(dataObject)
    })
    
    const response = await (await fetch(request)).json()
    console.log(response)
    if (response.error == "") {
        localStorage.setItem("token", response.content.token)
        localStorage.setItem("guid", response.content.guid)
        window.location.replace("http://localhost:5000/")
    }

}

function moveToSignup() {
    const signupHtml = `
    <div class="username">
            <div class="exists">Такой пользователь существует</div>
            <div class="username-header">Логин</div>
            <input type="text" class="username-text">
        </div>
        <div class="password">
            <div class="password-header">Пароль</div>
            <input type="password" class="password-text">
        </div>
        <div class="password">
            <div class="error">Пароли не совпадают</div>
            <div class="password-header">Повторите пароль</div>
            <input type="password" class="password-text">
        </div>
        <div class="button-section">
        <button class="submit" onclick="signup()">Зарегистрироваться</button> 
        </div>
        <div class="registration">
        <a class="registration-link" onclick="moveToLogin()">Войти</a>
    </div>
    `

    const oldLogin = document.querySelector('div.login-section')

    oldLogin.innerHTML = signupHtml

}

function moveToLogin() {
    const loginHtml = `
    <div class="username">
        <div class="username-header">Логин</div>
        <input type="text" class="username-text">
    </div>
    <div class="password">
        <div class="password-invalid">Неверный пароль</div>
        <div class="password-header">Пароль</div>
        <input type="password" class="password-text">
    </div>
    <div class="button-section">
                    <button class="submit" onclick="send()">Войти</button>
                </div>
    <div class="registration">
        <a class="registration-link" onclick="moveToSignup()">Зарегистрироваться</a>
    </div>
    
    `

    const oldSignUp = document.querySelector(".login-section")

    oldSignUp.innerHTML = loginHtml
}

async function signup() {
    const username = document.querySelector('.username-text').value
    const passwords = document.querySelectorAll('.password-text')
    const error = document.querySelector('.error')

    const data = {
        username: username,
        password: passwords[0].value,
    }

    if (passwords[0].value != passwords[1].value) {
        error.style.display = 'block'
        return
    } else {
        error.style.display = 'none'
    }

    console.log(data)

    const request = new Request('http://localhost:3000/user/create/', {
        method: 'POST',
        mode: 'cors',
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data)
    })

    const response = await (await fetch(request)).json()
    console.log(response.content)
    localStorage.setItem("guid", response.content)
    console.log(response.error)
    if (response.error == "") {
        console.log("hello")
        document.querySelector('.exists').style.display = 'none'
        const request = new Request('http://localhost:3000/user/login/', {
            method: "POST",
            mode: 'cors',
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        })
        const response = await (await fetch(request)).json()
        console.log(response)
        
        if (response.error == "") {
            localStorage.setItem("token", response.content.token)
            window.location.replace('http://localhost:5000/')
        }
        
    } else {
        if (response.status == 400) {
            document.querySelector('.exists').style.display = 'block'
        }
    }

}