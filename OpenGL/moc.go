package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
	std_gui "github.com/therecipe/qt/gui"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type OpenGLWindow_ITF interface {
	std_gui.QWindow_ITF
	OpenGLWindow_PTR() *OpenGLWindow
}

func (ptr *OpenGLWindow) OpenGLWindow_PTR() *OpenGLWindow {
	return ptr
}

func (ptr *OpenGLWindow) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWindow_PTR().Pointer()
	}
	return nil
}

func (ptr *OpenGLWindow) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWindow_PTR().SetPointer(p)
	}
}

func PointerFromOpenGLWindow(ptr OpenGLWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.OpenGLWindow_PTR().Pointer()
	}
	return nil
}

func NewOpenGLWindowFromPointer(ptr unsafe.Pointer) (n *OpenGLWindow) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(OpenGLWindow)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *OpenGLWindow:
			n = deduced

		case *std_gui.QWindow:
			n = &OpenGLWindow{QWindow: *deduced}

		default:
			n = new(OpenGLWindow)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackOpenGLWindow228a58_Constructor
func callbackOpenGLWindow228a58_Constructor(ptr unsafe.Pointer) {
	this := NewOpenGLWindowFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

func OpenGLWindow_QRegisterMetaType() int {
	return int(int32(C.OpenGLWindow228a58_OpenGLWindow228a58_QRegisterMetaType()))
}

func (ptr *OpenGLWindow) QRegisterMetaType() int {
	return int(int32(C.OpenGLWindow228a58_OpenGLWindow228a58_QRegisterMetaType()))
}

func OpenGLWindow_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.OpenGLWindow228a58_OpenGLWindow228a58_QRegisterMetaType2(typeNameC)))
}

func (ptr *OpenGLWindow) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.OpenGLWindow228a58_OpenGLWindow228a58_QRegisterMetaType2(typeNameC)))
}

func OpenGLWindow_QmlRegisterType() int {
	return int(int32(C.OpenGLWindow228a58_OpenGLWindow228a58_QmlRegisterType()))
}

func (ptr *OpenGLWindow) QmlRegisterType() int {
	return int(int32(C.OpenGLWindow228a58_OpenGLWindow228a58_QmlRegisterType()))
}

func OpenGLWindow_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.OpenGLWindow228a58_OpenGLWindow228a58_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *OpenGLWindow) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.OpenGLWindow228a58_OpenGLWindow228a58_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *OpenGLWindow) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.OpenGLWindow228a58___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *OpenGLWindow) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *OpenGLWindow) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.OpenGLWindow228a58___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *OpenGLWindow) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.OpenGLWindow228a58___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *OpenGLWindow) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *OpenGLWindow) __findChildren_newList2() unsafe.Pointer {
	return C.OpenGLWindow228a58___findChildren_newList2(ptr.Pointer())
}

func (ptr *OpenGLWindow) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.OpenGLWindow228a58___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *OpenGLWindow) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *OpenGLWindow) __findChildren_newList3() unsafe.Pointer {
	return C.OpenGLWindow228a58___findChildren_newList3(ptr.Pointer())
}

func (ptr *OpenGLWindow) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.OpenGLWindow228a58___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *OpenGLWindow) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *OpenGLWindow) __findChildren_newList() unsafe.Pointer {
	return C.OpenGLWindow228a58___findChildren_newList(ptr.Pointer())
}

func (ptr *OpenGLWindow) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.OpenGLWindow228a58___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *OpenGLWindow) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *OpenGLWindow) __children_newList() unsafe.Pointer {
	return C.OpenGLWindow228a58___children_newList(ptr.Pointer())
}

