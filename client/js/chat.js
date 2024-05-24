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
}

getMessages()