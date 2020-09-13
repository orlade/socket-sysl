// ChatClient

socket = io('http://localhost:3001');
setTimeout(() => socket.emit('Connect'), 100)

socket.on('MOTD', msg => {
    const li = document.createElement('li');
    li.innerText = JSON.stringify(msg, null, 2);
    document.getElementById('MOTD').append(li);
})
socket.on('SendMessage', msg => {
    const li = document.createElement('li');
    li.innerText = JSON.stringify(msg, null, 2);
    document.getElementById('SendMessage').append(li);
})

function emitSendMessage() {
    const input = document.getElementById('SendMessage_input').value;
    socket.emit('SendMessage', {userId, input}) 
}