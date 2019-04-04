import QtQuick 2.6
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3


Item {
    id: hostnameWindow
    anchors.fill: parent

    property Text errText: errorText
    signal found(bool ok)

    Column {
        anchors.centerIn: parent
        Row {
            Text {
                id: errorText
                text: ""
            }
        }

        Row {
            spacing: 10
            TextField {
                id: inputHostname
                width: 270
                placeholderText: "hostname"
                selectByMouse: true
            }
            Button {
                text: "find"
                onClicked: hostnameWindow.found(goHostname.find(inputHostname.text))
            }
        }
    }
}


