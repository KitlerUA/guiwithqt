import QtQuick 2.6
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3


Item {
    id: loginWindow
    anchors.fill: parent

    property Text errText: errorText
    signal loggedIn(bool ok)

    Column {
        anchors.centerIn: parent
        Row {
            Text {
                id: errorText
                text: ""
            }
        }

        Row {
            TextField {
                id: inputLogin
                width: 270
                placeholderText: "email"
                selectByMouse: true
            }
        }

        Row {
            spacing: 10
            TextField {
                id: inputPassword
                width: 270
                echoMode: TextInput.Password
                placeholderText: "password"
                selectByMouse: true
            }
            Button {
                text: "login"
                onClicked: loginWindow.loggedIn(goLogin.login(inputLogin.text,inputPassword.text))
            }
        }
    }
}
