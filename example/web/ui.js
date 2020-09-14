document.title = window.gen.app.name;

window.onload = function () {
    function App() {
        return React.createElement('div', null,
            window.gen.ui.MOTD(),
            window.gen.ui.Title(),
            window.gen.ui.SendMessageLog(),
            window.gen.ui.SendMessageInput()
        )
    }

    ReactDOM.render(e(App), document.getElementById('root'));
};
