import QtQuick 2.6
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3


Item {
    id: loginWindow
    anchors.fill: parent


    Column {
        Row {
            ListView {
                id: listView
                model: FilesModel
                delegate: Text {
                    text: fileName + "  | " + mTime
                }
                height: count*20
            }

            Connections {
                     target: goHelper
                     onChanged: { listView.model = 0; listView.model = FilesModel }
                 }
        }
    }
}