func NewOpenGLWindow(targetScreen std_gui.QScreen_ITF) *OpenGLWindow {
	tmpValue := NewOpenGLWindowFromPointer(C.OpenGLWindow228a58_NewOpenGLWindow(std_gui.PointerFromQScreen(targetScreen)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewOpenGLWindow2(parent std_gui.QWindow_ITF) *OpenGLWindow {
	tmpValue := NewOpenGLWindowFromPointer(C.OpenGLWindow228a58_NewOpenGLWindow2(std_gui.PointerFromQWindow(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackOpenGLWindow228a58_DestroyOpenGLWindow
func callbackOpenGLWindow228a58_DestroyOpenGLWindow(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~OpenGLWindow"); signal != nil {
		signal.(func())()
	} else {
		NewOpenGLWindowFromPointer(ptr).DestroyOpenGLWindowDefault()
	}
}

func (ptr *OpenGLWindow) ConnectDestroyOpenGLWindow(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~OpenGLWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~OpenGLWindow", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~OpenGLWindow", f)
		}
	}
}

func (ptr *OpenGLWindow) DisconnectDestroyOpenGLWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~OpenGLWindow")
	}
}

func (ptr *OpenGLWindow) DestroyOpenGLWindow() {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_DestroyOpenGLWindow(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *OpenGLWindow) DestroyOpenGLWindowDefault() {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_DestroyOpenGLWindowDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackOpenGLWindow228a58_Close
func callbackOpenGLWindow228a58_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewOpenGLWindowFromPointer(ptr).CloseDefault())))
}

func (ptr *OpenGLWindow) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.OpenGLWindow228a58_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackOpenGLWindow228a58_Event
func callbackOpenGLWindow228a58_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewOpenGLWindowFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(ev)))))
}

func (ptr *OpenGLWindow) EventDefault(ev std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.OpenGLWindow228a58_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(ev))) != 0
	}
	return false
}

//export callbackOpenGLWindow228a58_NativeEvent
func callbackOpenGLWindow228a58_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QByteArray, unsafe.Pointer, int) bool)(std_core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewOpenGLWindowFromPointer(ptr).NativeEventDefault(std_core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *OpenGLWindow) NativeEventDefault(eventType std_core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return int8(C.OpenGLWindow228a58_NativeEventDefault(ptr.Pointer(), std_core.PointerFromQByteArray(eventType), message, C.long(int32(result)))) != 0
	}
	return false
}

//export callbackOpenGLWindow228a58_ActiveChanged
func callbackOpenGLWindow228a58_ActiveChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activeChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackOpenGLWindow228a58_Alert
func callbackOpenGLWindow228a58_Alert(ptr unsafe.Pointer, msec C.int) {
	if signal := qt.GetSignal(ptr, "alert"); signal != nil {
		signal.(func(int))(int(int32(msec)))
	} else {
		NewOpenGLWindowFromPointer(ptr).AlertDefault(int(int32(msec)))
	}
}

func (ptr *OpenGLWindow) AlertDefault(msec int) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_AlertDefault(ptr.Pointer(), C.int(int32(msec)))
	}
}

//export callbackOpenGLWindow228a58_ContentOrientationChanged
func callbackOpenGLWindow228a58_ContentOrientationChanged(ptr unsafe.Pointer, orientation C.longlong) {
	if signal := qt.GetSignal(ptr, "contentOrientationChanged"); signal != nil {
		signal.(func(std_core.Qt__ScreenOrientation))(std_core.Qt__ScreenOrientation(orientation))
	}

}

//export callbackOpenGLWindow228a58_ExposeEvent
func callbackOpenGLWindow228a58_ExposeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "exposeEvent"); signal != nil {
		signal.(func(*std_gui.QExposeEvent))(std_gui.NewQExposeEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).ExposeEventDefault(std_gui.NewQExposeEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) ExposeEventDefault(ev std_gui.QExposeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_ExposeEventDefault(ptr.Pointer(), std_gui.PointerFromQExposeEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_FocusInEvent
func callbackOpenGLWindow228a58_FocusInEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) FocusInEventDefault(ev std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_FocusObjectChanged
func callbackOpenGLWindow228a58_FocusObjectChanged(ptr unsafe.Pointer, object unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusObjectChanged"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(object))
	}

}

//export callbackOpenGLWindow228a58_FocusOutEvent
func callbackOpenGLWindow228a58_FocusOutEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) FocusOutEventDefault(ev std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_HeightChanged
func callbackOpenGLWindow228a58_HeightChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "heightChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackOpenGLWindow228a58_Hide
func callbackOpenGLWindow228a58_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewOpenGLWindowFromPointer(ptr).HideDefault()
	}
}

