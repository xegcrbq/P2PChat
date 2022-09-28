class MySocket {
    constructor() {
        this.mysocket = null;

        this.messageHistory = document.querySelector('.messageHistory');

    }
    showMessage(text){
        const jtext = JSON.parse(text)
        jtext.forEach((element) => {
            this.addMessage(element['MessageText'],element['SenderId'])
        })
    }
    send(){
        this.mysocket.send(inputField.value)
        inputField.value = ""
    }
    connectSocket(){
        const socket = new WebSocket("ws://localhost:8080/api/socket/" + this.getCookie('session_id'));
        this.mysocket = socket;

        socket.onmessage = (e)=>{
            this.showMessage(e.data);
        }
        socket.onopen =  ()=> {
            console.log("socket opened")
        };
        socket.onclose = ()=>{
            console.log("socket close")
        }
    }
    getCookie(cookieName) {
        let cookie = {};
        document.cookie.split(';').forEach(function(el) {
            let [key,value] = el.split('=');
            cookie[key.trim()] = value;
        })
        return cookie[cookieName];
    }
    addMessage(text = '', author = 0) {
        const newMessage = document.createElement('div')
        newMessage.classList.add('message')
        if (author != 1) {
            newMessage.classList.add('green')
        }
        newMessage.innerText = text
        const messageContainer = document.createElement('div')
        messageContainer.classList.add('messageContainer')
        messageContainer.appendChild(newMessage)
        this.messageHistory.appendChild(messageContainer)
    }
}