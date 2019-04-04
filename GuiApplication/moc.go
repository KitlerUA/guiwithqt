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

type RasterWindow_ITF interface {
	std_gui.QWindow_ITF
	RasterWindow_PTR() *RasterWindow
}

func (ptr *RasterWindow) RasterWindow_PTR() *RasterWindow {
	return ptr
}

func (ptr *RasterWindow) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWindow_PTR().Pointer()
	}
	return nil
}

func (ptr *RasterWindow) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWindow_PTR().SetPointer(p)
	}
}

func PointerFromRasterWindow(ptr RasterWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.RasterWindow_PTR().Pointer()
	}
	return nil
}

func NewRasterWindowFromPointer(ptr unsafe.Pointer) (n *RasterWindow) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(RasterWindow)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *RasterWindow:
			n = deduced

		case *std_gui.QWindow:
			n = &RasterWindow{QWindow: *deduced}

		default:
			n = new(RasterWindow)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackRasterWindow000290_Constructor
func callbackRasterWindow000290_Constructor(ptr unsafe.Pointer) {
	this := NewRasterWindowFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

func RasterWindow_QRegisterMetaType() int {
	return int(int32(C.RasterWindow000290_RasterWindow000290_QRegisterMetaType()))
}

func (ptr *RasterWindow) QRegisterMetaType() int {
	return int(int32(C.RasterWindow000290_RasterWindow000290_QRegisterMetaType()))
}

func RasterWindow_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.RasterWindow000290_RasterWindow000290_QRegisterMetaType2(typeNameC)))
}

func (ptr *RasterWindow) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.RasterWindow000290_RasterWindow000290_QRegisterMetaType2(typeNameC)))
}

func RasterWindow_QmlRegisterType() int {
	return int(int32(C.RasterWindow000290_RasterWindow000290_QmlRegisterType()))
}

func (ptr *RasterWindow) QmlRegisterType() int {
	return int(int32(C.RasterWindow000290_RasterWindow000290_QmlRegisterType()))
}

func RasterWindow_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.RasterWindow000290_RasterWindow000290_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *RasterWindow) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.RasterWindow000290_RasterWindow000290_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *RasterWindow) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.RasterWindow000290___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *RasterWindow) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *RasterWindow) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.RasterWindow000290___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *RasterWindow) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.RasterWindow000290___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *RasterWindow) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *RasterWindow) __findChildren_newList2() unsafe.Pointer {
	return C.RasterWindow000290___findChildren_newList2(ptr.Pointer())
}

func (ptr *RasterWindow) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.RasterWindow000290___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *RasterWindow) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *RasterWindow) __findChildren_newList3() unsafe.Pointer {
	return C.RasterWindow000290___findChildren_newList3(ptr.Pointer())
}

func (ptr *RasterWindow) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.RasterWindow000290___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *RasterWindow) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *RasterWindow) __findChildren_newList() unsafe.Pointer {
	return C.RasterWindow000290___findChildren_newList(ptr.Pointer())
}

func (ptr *RasterWindow) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.RasterWindow000290___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *RasterWindow) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *RasterWindow) __children_newList() unsafe.Pointer {
	return C.RasterWindow000290___children_newList(ptr.Pointer())
}

func NewRasterWindow(targetScreen std_gui.QScreen_ITF) *RasterWindow {
	tmpValue := NewRasterWindowFromPointer(C.RasterWindow000290_NewRasterWindow(std_gui.PointerFromQScreen(targetScreen)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewRasterWindow2(parent std_gui.QWindow_ITF) *RasterWindow {
	tmpValue := NewRasterWindowFromPointer(C.RasterWindow000290_NewRasterWindow2(std_gui.PointerFromQWindow(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackRasterWindow000290_DestroyRasterWindow
func callbackRasterWindow000290_DestroyRasterWindow(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~RasterWindow"); signal != nil {
		signal.(func())()
	} else {
		NewRasterWindowFromPointer(ptr).DestroyRasterWindowDefault()
	}
}

func (ptr *RasterWindow) ConnectDestroyRasterWindow(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~RasterWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~RasterWindow", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~RasterWindow", f)
		}
	}
}

func (ptr *RasterWindow) DisconnectDestroyRasterWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~RasterWindow")
	}
}

func (ptr *RasterWindow) DestroyRasterWindow() {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_DestroyRasterWindow(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *RasterWindow) DestroyRasterWindowDefault() {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_DestroyRasterWindowDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackRasterWindow000290_Close
func callbackRasterWindow000290_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewRasterWindowFromPointer(ptr).CloseDefault())))
}

