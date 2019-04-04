import QtQuick 2.0
import QtQuick.Controls 2.0

Item {
  anchors.fill: parent
  id: rootItem

  Button {
    text: "Button"

    anchors {
      top: parent.top
      left: parent.left
      right: parent.right
    }
  }
}