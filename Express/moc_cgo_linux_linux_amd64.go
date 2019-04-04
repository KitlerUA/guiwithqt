package main

/*
#cgo CFLAGS: -pipe -O2 -Wall -W -D_REENTRANT -fPIC -DQT_NO_DEBUG -DQT_MULTIMEDIA_LIB -DQT_DESIGNER_LIB -DQT_UIPLUGIN_LIB -DQT_WIDGETS_LIB -DQT_QUICK_LIB -DQT_GUI_LIB -DQT_QML_LIB -DQT_NETWORK_LIB -DQT_DBUS_LIB -DQT_XML_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -pipe -O2 -Wall -W -D_REENTRANT -fPIC -DQT_NO_DEBUG -DQT_MULTIMEDIA_LIB -DQT_DESIGNER_LIB -DQT_UIPLUGIN_LIB -DQT_WIDGETS_LIB -DQT_QUICK_LIB -DQT_GUI_LIB -DQT_QML_LIB -DQT_NETWORK_LIB -DQT_DBUS_LIB -DQT_XML_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../QuatrixGExpress -I. -I/home/vkit/Qt/5.12.0/gcc_64/include -I/home/vkit/Qt/5.12.0/gcc_64/include/QtMultimedia -I/home/vkit/Qt/5.12.0/gcc_64/include/QtDesigner -I/home/vkit/Qt/5.12.0/gcc_64/include/QtUiPlugin -I/home/vkit/Qt/5.12.0/gcc_64/include/QtWidgets -I/home/vkit/Qt/5.12.0/gcc_64/include/QtQuick -I/home/vkit/Qt/5.12.0/gcc_64/include/QtGui -I/home/vkit/Qt/5.12.0/gcc_64/include/QtQml -I/home/vkit/Qt/5.12.0/gcc_64/include/QtNetwork -I/home/vkit/Qt/5.12.0/gcc_64/include/QtDBus -I/home/vkit/Qt/5.12.0/gcc_64/include/QtXml -I/home/vkit/Qt/5.12.0/gcc_64/include/QtCore -I. -isystem /usr/include/libdrm -I/home/vkit/Qt/5.12.0/gcc_64/mkspecs/linux-g++
#cgo LDFLAGS: -O1 -Wl,-rpath,/home/vkit/Qt/5.12.0/gcc_64/lib
#cgo LDFLAGS:  -L/home/vkit/Qt/5.12.0/gcc_64/lib -lQt5Multimedia -lQt5Designer -lQt5Widgets -lQt5Quick -lQt5Gui -lQt5Qml -lQt5Network -lQt5DBus -lQt5Xml -lQt5Core -lGL -lpthread
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"