func (ptr *RasterWindow) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.RasterWindow000290_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackRasterWindow000290_Event
func callbackRasterWindow000290_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewRasterWindowFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(ev)))))
}

func (ptr *RasterWindow) EventDefault(ev std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.RasterWindow000290_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(ev))) != 0
	}
	return false
}

//export callbackRasterWindow000290_NativeEvent
func callbackRasterWindow000290_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QByteArray, unsafe.Pointer, int) bool)(std_core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewRasterWindowFromPointer(ptr).NativeEventDefault(std_core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *RasterWindow) NativeEventDefault(eventType std_core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return int8(C.RasterWindow000290_NativeEventDefault(ptr.Pointer(), std_core.PointerFromQByteArray(eventType), message, C.long(int32(result)))) != 0
	}
	return false
}

//export callbackRasterWindow000290_ActiveChanged
func callbackRasterWindow000290_ActiveChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activeChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackRasterWindow000290_Alert
func callbackRasterWindow000290_Alert(ptr unsafe.Pointer, msec C.int) {
	if signal := qt.GetSignal(ptr, "alert"); signal != nil {
		signal.(func(int))(int(int32(msec)))
	} else {
		NewRasterWindowFromPointer(ptr).AlertDefault(int(int32(msec)))
	}
}

func (ptr *RasterWindow) AlertDefault(msec int) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_AlertDefault(ptr.Pointer(), C.int(int32(msec)))
	}
}

//export callbackRasterWindow000290_ContentOrientationChanged
func callbackRasterWindow000290_ContentOrientationChanged(ptr unsafe.Pointer, orientation C.longlong) {
	if signal := qt.GetSignal(ptr, "contentOrientationChanged"); signal != nil {
		signal.(func(std_core.Qt__ScreenOrientation))(std_core.Qt__ScreenOrientation(orientation))
	}

}

//export callbackRasterWindow000290_ExposeEvent
func callbackRasterWindow000290_ExposeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "exposeEvent"); signal != nil {
		signal.(func(*std_gui.QExposeEvent))(std_gui.NewQExposeEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).ExposeEventDefault(std_gui.NewQExposeEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) ExposeEventDefault(ev std_gui.QExposeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_ExposeEventDefault(ptr.Pointer(), std_gui.PointerFromQExposeEvent(ev))
	}
}

//export callbackRasterWindow000290_FocusInEvent
func callbackRasterWindow000290_FocusInEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) FocusInEventDefault(ev std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackRasterWindow000290_FocusObjectChanged
func callbackRasterWindow000290_FocusObjectChanged(ptr unsafe.Pointer, object unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusObjectChanged"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(object))
	}

}

//export callbackRasterWindow000290_FocusOutEvent
func callbackRasterWindow000290_FocusOutEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) FocusOutEventDefault(ev std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackRasterWindow000290_HeightChanged
func callbackRasterWindow000290_HeightChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "heightChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackRasterWindow000290_Hide
func callbackRasterWindow000290_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewRasterWindowFromPointer(ptr).HideDefault()
	}
}

func (ptr *RasterWindow) HideDefault() {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_HideDefault(ptr.Pointer())
	}
}

//export callbackRasterWindow000290_HideEvent
func callbackRasterWindow000290_HideEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*std_gui.QHideEvent))(std_gui.NewQHideEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) HideEventDefault(ev std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(ev))
	}
}

