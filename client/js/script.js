if (localStorage.getItem("token") == "") {
    window.location.replace("http://localhost:5000/auth/")
}

async function fetchChatrooms() {

    const request = new Request(`http://localhost:3000/chatroom/get/`, {
        method: "GET",
        mode: "cors",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`
        }
    })
    
    const response = await (await fetch(request)).json()

    if (response.error == "") {
        const chatrooms = response.content
        const section = document.querySelector(".chatrooms-section")
        for (let i = 0; i < chatrooms.length; i++) {
            const chat = document.createElement("div")
            chat.classList.add("chatroom")
            chat.setAttribute("id", chatrooms[i].ChatroomId)
            chat.setAttribute("onclick", "moveToChat(this)")
            chat.innerHTML = `<div class="chatroom-name">${chatrooms[i].Name}</div>`  
            section.append(chat)
        }
    }
}

async function moveToChat(element) {
    const chatId = element.getAttribute("id")

    localStorage.setItem("chatroom_id", chatId)

    const request = new Request(`http://localhost:3000/user/enterChatroom/${chatId}/${localStorage.getItem("guid")}`, {
        method: "GET",
        mode: "cors",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`
        }
    })

    const response = await (await fetch(request)).json()
    console.log(response)

    window.location.replace("http://localhost:5000/chat")
}

fetchChatrooms()
var creation = false
function addChatroom() {
    const createButton = document.querySelector(".create")
    const addChat = document.querySelector(".creation")
    if (!creation) {
        createButton.innerText = "Убрать"
        addChat.style.display = "flex"
        creation = !creation
    } else {
        createButton.innerText = "Добавить чат"
        addChat.style.display = "none"
        creation = !creation
    }
}

async function addChat() {
    const chatName = document.querySelector(".name-input").value

    const guid = localStorage.getItem("guid")

    const data = {
        "name": chatName,
        "guid": guid,
        
    }

    console.log(JSON.stringify(data))

    const request = new Request("http://localhost:3000/chatroom/create", {
        method: "POST",
        mode: "cors",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    })

    const response = await (await fetch(request)).json()

    console.log(response)
}