func (ptr *OpenGLWindow) HideDefault() {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_HideDefault(ptr.Pointer())
	}
}

//export callbackOpenGLWindow228a58_HideEvent
func callbackOpenGLWindow228a58_HideEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*std_gui.QHideEvent))(std_gui.NewQHideEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) HideEventDefault(ev std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_KeyPressEvent
func callbackOpenGLWindow228a58_KeyPressEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) KeyPressEventDefault(ev std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_KeyReleaseEvent
func callbackOpenGLWindow228a58_KeyReleaseEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) KeyReleaseEventDefault(ev std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_Lower
func callbackOpenGLWindow228a58_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewOpenGLWindowFromPointer(ptr).LowerDefault()
	}
}

func (ptr *OpenGLWindow) LowerDefault() {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_LowerDefault(ptr.Pointer())
	}
}

//export callbackOpenGLWindow228a58_MaximumHeightChanged
func callbackOpenGLWindow228a58_MaximumHeightChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "maximumHeightChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackOpenGLWindow228a58_MaximumWidthChanged
func callbackOpenGLWindow228a58_MaximumWidthChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "maximumWidthChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackOpenGLWindow228a58_MinimumHeightChanged
func callbackOpenGLWindow228a58_MinimumHeightChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "minimumHeightChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackOpenGLWindow228a58_MinimumWidthChanged
func callbackOpenGLWindow228a58_MinimumWidthChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "minimumWidthChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackOpenGLWindow228a58_ModalityChanged
func callbackOpenGLWindow228a58_ModalityChanged(ptr unsafe.Pointer, modality C.longlong) {
	if signal := qt.GetSignal(ptr, "modalityChanged"); signal != nil {
		signal.(func(std_core.Qt__WindowModality))(std_core.Qt__WindowModality(modality))
	}

}

//export callbackOpenGLWindow228a58_MouseDoubleClickEvent
func callbackOpenGLWindow228a58_MouseDoubleClickEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) MouseDoubleClickEventDefault(ev std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_MouseMoveEvent
func callbackOpenGLWindow228a58_MouseMoveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) MouseMoveEventDefault(ev std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_MousePressEvent
func callbackOpenGLWindow228a58_MousePressEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) MousePressEventDefault(ev std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_MouseReleaseEvent
func callbackOpenGLWindow228a58_MouseReleaseEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) MouseReleaseEventDefault(ev std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_MoveEvent
func callbackOpenGLWindow228a58_MoveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*std_gui.QMoveEvent))(std_gui.NewQMoveEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) MoveEventDefault(ev std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_OpacityChanged
func callbackOpenGLWindow228a58_OpacityChanged(ptr unsafe.Pointer, opacity C.double) {
	if signal := qt.GetSignal(ptr, "opacityChanged"); signal != nil {
		signal.(func(float64))(float64(opacity))
	}

}

//export callbackOpenGLWindow228a58_Raise
func callbackOpenGLWindow228a58_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewOpenGLWindowFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *OpenGLWindow) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_RaiseDefault(ptr.Pointer())
	}
}

//export callbackOpenGLWindow228a58_RequestActivate
func callbackOpenGLWindow228a58_RequestActivate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "requestActivate"); signal != nil {
		signal.(func())()
	} else {
		NewOpenGLWindowFromPointer(ptr).RequestActivateDefault()
	}
}

func (ptr *OpenGLWindow) RequestActivateDefault() {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_RequestActivateDefault(ptr.Pointer())
	}
}

//export callbackOpenGLWindow228a58_RequestUpdate
func callbackOpenGLWindow228a58_RequestUpdate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "requestUpdate"); signal != nil {
		signal.(func())()
	} else {
		NewOpenGLWindowFromPointer(ptr).RequestUpdateDefault()
	}
}

func (ptr *OpenGLWindow) RequestUpdateDefault() {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_RequestUpdateDefault(ptr.Pointer())
	}
}