//export callbackRasterWindow000290_KeyPressEvent
func callbackRasterWindow000290_KeyPressEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) KeyPressEventDefault(ev std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(ev))
	}
}

//export callbackRasterWindow000290_KeyReleaseEvent
func callbackRasterWindow000290_KeyReleaseEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) KeyReleaseEventDefault(ev std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(ev))
	}
}

//export callbackRasterWindow000290_Lower
func callbackRasterWindow000290_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewRasterWindowFromPointer(ptr).LowerDefault()
	}
}

func (ptr *RasterWindow) LowerDefault() {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_LowerDefault(ptr.Pointer())
	}
}

//export callbackRasterWindow000290_MaximumHeightChanged
func callbackRasterWindow000290_MaximumHeightChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "maximumHeightChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackRasterWindow000290_MaximumWidthChanged
func callbackRasterWindow000290_MaximumWidthChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "maximumWidthChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackRasterWindow000290_MinimumHeightChanged
func callbackRasterWindow000290_MinimumHeightChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "minimumHeightChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackRasterWindow000290_MinimumWidthChanged
func callbackRasterWindow000290_MinimumWidthChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "minimumWidthChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackRasterWindow000290_ModalityChanged
func callbackRasterWindow000290_ModalityChanged(ptr unsafe.Pointer, modality C.longlong) {
	if signal := qt.GetSignal(ptr, "modalityChanged"); signal != nil {
		signal.(func(std_core.Qt__WindowModality))(std_core.Qt__WindowModality(modality))
	}

}

//export callbackRasterWindow000290_MouseDoubleClickEvent
func callbackRasterWindow000290_MouseDoubleClickEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) MouseDoubleClickEventDefault(ev std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackRasterWindow000290_MouseMoveEvent
func callbackRasterWindow000290_MouseMoveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) MouseMoveEventDefault(ev std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackRasterWindow000290_MousePressEvent
func callbackRasterWindow000290_MousePressEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) MousePressEventDefault(ev std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackRasterWindow000290_MouseReleaseEvent
func callbackRasterWindow000290_MouseReleaseEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) MouseReleaseEventDefault(ev std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackRasterWindow000290_MoveEvent
func callbackRasterWindow000290_MoveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*std_gui.QMoveEvent))(std_gui.NewQMoveEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) MoveEventDefault(ev std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(ev))
	}
}

//export callbackRasterWindow000290_OpacityChanged
func callbackRasterWindow000290_OpacityChanged(ptr unsafe.Pointer, opacity C.double) {
	if signal := qt.GetSignal(ptr, "opacityChanged"); signal != nil {
		signal.(func(float64))(float64(opacity))
	}

}

//export callbackRasterWindow000290_Raise
func callbackRasterWindow000290_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewRasterWindowFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *RasterWindow) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_RaiseDefault(ptr.Pointer())
	}
}

//export callbackRasterWindow000290_RequestActivate
func callbackRasterWindow000290_RequestActivate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "requestActivate"); signal != nil {
		signal.(func())()
	} else {
		NewRasterWindowFromPointer(ptr).RequestActivateDefault()
	}
}

func (ptr *RasterWindow) RequestActivateDefault() {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_RequestActivateDefault(ptr.Pointer())
	}
}

//export callbackRasterWindow000290_RequestUpdate
func callbackRasterWindow000290_RequestUpdate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "requestUpdate"); signal != nil {
		signal.(func())()
	} else {
		NewRasterWindowFromPointer(ptr).RequestUpdateDefault()
	}
}

func (ptr *RasterWindow) RequestUpdateDefault() {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_RequestUpdateDefault(ptr.Pointer())
	}
}

//export callbackRasterWindow000290_ResizeEvent
func callbackRasterWindow000290_ResizeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*std_gui.QResizeEvent))(std_gui.NewQResizeEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) ResizeEventDefault(ev std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(ev))
	}
}

//export callbackRasterWindow000290_ScreenChanged
func callbackRasterWindow000290_ScreenChanged(ptr unsafe.Pointer, screen unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "screenChanged"); signal != nil {
		signal.(func(*std_gui.QScreen))(std_gui.NewQScreenFromPointer(screen))
	}

}

