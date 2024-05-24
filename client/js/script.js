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
        console.log(chatrooms)
        const section = document.querySelector(".chatrooms-section")
        for (let i = 0; i < chatrooms.length; i++) {

            const chat = document.createElement("div")
            chat.classList.add("chatroom-wrapper")
            chat.setAttribute("id", chatrooms[i].ChatroomId)
            
            chat.innerHTML = `
            <div class="chatroom" onclick="moveToChat(this)">
                <div class="chatroom-name">${chatrooms[i].Name}</div>  
            </div>
            <button class="delete-button" onclick="deleteChatroom(this)">Удалить</button>
            `  
            section.append(chat)
        }
    }
}

async function deleteChatroom(element) {
    const cid = element.parentElement.getAttribute("id")

    const request = new Request(`http://localhost:3000/chatroom/${cid}/skadflj`, {
        method: "DELETE",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
        mode: "cors",
    })

    const response = await (await fetch(request)).json()

    console.log(response)

    if (response.error == "") {
        window.location.reload()
    }
} 

async function moveToChat(element) {
    const chatId = element.parentElement.getAttribute("id")
    console.log(chatId)
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
    window.location.reload()
}