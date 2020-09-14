// ChatClient UI

window.gen = window.gen || {};
window.gen.ui = {};

const e = React.createElement;
const $ = q => document.querySelector(q);
const net = window.gen.net;
const userId = window.gen.app.userId;

window.gen.ui.Title = function Title() {
    return e('h1', null, window.gen.app.name);
};

window.gen.ui.MOTD = function MOTD() {
    const [val, setVal] = React.useState();
    React.useEffect(() => net.onMOTD(setVal));
    return e('div', null, `MOTD: ${JSON.stringify(val)}`);
}

window.gen.ui.MOTDLog = function MOTDLog() {
    const [vals, setVals] = React.useState([]);
    React.useEffect(() => net.onMOTD(val => setVals(vals.concat([val]))));
    return e('ul', null, vals.map(({userId, input, timestamp}) =>
        e('li', {key: timestamp}, `${userId}: ${input}`)
    ));
}

window.gen.ui.SendMessage = function SendMessage() {
    const [val, setVal] = React.useState();
    React.useEffect(() => net.onSendMessage(setVal));
    return e('div', null, `SendMessage: ${JSON.stringify(val)}`);
}

window.gen.ui.SendMessageLog = function SendMessageLog() {
    const [vals, setVals] = React.useState([]);
    React.useEffect(() => net.onSendMessage(val => setVals(vals.concat([val]))));
    return e('ul', null, vals.map(({userId, input, timestamp}) =>
        e('li', {key: timestamp}, `${userId}: ${input}`)
    ));
}

window.gen.ui.SendMessageInput = function SendMessageInput() {
    let id = 'SendMessage_input';
    const SendMessage = () => {
        const input = $('#'+id);
        input.value && net.emitSendMessage(input.value);
        input.value = '';
    };
    return e('div', null,
        e('textarea', {id}),
        e('button', {onClick: SendMessage}, 'SendMessage'),
    );
};