//export callbackOpenGLWindow228a58_ResizeEvent
func callbackOpenGLWindow228a58_ResizeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*std_gui.QResizeEvent))(std_gui.NewQResizeEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) ResizeEventDefault(ev std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_ScreenChanged
func callbackOpenGLWindow228a58_ScreenChanged(ptr unsafe.Pointer, screen unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "screenChanged"); signal != nil {
		signal.(func(*std_gui.QScreen))(std_gui.NewQScreenFromPointer(screen))
	}

}

//export callbackOpenGLWindow228a58_SetGeometry2
func callbackOpenGLWindow228a58_SetGeometry2(ptr unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry2"); signal != nil {
		signal.(func(*std_core.QRect))(std_core.NewQRectFromPointer(rect))
	} else {
		NewOpenGLWindowFromPointer(ptr).SetGeometry2Default(std_core.NewQRectFromPointer(rect))
	}
}

func (ptr *OpenGLWindow) SetGeometry2Default(rect std_core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_SetGeometry2Default(ptr.Pointer(), std_core.PointerFromQRect(rect))
	}
}

//export callbackOpenGLWindow228a58_SetGeometry
func callbackOpenGLWindow228a58_SetGeometry(ptr unsafe.Pointer, posx C.int, posy C.int, w C.int, h C.int) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		signal.(func(int, int, int, int))(int(int32(posx)), int(int32(posy)), int(int32(w)), int(int32(h)))
	} else {
		NewOpenGLWindowFromPointer(ptr).SetGeometryDefault(int(int32(posx)), int(int32(posy)), int(int32(w)), int(int32(h)))
	}
}

func (ptr *OpenGLWindow) SetGeometryDefault(posx int, posy int, w int, h int) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_SetGeometryDefault(ptr.Pointer(), C.int(int32(posx)), C.int(int32(posy)), C.int(int32(w)), C.int(int32(h)))
	}
}

//export callbackOpenGLWindow228a58_SetHeight
func callbackOpenGLWindow228a58_SetHeight(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "setHeight"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewOpenGLWindowFromPointer(ptr).SetHeightDefault(int(int32(arg)))
	}
}

func (ptr *OpenGLWindow) SetHeightDefault(arg int) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_SetHeightDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackOpenGLWindow228a58_SetMaximumHeight
func callbackOpenGLWindow228a58_SetMaximumHeight(ptr unsafe.Pointer, h C.int) {
	if signal := qt.GetSignal(ptr, "setMaximumHeight"); signal != nil {
		signal.(func(int))(int(int32(h)))
	} else {
		NewOpenGLWindowFromPointer(ptr).SetMaximumHeightDefault(int(int32(h)))
	}
}

func (ptr *OpenGLWindow) SetMaximumHeightDefault(h int) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_SetMaximumHeightDefault(ptr.Pointer(), C.int(int32(h)))
	}
}

//export callbackOpenGLWindow228a58_SetMaximumWidth
func callbackOpenGLWindow228a58_SetMaximumWidth(ptr unsafe.Pointer, w C.int) {
	if signal := qt.GetSignal(ptr, "setMaximumWidth"); signal != nil {
		signal.(func(int))(int(int32(w)))
	} else {
		NewOpenGLWindowFromPointer(ptr).SetMaximumWidthDefault(int(int32(w)))
	}
}

func (ptr *OpenGLWindow) SetMaximumWidthDefault(w int) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_SetMaximumWidthDefault(ptr.Pointer(), C.int(int32(w)))
	}
}

//export callbackOpenGLWindow228a58_SetMinimumHeight
func callbackOpenGLWindow228a58_SetMinimumHeight(ptr unsafe.Pointer, h C.int) {
	if signal := qt.GetSignal(ptr, "setMinimumHeight"); signal != nil {
		signal.(func(int))(int(int32(h)))
	} else {
		NewOpenGLWindowFromPointer(ptr).SetMinimumHeightDefault(int(int32(h)))
	}
}

func (ptr *OpenGLWindow) SetMinimumHeightDefault(h int) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_SetMinimumHeightDefault(ptr.Pointer(), C.int(int32(h)))
	}
}

//export callbackOpenGLWindow228a58_SetMinimumWidth
func callbackOpenGLWindow228a58_SetMinimumWidth(ptr unsafe.Pointer, w C.int) {
	if signal := qt.GetSignal(ptr, "setMinimumWidth"); signal != nil {
		signal.(func(int))(int(int32(w)))
	} else {
		NewOpenGLWindowFromPointer(ptr).SetMinimumWidthDefault(int(int32(w)))
	}
}

