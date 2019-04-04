import QtQuick 2.6
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3
import "content"

Item {
    id: mainWindow
    visible: true

    //name = fileName (notice that file should start with capital letter to be visible)
    Hostname {
        id: hostnameLocal
        visible: true
        //there we decide what to do, when Hostname emit found signal
        //and yes, it's on+signalname
        onFound: {
            //ok is signal's argument
            if (ok) {
                hostnameLocal.errText.text = "found"
                hostnameLocal.visible = false
                loginLocal.visible = true
            } else {
                hostnameLocal.errText.text = "not found"
            }
        }
    }

    Login {
        id: loginLocal
        visible: false
        onLoggedIn: {
            if (ok) {
                goHelper.metadata()
                loginLocal.visible = false
                explorerLocal.visible = true

            } else {
                loginLocal.errText.text = "authorization failed"
            }
        }
    }

    Explorer {
        id: explorerLocal
        visible: false
    }
}