//export callbackRasterWindow000290_SetGeometry2
func callbackRasterWindow000290_SetGeometry2(ptr unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry2"); signal != nil {
		signal.(func(*std_core.QRect))(std_core.NewQRectFromPointer(rect))
	} else {
		NewRasterWindowFromPointer(ptr).SetGeometry2Default(std_core.NewQRectFromPointer(rect))
	}
}

func (ptr *RasterWindow) SetGeometry2Default(rect std_core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_SetGeometry2Default(ptr.Pointer(), std_core.PointerFromQRect(rect))
	}
}

//export callbackRasterWindow000290_SetGeometry
func callbackRasterWindow000290_SetGeometry(ptr unsafe.Pointer, posx C.int, posy C.int, w C.int, h C.int) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		signal.(func(int, int, int, int))(int(int32(posx)), int(int32(posy)), int(int32(w)), int(int32(h)))
	} else {
		NewRasterWindowFromPointer(ptr).SetGeometryDefault(int(int32(posx)), int(int32(posy)), int(int32(w)), int(int32(h)))
	}
}

func (ptr *RasterWindow) SetGeometryDefault(posx int, posy int, w int, h int) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_SetGeometryDefault(ptr.Pointer(), C.int(int32(posx)), C.int(int32(posy)), C.int(int32(w)), C.int(int32(h)))
	}
}

//export callbackRasterWindow000290_SetHeight
func callbackRasterWindow000290_SetHeight(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "setHeight"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewRasterWindowFromPointer(ptr).SetHeightDefault(int(int32(arg)))
	}
}

func (ptr *RasterWindow) SetHeightDefault(arg int) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_SetHeightDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackRasterWindow000290_SetMaximumHeight
func callbackRasterWindow000290_SetMaximumHeight(ptr unsafe.Pointer, h C.int) {
	if signal := qt.GetSignal(ptr, "setMaximumHeight"); signal != nil {
		signal.(func(int))(int(int32(h)))
	} else {
		NewRasterWindowFromPointer(ptr).SetMaximumHeightDefault(int(int32(h)))
	}
}

func (ptr *RasterWindow) SetMaximumHeightDefault(h int) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_SetMaximumHeightDefault(ptr.Pointer(), C.int(int32(h)))
	}
}

//export callbackRasterWindow000290_SetMaximumWidth
func callbackRasterWindow000290_SetMaximumWidth(ptr unsafe.Pointer, w C.int) {
	if signal := qt.GetSignal(ptr, "setMaximumWidth"); signal != nil {
		signal.(func(int))(int(int32(w)))
	} else {
		NewRasterWindowFromPointer(ptr).SetMaximumWidthDefault(int(int32(w)))
	}
}

func (ptr *RasterWindow) SetMaximumWidthDefault(w int) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_SetMaximumWidthDefault(ptr.Pointer(), C.int(int32(w)))
	}
}

//export callbackRasterWindow000290_SetMinimumHeight
func callbackRasterWindow000290_SetMinimumHeight(ptr unsafe.Pointer, h C.int) {
	if signal := qt.GetSignal(ptr, "setMinimumHeight"); signal != nil {
		signal.(func(int))(int(int32(h)))
	} else {
		NewRasterWindowFromPointer(ptr).SetMinimumHeightDefault(int(int32(h)))
	}
}

func (ptr *RasterWindow) SetMinimumHeightDefault(h int) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_SetMinimumHeightDefault(ptr.Pointer(), C.int(int32(h)))
	}
}

//export callbackRasterWindow000290_SetMinimumWidth
func callbackRasterWindow000290_SetMinimumWidth(ptr unsafe.Pointer, w C.int) {
	if signal := qt.GetSignal(ptr, "setMinimumWidth"); signal != nil {
		signal.(func(int))(int(int32(w)))
	} else {
		NewRasterWindowFromPointer(ptr).SetMinimumWidthDefault(int(int32(w)))
	}
}