func (ptr *OpenGLWindow) SetMinimumWidthDefault(w int) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_SetMinimumWidthDefault(ptr.Pointer(), C.int(int32(w)))
	}
}

//export callbackOpenGLWindow228a58_SetTitle
func callbackOpenGLWindow228a58_SetTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewOpenGLWindowFromPointer(ptr).SetTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *OpenGLWindow) SetTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.OpenGLWindow228a58_SetTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackOpenGLWindow228a58_SetVisible
func callbackOpenGLWindow228a58_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewOpenGLWindowFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *OpenGLWindow) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackOpenGLWindow228a58_SetWidth
func callbackOpenGLWindow228a58_SetWidth(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "setWidth"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewOpenGLWindowFromPointer(ptr).SetWidthDefault(int(int32(arg)))
	}
}

func (ptr *OpenGLWindow) SetWidthDefault(arg int) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_SetWidthDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackOpenGLWindow228a58_SetX
func callbackOpenGLWindow228a58_SetX(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "setX"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewOpenGLWindowFromPointer(ptr).SetXDefault(int(int32(arg)))
	}
}

func (ptr *OpenGLWindow) SetXDefault(arg int) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_SetXDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackOpenGLWindow228a58_SetY
func callbackOpenGLWindow228a58_SetY(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "setY"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewOpenGLWindowFromPointer(ptr).SetYDefault(int(int32(arg)))
	}
}

func (ptr *OpenGLWindow) SetYDefault(arg int) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_SetYDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackOpenGLWindow228a58_Show
func callbackOpenGLWindow228a58_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewOpenGLWindowFromPointer(ptr).ShowDefault()
	}
}

func (ptr *OpenGLWindow) ShowDefault() {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_ShowDefault(ptr.Pointer())
	}
}

//export callbackOpenGLWindow228a58_ShowEvent
func callbackOpenGLWindow228a58_ShowEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*std_gui.QShowEvent))(std_gui.NewQShowEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) ShowEventDefault(ev std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_ShowFullScreen
func callbackOpenGLWindow228a58_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewOpenGLWindowFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *OpenGLWindow) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackOpenGLWindow228a58_ShowMaximized
func callbackOpenGLWindow228a58_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewOpenGLWindowFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *OpenGLWindow) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackOpenGLWindow228a58_ShowMinimized
func callbackOpenGLWindow228a58_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewOpenGLWindowFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *OpenGLWindow) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackOpenGLWindow228a58_ShowNormal
func callbackOpenGLWindow228a58_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewOpenGLWindowFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *OpenGLWindow) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackOpenGLWindow228a58_TabletEvent
func callbackOpenGLWindow228a58_TabletEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*std_gui.QTabletEvent))(std_gui.NewQTabletEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) TabletEventDefault(ev std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_TouchEvent
func callbackOpenGLWindow228a58_TouchEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "touchEvent"); signal != nil {
		signal.(func(*std_gui.QTouchEvent))(std_gui.NewQTouchEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).TouchEventDefault(std_gui.NewQTouchEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) TouchEventDefault(ev std_gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_TouchEventDefault(ptr.Pointer(), std_gui.PointerFromQTouchEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_VisibilityChanged
func callbackOpenGLWindow228a58_VisibilityChanged(ptr unsafe.Pointer, visibility C.longlong) {
	if signal := qt.GetSignal(ptr, "visibilityChanged"); signal != nil {
		signal.(func(std_gui.QWindow__Visibility))(std_gui.QWindow__Visibility(visibility))
	}

}

//export callbackOpenGLWindow228a58_VisibleChanged
func callbackOpenGLWindow228a58_VisibleChanged(ptr unsafe.Pointer, arg C.char) {
	if signal := qt.GetSignal(ptr, "visibleChanged"); signal != nil {
		signal.(func(bool))(int8(arg) != 0)
	}

}

//export callbackOpenGLWindow228a58_WheelEvent
func callbackOpenGLWindow228a58_WheelEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*std_gui.QWheelEvent))(std_gui.NewQWheelEventFromPointer(ev))
	} else {
		NewOpenGLWindowFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(ev))
	}
}

