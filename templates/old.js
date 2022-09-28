// const inputField = document.querySelector('.inputField');
// const inputButton = document.querySelector('.inputButton');
// const inputFile = document.querySelector('.inputFile');
// const messageHistory = document.querySelector('.messageHistory');
// inputField.addEventListener('keypress', function (e) {
//     if (e.key === 'Enter') {
//         sendMessage()
//
//     }
// });
// inputButton.addEventListener('click', function (e) {
//     sendMessage()
//
// });
// async function postData(url = '', data = {}) {
//     // Default options are marked with *
//     const response = await fetch(url, {
//         method: 'POST', // *GET, POST, PUT, DELETE, etc.
//         mode: 'cors', // no-cors, *cors, same-origin
//         cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
//         credentials: 'same-origin', // include, *same-origin, omit
//         headers: {
//             'Content-Type': 'application/json'
//         },
//         cookie: document.cookie,
//         redirect: 'follow', // manual, *follow, error
//         referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
//         body: JSON.stringify(data) // body data type must match "Content-Type" header
//     });
//     return response; // parses JSON response into native JavaScript objects
// }
// async function postFile(url = '', data = {}) {
//     // Default options are marked with *
//     const response = await fetch(url, {
//         method: 'POST', // *GET, POST, PUT, DELETE, etc.
//         mode: 'cors', // no-cors, *cors, same-origin
//         cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
//         credentials: 'same-origin', // include, *same-origin, omit
//         headers: {
//             'Content-Type': 'file'
//         },
//         redirect: 'follow', // manual, *follow, error
//         referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
//         body: data // body data type must match "Content-Type" header
//     });
//     return response; // parses JSON response into native JavaScript objects
// }
// function addMessage(text = '', author = 0) {
//     const newMessage = document.createElement('div')
//     newMessage.classList.add('message')
//     if (author != 1) {
//         newMessage.classList.add('green')
//     }
//     newMessage.innerText = text
//     const messageContainer = document.createElement('div')
//     messageContainer.classList.add('messageContainer')
//     messageContainer.appendChild(newMessage)
//     messageHistory.appendChild(messageContainer)
// }
// function sendMessage(){
//     if (inputField.value != ""){
//         addMessage(inputField.value)
//         postData('send/', {Message:inputField.value})
//         inputField.value = ""
//     }
//
//     if (inputFile.files.length != 0){
//         let formData = new FormData();
//
//         formData.append("photo", inputFile.files[0]);
//         postFile('uploadFile/', formData);
//     }
// }
// update()
// setInterval(update, 5000);
// async function update () {
//     const resp = (await postData('update/v2/', {
//         MessageCount: document.querySelectorAll('.messageContainer').length,
//     }).then(r => r.json()));
//     resp.forEach((element) => {
//         addMessage(element['MessageText'],element['SenderId'])
//     })
// }