func (ptr *RasterWindow) SetMinimumWidthDefault(w int) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_SetMinimumWidthDefault(ptr.Pointer(), C.int(int32(w)))
	}
}

//export callbackRasterWindow000290_SetTitle
func callbackRasterWindow000290_SetTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewRasterWindowFromPointer(ptr).SetTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *RasterWindow) SetTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.RasterWindow000290_SetTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackRasterWindow000290_SetVisible
func callbackRasterWindow000290_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewRasterWindowFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *RasterWindow) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackRasterWindow000290_SetWidth
func callbackRasterWindow000290_SetWidth(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "setWidth"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewRasterWindowFromPointer(ptr).SetWidthDefault(int(int32(arg)))
	}
}

func (ptr *RasterWindow) SetWidthDefault(arg int) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_SetWidthDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackRasterWindow000290_SetX
func callbackRasterWindow000290_SetX(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "setX"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewRasterWindowFromPointer(ptr).SetXDefault(int(int32(arg)))
	}
}

func (ptr *RasterWindow) SetXDefault(arg int) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_SetXDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackRasterWindow000290_SetY
func callbackRasterWindow000290_SetY(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "setY"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewRasterWindowFromPointer(ptr).SetYDefault(int(int32(arg)))
	}
}

func (ptr *RasterWindow) SetYDefault(arg int) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_SetYDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackRasterWindow000290_Show
func callbackRasterWindow000290_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewRasterWindowFromPointer(ptr).ShowDefault()
	}
}

func (ptr *RasterWindow) ShowDefault() {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_ShowDefault(ptr.Pointer())
	}
}

//export callbackRasterWindow000290_ShowEvent
func callbackRasterWindow000290_ShowEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*std_gui.QShowEvent))(std_gui.NewQShowEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) ShowEventDefault(ev std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(ev))
	}
}

//export callbackRasterWindow000290_ShowFullScreen
func callbackRasterWindow000290_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewRasterWindowFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *RasterWindow) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackRasterWindow000290_ShowMaximized
func callbackRasterWindow000290_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewRasterWindowFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *RasterWindow) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackRasterWindow000290_ShowMinimized
func callbackRasterWindow000290_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewRasterWindowFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *RasterWindow) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackRasterWindow000290_ShowNormal
func callbackRasterWindow000290_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewRasterWindowFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *RasterWindow) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackRasterWindow000290_TabletEvent
func callbackRasterWindow000290_TabletEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*std_gui.QTabletEvent))(std_gui.NewQTabletEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) TabletEventDefault(ev std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(ev))
	}
}

//export callbackRasterWindow000290_TouchEvent
func callbackRasterWindow000290_TouchEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "touchEvent"); signal != nil {
		signal.(func(*std_gui.QTouchEvent))(std_gui.NewQTouchEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).TouchEventDefault(std_gui.NewQTouchEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) TouchEventDefault(ev std_gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_TouchEventDefault(ptr.Pointer(), std_gui.PointerFromQTouchEvent(ev))
	}
}

//export callbackRasterWindow000290_VisibilityChanged
func callbackRasterWindow000290_VisibilityChanged(ptr unsafe.Pointer, visibility C.longlong) {
	if signal := qt.GetSignal(ptr, "visibilityChanged"); signal != nil {
		signal.(func(std_gui.QWindow__Visibility))(std_gui.QWindow__Visibility(visibility))
	}

}

//export callbackRasterWindow000290_VisibleChanged
func callbackRasterWindow000290_VisibleChanged(ptr unsafe.Pointer, arg C.char) {
	if signal := qt.GetSignal(ptr, "visibleChanged"); signal != nil {
		signal.(func(bool))(int8(arg) != 0)
	}

}

//export callbackRasterWindow000290_WheelEvent
func callbackRasterWindow000290_WheelEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*std_gui.QWheelEvent))(std_gui.NewQWheelEventFromPointer(ev))
	} else {
		NewRasterWindowFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(ev))
	}
}