func (ptr *OpenGLWindow) WheelEventDefault(ev std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(ev))
	}
}

//export callbackOpenGLWindow228a58_WidthChanged
func callbackOpenGLWindow228a58_WidthChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "widthChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackOpenGLWindow228a58_WindowStateChanged
func callbackOpenGLWindow228a58_WindowStateChanged(ptr unsafe.Pointer, windowState C.longlong) {
	if signal := qt.GetSignal(ptr, "windowStateChanged"); signal != nil {
		signal.(func(std_core.Qt__WindowState))(std_core.Qt__WindowState(windowState))
	}

}

//export callbackOpenGLWindow228a58_WindowTitleChanged
func callbackOpenGLWindow228a58_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackOpenGLWindow228a58_XChanged
func callbackOpenGLWindow228a58_XChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "xChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackOpenGLWindow228a58_YChanged
func callbackOpenGLWindow228a58_YChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "yChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackOpenGLWindow228a58_FocusObject
func callbackOpenGLWindow228a58_FocusObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "focusObject"); signal != nil {
		return std_core.PointerFromQObject(signal.(func() *std_core.QObject)())
	}

	return std_core.PointerFromQObject(NewOpenGLWindowFromPointer(ptr).FocusObjectDefault())
}

func (ptr *OpenGLWindow) FocusObjectDefault() *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.OpenGLWindow228a58_FocusObjectDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackOpenGLWindow228a58_Size
func callbackOpenGLWindow228a58_Size(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "size"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewOpenGLWindowFromPointer(ptr).SizeDefault())
}

func (ptr *OpenGLWindow) SizeDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.OpenGLWindow228a58_SizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackOpenGLWindow228a58_SurfaceType
func callbackOpenGLWindow228a58_SurfaceType(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "surfaceType"); signal != nil {
		return C.longlong(signal.(func() std_gui.QSurface__SurfaceType)())
	}

	return C.longlong(NewOpenGLWindowFromPointer(ptr).SurfaceTypeDefault())
}

func (ptr *OpenGLWindow) SurfaceTypeDefault() std_gui.QSurface__SurfaceType {
	if ptr.Pointer() != nil {
		return std_gui.QSurface__SurfaceType(C.OpenGLWindow228a58_SurfaceTypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackOpenGLWindow228a58_Format
func callbackOpenGLWindow228a58_Format(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "format"); signal != nil {
		return std_gui.PointerFromQSurfaceFormat(signal.(func() *std_gui.QSurfaceFormat)())
	}

	return std_gui.PointerFromQSurfaceFormat(NewOpenGLWindowFromPointer(ptr).FormatDefault())
}

func (ptr *OpenGLWindow) FormatDefault() *std_gui.QSurfaceFormat {
	if ptr.Pointer() != nil {
		tmpValue := std_gui.NewQSurfaceFormatFromPointer(C.OpenGLWindow228a58_FormatDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_gui.QSurfaceFormat).DestroyQSurfaceFormat)
		return tmpValue
	}
	return nil
}

//export callbackOpenGLWindow228a58_EventFilter
func callbackOpenGLWindow228a58_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewOpenGLWindowFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *OpenGLWindow) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.OpenGLWindow228a58_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackOpenGLWindow228a58_ChildEvent
func callbackOpenGLWindow228a58_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewOpenGLWindowFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *OpenGLWindow) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackOpenGLWindow228a58_ConnectNotify
func callbackOpenGLWindow228a58_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewOpenGLWindowFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *OpenGLWindow) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackOpenGLWindow228a58_CustomEvent
func callbackOpenGLWindow228a58_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewOpenGLWindowFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *OpenGLWindow) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackOpenGLWindow228a58_DeleteLater
func callbackOpenGLWindow228a58_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewOpenGLWindowFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *OpenGLWindow) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackOpenGLWindow228a58_Destroyed
func callbackOpenGLWindow228a58_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackOpenGLWindow228a58_DisconnectNotify
func callbackOpenGLWindow228a58_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewOpenGLWindowFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *OpenGLWindow) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackOpenGLWindow228a58_ObjectNameChanged
func callbackOpenGLWindow228a58_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackOpenGLWindow228a58_TimerEvent
func callbackOpenGLWindow228a58_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewOpenGLWindowFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *OpenGLWindow) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.OpenGLWindow228a58_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type TriangleWindow_ITF interface {
	OpenGLWindow_ITF
	TriangleWindow_PTR() *TriangleWindow
}

