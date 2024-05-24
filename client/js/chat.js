if (localStorage.getItem("token") == "") {
    window.location.replace("http://localhost:5000/auth/")
}

async function getMessages() {
    const chatroom_id = localStorage.getItem("chatroom_id")

    const request = new Request(`http://localhost:3000/message/${chatroom_id}`, {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
    })

    const response = await (await fetch(request)).json()

    console.log(response)

    ///
    const msgArray = response.content
    for (let i = 0; i < msgArray.length; i++) {
        fetchMessages(msgArray[i].Content, msgArray[i].sender_guid, msgArray[i].SenderName)
    }
    ///

    
    
}

getMessages()

const url = `ws://localhost:3000/ws/${localStorage.getItem("guid")}/${localStorage.getItem("chatroom_id")}`

const ws = new WebSocket(url)


ws.onopen = () => {
    console.log("hui")
}

function sendMessage() {
    const message = document.querySelector(".msger-input")
    const text = message.value
    const data = {
        "chatroom_id": localStorage.getItem("chatroom_id"),
        "sender_guid": localStorage.getItem("guid"),
        "content": text,
        "image": false,
    }
    message.value = ""
    drawMessage(text, "Я", "right")


    ws.send(JSON.stringify(data))
    scroll()
}

ws.onmessage = (event) => {
    const msg = JSON.parse(event.data)
    console.log(msg)
    const message = msg.content
    const senderName = msg.sender_name

    drawMessage(message, senderName, "left")
    scroll()
}

function drawMessage(message, name, path) {
    const msg = document.createElement("div")
    msg.classList.add("msg")
    msg.classList.add(path+"-msg")
    const msgInner = `
          
                <div class="msg-bubble">
                  <div class="msg-info">
                    <div class="msg-info-name">${name}</div>
                    <div class="msg-info-time"></div>
                  </div>
          
                  <div class="msg-text">
                    ${message}
                  </div>
                </div>
    `
    msg.innerHTML = msgInner
    const chat = document.querySelector("main.msger-chat")
    chat.append(msg)
    
}

function fetchMessages(message, guid, name) {
    let sender
    const msg = document.createElement("div")
    msg.classList.add("msg")
    if (guid == localStorage.getItem("guid")) {
        sender = "right"
        name = "Я"
    } else {
        sender = "left"
    }
    msg.classList.add(sender + "-msg")
    const msgInner = `
                
          
                <div class="msg-bubble">
                  <div class="msg-info">
                    <div class="msg-info-name">${name}</div>
                    <div class="msg-info-time"></div>
                  </div>
          
                  <div class="msg-text">
                    ${message}
                  </div>
                </div>
    `
    msg.innerHTML = msgInner
    const chat = document.querySelector("main.msger-chat")
    chat.append(msg)
    scroll()
}

function scroll() {
    const element = document.querySelector(".msger-chat");
    element.scrollTop = element.scrollHeight - element.clientHeight;
}

window.addEventListener("unload", async () => {
    const request = new Request(`http://localhost:3000/user/exitChatroom/${localStorage.getItem("chatroom_id")}/${localStorage.getItem("guid")}`, {
        method: "GET",
        mode: "cors",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`
        }
    })

    const response = await (await fetch(request)).json()

})