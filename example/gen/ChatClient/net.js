// ChatClient Networking

socket = io('http://localhost:3001');
setTimeout(() => socket.emit('Connect'), 100)

window.gen = window.gen || {};
window.gen.net = {};

window.gen.net.onMOTD = function onMOTD(callback) {
    socket.on('MOTD', callback)
}

window.gen.net.onSendMessage = function onSendMessage(callback) {
    socket.on('SendMessage', callback)
}

window.gen.net.emitSendMessage = function emitSendMessage(input) {
    socket.emit('SendMessage', {userId, input, timestamp: Date.now()}) 
}