func (ptr *RasterWindow) WheelEventDefault(ev std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(ev))
	}
}

//export callbackRasterWindow000290_WidthChanged
func callbackRasterWindow000290_WidthChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "widthChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackRasterWindow000290_WindowStateChanged
func callbackRasterWindow000290_WindowStateChanged(ptr unsafe.Pointer, windowState C.longlong) {
	if signal := qt.GetSignal(ptr, "windowStateChanged"); signal != nil {
		signal.(func(std_core.Qt__WindowState))(std_core.Qt__WindowState(windowState))
	}

}

//export callbackRasterWindow000290_WindowTitleChanged
func callbackRasterWindow000290_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackRasterWindow000290_XChanged
func callbackRasterWindow000290_XChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "xChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackRasterWindow000290_YChanged
func callbackRasterWindow000290_YChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "yChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackRasterWindow000290_FocusObject
func callbackRasterWindow000290_FocusObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "focusObject"); signal != nil {
		return std_core.PointerFromQObject(signal.(func() *std_core.QObject)())
	}

	return std_core.PointerFromQObject(NewRasterWindowFromPointer(ptr).FocusObjectDefault())
}

func (ptr *RasterWindow) FocusObjectDefault() *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.RasterWindow000290_FocusObjectDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackRasterWindow000290_Size
func callbackRasterWindow000290_Size(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "size"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewRasterWindowFromPointer(ptr).SizeDefault())
}

func (ptr *RasterWindow) SizeDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.RasterWindow000290_SizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackRasterWindow000290_SurfaceType
func callbackRasterWindow000290_SurfaceType(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "surfaceType"); signal != nil {
		return C.longlong(signal.(func() std_gui.QSurface__SurfaceType)())
	}

	return C.longlong(NewRasterWindowFromPointer(ptr).SurfaceTypeDefault())
}

func (ptr *RasterWindow) SurfaceTypeDefault() std_gui.QSurface__SurfaceType {
	if ptr.Pointer() != nil {
		return std_gui.QSurface__SurfaceType(C.RasterWindow000290_SurfaceTypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackRasterWindow000290_Format
func callbackRasterWindow000290_Format(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "format"); signal != nil {
		return std_gui.PointerFromQSurfaceFormat(signal.(func() *std_gui.QSurfaceFormat)())
	}

	return std_gui.PointerFromQSurfaceFormat(NewRasterWindowFromPointer(ptr).FormatDefault())
}

func (ptr *RasterWindow) FormatDefault() *std_gui.QSurfaceFormat {
	if ptr.Pointer() != nil {
		tmpValue := std_gui.NewQSurfaceFormatFromPointer(C.RasterWindow000290_FormatDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_gui.QSurfaceFormat).DestroyQSurfaceFormat)
		return tmpValue
	}
	return nil
}

//export callbackRasterWindow000290_EventFilter
func callbackRasterWindow000290_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewRasterWindowFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *RasterWindow) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.RasterWindow000290_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackRasterWindow000290_ChildEvent
func callbackRasterWindow000290_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewRasterWindowFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *RasterWindow) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackRasterWindow000290_ConnectNotify
func callbackRasterWindow000290_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewRasterWindowFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *RasterWindow) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackRasterWindow000290_CustomEvent
func callbackRasterWindow000290_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewRasterWindowFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *RasterWindow) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackRasterWindow000290_DeleteLater
func callbackRasterWindow000290_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewRasterWindowFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *RasterWindow) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackRasterWindow000290_Destroyed
func callbackRasterWindow000290_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackRasterWindow000290_DisconnectNotify
func callbackRasterWindow000290_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewRasterWindowFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *RasterWindow) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackRasterWindow000290_ObjectNameChanged
func callbackRasterWindow000290_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackRasterWindow000290_TimerEvent
func callbackRasterWindow000290_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewRasterWindowFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *RasterWindow) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.RasterWindow000290_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