func (ptr *TriangleWindow) TriangleWindow_PTR() *TriangleWindow {
	return ptr
}

func (ptr *TriangleWindow) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.OpenGLWindow_PTR().Pointer()
	}
	return nil
}

func (ptr *TriangleWindow) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.OpenGLWindow_PTR().SetPointer(p)
	}
}

func PointerFromTriangleWindow(ptr TriangleWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.TriangleWindow_PTR().Pointer()
	}
	return nil
}

func NewTriangleWindowFromPointer(ptr unsafe.Pointer) (n *TriangleWindow) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(TriangleWindow)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *TriangleWindow:
			n = deduced

		case *OpenGLWindow:
			n = &TriangleWindow{OpenGLWindow: *deduced}

		default:
			n = new(TriangleWindow)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackTriangleWindow228a58_Constructor
func callbackTriangleWindow228a58_Constructor(ptr unsafe.Pointer) {
	this := NewTriangleWindowFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

func TriangleWindow_QRegisterMetaType() int {
	return int(int32(C.TriangleWindow228a58_TriangleWindow228a58_QRegisterMetaType()))
}

func (ptr *TriangleWindow) QRegisterMetaType() int {
	return int(int32(C.TriangleWindow228a58_TriangleWindow228a58_QRegisterMetaType()))
}

func TriangleWindow_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TriangleWindow228a58_TriangleWindow228a58_QRegisterMetaType2(typeNameC)))
}

func (ptr *TriangleWindow) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TriangleWindow228a58_TriangleWindow228a58_QRegisterMetaType2(typeNameC)))
}

func TriangleWindow_QmlRegisterType() int {
	return int(int32(C.TriangleWindow228a58_TriangleWindow228a58_QmlRegisterType()))
}

func (ptr *TriangleWindow) QmlRegisterType() int {
	return int(int32(C.TriangleWindow228a58_TriangleWindow228a58_QmlRegisterType()))
}

func TriangleWindow_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.TriangleWindow228a58_TriangleWindow228a58_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TriangleWindow) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.TriangleWindow228a58_TriangleWindow228a58_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func NewTriangleWindow(targetScreen std_gui.QScreen_ITF) *TriangleWindow {
	tmpValue := NewTriangleWindowFromPointer(C.TriangleWindow228a58_NewTriangleWindow(std_gui.PointerFromQScreen(targetScreen)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewTriangleWindow2(parent std_gui.QWindow_ITF) *TriangleWindow {
	tmpValue := NewTriangleWindowFromPointer(C.TriangleWindow228a58_NewTriangleWindow2(std_gui.PointerFromQWindow(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackTriangleWindow228a58_DestroyTriangleWindow
func callbackTriangleWindow228a58_DestroyTriangleWindow(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~TriangleWindow"); signal != nil {
		signal.(func())()
	} else {
		NewTriangleWindowFromPointer(ptr).DestroyTriangleWindowDefault()
	}
}

func (ptr *TriangleWindow) ConnectDestroyTriangleWindow(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~TriangleWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~TriangleWindow", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~TriangleWindow", f)
		}
	}
}

func (ptr *TriangleWindow) DisconnectDestroyTriangleWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~TriangleWindow")
	}
}

func (ptr *TriangleWindow) DestroyTriangleWindow() {
	if ptr.Pointer() != nil {
		C.TriangleWindow228a58_DestroyTriangleWindow(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *TriangleWindow) DestroyTriangleWindowDefault() {
	if ptr.Pointer() != nil {
		C.TriangleWindow228a58_DestroyTriangleWindowDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}
