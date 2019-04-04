package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"net/http"
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
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

type Hostname_ITF interface {
	std_core.QObject_ITF
	Hostname_PTR() *Hostname
}

func (ptr *Hostname) Hostname_PTR() *Hostname {
	return ptr
}

func (ptr *Hostname) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Hostname) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromHostname(ptr Hostname_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Hostname_PTR().Pointer()
	}
	return nil
}

func NewHostnameFromPointer(ptr unsafe.Pointer) (n *Hostname) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Hostname)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Hostname:
			n = deduced

		case *std_core.QObject:
			n = &Hostname{QObject: *deduced}

		default:
			n = new(Hostname)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackHostnamed48130_Constructor
func callbackHostnamed48130_Constructor(ptr unsafe.Pointer) {
	this := NewHostnameFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackHostnamed48130_Find
func callbackHostnamed48130_Find(ptr unsafe.Pointer, hostname C.struct_Moc_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "find"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(hostname)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *Hostname) ConnectFind(f func(hostname string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "find"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "find", func(hostname string) bool {
				signal.(func(string) bool)(hostname)
				return f(hostname)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "find", f)
		}
	}
}

func (ptr *Hostname) DisconnectFind() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "find")
	}
}

func (ptr *Hostname) Find(hostname string) bool {
	if ptr.Pointer() != nil {
		var hostnameC *C.char
		if hostname != "" {
			hostnameC = C.CString(hostname)
			defer C.free(unsafe.Pointer(hostnameC))
		}
		return int8(C.Hostnamed48130_Find(ptr.Pointer(), C.struct_Moc_PackedString{data: hostnameC, len: C.longlong(len(hostname))})) != 0
	}
	return false
}

func Hostname_QRegisterMetaType() int {
	return int(int32(C.Hostnamed48130_Hostnamed48130_QRegisterMetaType()))
}

func (ptr *Hostname) QRegisterMetaType() int {
	return int(int32(C.Hostnamed48130_Hostnamed48130_QRegisterMetaType()))
}

func Hostname_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Hostnamed48130_Hostnamed48130_QRegisterMetaType2(typeNameC)))
}

func (ptr *Hostname) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Hostnamed48130_Hostnamed48130_QRegisterMetaType2(typeNameC)))
}

func Hostname_QmlRegisterType() int {
	return int(int32(C.Hostnamed48130_Hostnamed48130_QmlRegisterType()))
}

func (ptr *Hostname) QmlRegisterType() int {
	return int(int32(C.Hostnamed48130_Hostnamed48130_QmlRegisterType()))
}

func Hostname_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Hostnamed48130_Hostnamed48130_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Hostname) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Hostnamed48130_Hostnamed48130_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Hostname) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.Hostnamed48130___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Hostname) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Hostnamed48130___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Hostname) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Hostnamed48130___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Hostname) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Hostnamed48130___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Hostname) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Hostnamed48130___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Hostname) __findChildren_newList2() unsafe.Pointer {
	return C.Hostnamed48130___findChildren_newList2(ptr.Pointer())
}

func (ptr *Hostname) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Hostnamed48130___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Hostname) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Hostnamed48130___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Hostname) __findChildren_newList3() unsafe.Pointer {
	return C.Hostnamed48130___findChildren_newList3(ptr.Pointer())
}

func (ptr *Hostname) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Hostnamed48130___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Hostname) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Hostnamed48130___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Hostname) __findChildren_newList() unsafe.Pointer {
	return C.Hostnamed48130___findChildren_newList(ptr.Pointer())
}

func (ptr *Hostname) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Hostnamed48130___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Hostname) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Hostnamed48130___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Hostname) __children_newList() unsafe.Pointer {
	return C.Hostnamed48130___children_newList(ptr.Pointer())
}

func NewHostname(parent std_core.QObject_ITF) *Hostname {
	tmpValue := NewHostnameFromPointer(C.Hostnamed48130_NewHostname(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackHostnamed48130_DestroyHostname
func callbackHostnamed48130_DestroyHostname(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Hostname"); signal != nil {
		signal.(func())()
	} else {
		NewHostnameFromPointer(ptr).DestroyHostnameDefault()
	}
}

func (ptr *Hostname) ConnectDestroyHostname(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Hostname"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~Hostname", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Hostname", f)
		}
	}
}

func (ptr *Hostname) DisconnectDestroyHostname() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Hostname")
	}
}

func (ptr *Hostname) DestroyHostname() {
	if ptr.Pointer() != nil {
		C.Hostnamed48130_DestroyHostname(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Hostname) DestroyHostnameDefault() {
	if ptr.Pointer() != nil {
		C.Hostnamed48130_DestroyHostnameDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackHostnamed48130_Event
func callbackHostnamed48130_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHostnameFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Hostname) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Hostnamed48130_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackHostnamed48130_EventFilter
func callbackHostnamed48130_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHostnameFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Hostname) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Hostnamed48130_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackHostnamed48130_ChildEvent
func callbackHostnamed48130_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewHostnameFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Hostname) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Hostnamed48130_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackHostnamed48130_ConnectNotify
func callbackHostnamed48130_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewHostnameFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Hostname) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Hostnamed48130_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackHostnamed48130_CustomEvent
func callbackHostnamed48130_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewHostnameFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Hostname) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Hostnamed48130_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackHostnamed48130_DeleteLater
func callbackHostnamed48130_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewHostnameFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Hostname) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Hostnamed48130_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackHostnamed48130_Destroyed
func callbackHostnamed48130_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackHostnamed48130_DisconnectNotify
func callbackHostnamed48130_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewHostnameFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Hostname) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Hostnamed48130_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackHostnamed48130_ObjectNameChanged
func callbackHostnamed48130_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackHostnamed48130_TimerEvent
func callbackHostnamed48130_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewHostnameFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Hostname) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Hostnamed48130_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type Login_ITF interface {
	std_core.QObject_ITF
	Login_PTR() *Login
}

func (ptr *Login) Login_PTR() *Login {
	return ptr
}

func (ptr *Login) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Login) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromLogin(ptr Login_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Login_PTR().Pointer()
	}
	return nil
}

func NewLoginFromPointer(ptr unsafe.Pointer) (n *Login) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Login)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Login:
			n = deduced

		case *std_core.QObject:
			n = &Login{QObject: *deduced}

		default:
			n = new(Login)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackLogind48130_Constructor
func callbackLogind48130_Constructor(ptr unsafe.Pointer) {
	this := NewLoginFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackLogind48130_Logind48130
func callbackLogind48130_Logind48130(ptr unsafe.Pointer, username C.struct_Moc_PackedString, password C.struct_Moc_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "login"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string) bool)(cGoUnpackString(username), cGoUnpackString(password)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *Login) ConnectLogin(f func(username string, password string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "login"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "login", func(username string, password string) bool {
				signal.(func(string, string) bool)(username, password)
				return f(username, password)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "login", f)
		}
	}
}

func (ptr *Login) DisconnectLogin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "login")
	}
}

func (ptr *Login) Login(username string, password string) bool {
	if ptr.Pointer() != nil {
		var usernameC *C.char
		if username != "" {
			usernameC = C.CString(username)
			defer C.free(unsafe.Pointer(usernameC))
		}
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		return int8(C.Logind48130_Logind48130(ptr.Pointer(), C.struct_Moc_PackedString{data: usernameC, len: C.longlong(len(username))}, C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})) != 0
	}
	return false
}

func Login_QRegisterMetaType() int {
	return int(int32(C.Logind48130_Logind48130_QRegisterMetaType()))
}

func (ptr *Login) QRegisterMetaType() int {
	return int(int32(C.Logind48130_Logind48130_QRegisterMetaType()))
}

func Login_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Logind48130_Logind48130_QRegisterMetaType2(typeNameC)))
}

func (ptr *Login) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Logind48130_Logind48130_QRegisterMetaType2(typeNameC)))
}

func Login_QmlRegisterType() int {
	return int(int32(C.Logind48130_Logind48130_QmlRegisterType()))
}

func (ptr *Login) QmlRegisterType() int {
	return int(int32(C.Logind48130_Logind48130_QmlRegisterType()))
}

func Login_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Logind48130_Logind48130_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Login) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Logind48130_Logind48130_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Login) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.Logind48130___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Login) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Logind48130___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Login) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Logind48130___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Login) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Logind48130___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Login) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Logind48130___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Login) __findChildren_newList2() unsafe.Pointer {
	return C.Logind48130___findChildren_newList2(ptr.Pointer())
}

func (ptr *Login) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Logind48130___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Login) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Logind48130___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Login) __findChildren_newList3() unsafe.Pointer {
	return C.Logind48130___findChildren_newList3(ptr.Pointer())
}

func (ptr *Login) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Logind48130___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Login) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Logind48130___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Login) __findChildren_newList() unsafe.Pointer {
	return C.Logind48130___findChildren_newList(ptr.Pointer())
}

func (ptr *Login) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Logind48130___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Login) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Logind48130___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Login) __children_newList() unsafe.Pointer {
	return C.Logind48130___children_newList(ptr.Pointer())
}

func NewLogin(parent std_core.QObject_ITF) *Login {
	tmpValue := NewLoginFromPointer(C.Logind48130_NewLogin(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackLogind48130_DestroyLogin
func callbackLogind48130_DestroyLogin(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Login"); signal != nil {
		signal.(func())()
	} else {
		NewLoginFromPointer(ptr).DestroyLoginDefault()
	}
}

func (ptr *Login) ConnectDestroyLogin(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Login"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~Login", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Login", f)
		}
	}
}

func (ptr *Login) DisconnectDestroyLogin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Login")
	}
}

func (ptr *Login) DestroyLogin() {
	if ptr.Pointer() != nil {
		C.Logind48130_DestroyLogin(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Login) DestroyLoginDefault() {
	if ptr.Pointer() != nil {
		C.Logind48130_DestroyLoginDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackLogind48130_Event
func callbackLogind48130_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewLoginFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Login) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Logind48130_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackLogind48130_EventFilter
func callbackLogind48130_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewLoginFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Login) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Logind48130_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackLogind48130_ChildEvent
func callbackLogind48130_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewLoginFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Login) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Logind48130_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackLogind48130_ConnectNotify
func callbackLogind48130_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewLoginFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Login) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Logind48130_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackLogind48130_CustomEvent
func callbackLogind48130_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewLoginFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Login) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Logind48130_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackLogind48130_DeleteLater
func callbackLogind48130_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewLoginFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Login) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Logind48130_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackLogind48130_Destroyed
func callbackLogind48130_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackLogind48130_DisconnectNotify
func callbackLogind48130_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewLoginFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Login) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Logind48130_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackLogind48130_ObjectNameChanged
func callbackLogind48130_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackLogind48130_TimerEvent
func callbackLogind48130_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewLoginFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Login) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Logind48130_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type QuatrixFile_ITF interface {
	std_core.QObject_ITF
	QuatrixFile_PTR() *QuatrixFile
}

func (ptr *QuatrixFile) QuatrixFile_PTR() *QuatrixFile {
	return ptr
}

func (ptr *QuatrixFile) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QuatrixFile) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQuatrixFile(ptr QuatrixFile_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QuatrixFile_PTR().Pointer()
	}
	return nil
}

func NewQuatrixFileFromPointer(ptr unsafe.Pointer) (n *QuatrixFile) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QuatrixFile)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QuatrixFile:
			n = deduced

		case *std_core.QObject:
			n = &QuatrixFile{QObject: *deduced}

		default:
			n = new(QuatrixFile)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackQuatrixFiled48130_Constructor
func callbackQuatrixFiled48130_Constructor(ptr unsafe.Pointer) {
	this := NewQuatrixFileFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackQuatrixFiled48130_FileName
func callbackQuatrixFiled48130_FileName(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "fileName"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQuatrixFileFromPointer(ptr).FileNameDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QuatrixFile) ConnectFileName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fileName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fileName", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fileName", f)
		}
	}
}

func (ptr *QuatrixFile) DisconnectFileName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fileName")
	}
}

func (ptr *QuatrixFile) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QuatrixFiled48130_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QuatrixFile) FileNameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QuatrixFiled48130_FileNameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQuatrixFiled48130_SetFileName
func callbackQuatrixFiled48130_SetFileName(ptr unsafe.Pointer, fileName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setFileName"); signal != nil {
		signal.(func(string))(cGoUnpackString(fileName))
	} else {
		NewQuatrixFileFromPointer(ptr).SetFileNameDefault(cGoUnpackString(fileName))
	}
}

func (ptr *QuatrixFile) ConnectSetFileName(f func(fileName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFileName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFileName", func(fileName string) {
				signal.(func(string))(fileName)
				f(fileName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFileName", f)
		}
	}
}

func (ptr *QuatrixFile) DisconnectSetFileName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFileName")
	}
}

func (ptr *QuatrixFile) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QuatrixFiled48130_SetFileName(ptr.Pointer(), C.struct_Moc_PackedString{data: fileNameC, len: C.longlong(len(fileName))})
	}
}

func (ptr *QuatrixFile) SetFileNameDefault(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QuatrixFiled48130_SetFileNameDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: fileNameC, len: C.longlong(len(fileName))})
	}
}

//export callbackQuatrixFiled48130_FileNameChanged
func callbackQuatrixFiled48130_FileNameChanged(ptr unsafe.Pointer, fileName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "fileNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(fileName))
	}

}

func (ptr *QuatrixFile) ConnectFileNameChanged(f func(fileName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fileNameChanged") {
			C.QuatrixFiled48130_ConnectFileNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "fileNameChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fileNameChanged", func(fileName string) {
				signal.(func(string))(fileName)
				f(fileName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fileNameChanged", f)
		}
	}
}

func (ptr *QuatrixFile) DisconnectFileNameChanged() {
	if ptr.Pointer() != nil {
		C.QuatrixFiled48130_DisconnectFileNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "fileNameChanged")
	}
}

func (ptr *QuatrixFile) FileNameChanged(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QuatrixFiled48130_FileNameChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: fileNameC, len: C.longlong(len(fileName))})
	}
}

//export callbackQuatrixFiled48130_MTime
func callbackQuatrixFiled48130_MTime(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mTime"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQuatrixFileFromPointer(ptr).MTimeDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QuatrixFile) ConnectMTime(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mTime"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "mTime", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mTime", f)
		}
	}
}

func (ptr *QuatrixFile) DisconnectMTime() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mTime")
	}
}

func (ptr *QuatrixFile) MTime() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QuatrixFiled48130_MTime(ptr.Pointer()))
	}
	return ""
}

func (ptr *QuatrixFile) MTimeDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QuatrixFiled48130_MTimeDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQuatrixFiled48130_SetMTime
func callbackQuatrixFiled48130_SetMTime(ptr unsafe.Pointer, mTime C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setMTime"); signal != nil {
		signal.(func(string))(cGoUnpackString(mTime))
	} else {
		NewQuatrixFileFromPointer(ptr).SetMTimeDefault(cGoUnpackString(mTime))
	}
}

func (ptr *QuatrixFile) ConnectSetMTime(f func(mTime string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setMTime"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setMTime", func(mTime string) {
				signal.(func(string))(mTime)
				f(mTime)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setMTime", f)
		}
	}
}

func (ptr *QuatrixFile) DisconnectSetMTime() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setMTime")
	}
}

func (ptr *QuatrixFile) SetMTime(mTime string) {
	if ptr.Pointer() != nil {
		var mTimeC *C.char
		if mTime != "" {
			mTimeC = C.CString(mTime)
			defer C.free(unsafe.Pointer(mTimeC))
		}
		C.QuatrixFiled48130_SetMTime(ptr.Pointer(), C.struct_Moc_PackedString{data: mTimeC, len: C.longlong(len(mTime))})
	}
}

func (ptr *QuatrixFile) SetMTimeDefault(mTime string) {
	if ptr.Pointer() != nil {
		var mTimeC *C.char
		if mTime != "" {
			mTimeC = C.CString(mTime)
			defer C.free(unsafe.Pointer(mTimeC))
		}
		C.QuatrixFiled48130_SetMTimeDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: mTimeC, len: C.longlong(len(mTime))})
	}
}

//export callbackQuatrixFiled48130_MTimeChanged
func callbackQuatrixFiled48130_MTimeChanged(ptr unsafe.Pointer, mTime C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "mTimeChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(mTime))
	}

}

func (ptr *QuatrixFile) ConnectMTimeChanged(f func(mTime string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "mTimeChanged") {
			C.QuatrixFiled48130_ConnectMTimeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "mTimeChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "mTimeChanged", func(mTime string) {
				signal.(func(string))(mTime)
				f(mTime)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mTimeChanged", f)
		}
	}
}

func (ptr *QuatrixFile) DisconnectMTimeChanged() {
	if ptr.Pointer() != nil {
		C.QuatrixFiled48130_DisconnectMTimeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "mTimeChanged")
	}
}

func (ptr *QuatrixFile) MTimeChanged(mTime string) {
	if ptr.Pointer() != nil {
		var mTimeC *C.char
		if mTime != "" {
			mTimeC = C.CString(mTime)
			defer C.free(unsafe.Pointer(mTimeC))
		}
		C.QuatrixFiled48130_MTimeChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: mTimeC, len: C.longlong(len(mTime))})
	}
}

func QuatrixFile_QRegisterMetaType() int {
	return int(int32(C.QuatrixFiled48130_QuatrixFiled48130_QRegisterMetaType()))
}

func (ptr *QuatrixFile) QRegisterMetaType() int {
	return int(int32(C.QuatrixFiled48130_QuatrixFiled48130_QRegisterMetaType()))
}

func QuatrixFile_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QuatrixFiled48130_QuatrixFiled48130_QRegisterMetaType2(typeNameC)))
}

func (ptr *QuatrixFile) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QuatrixFiled48130_QuatrixFiled48130_QRegisterMetaType2(typeNameC)))
}

func QuatrixFile_QmlRegisterType() int {
	return int(int32(C.QuatrixFiled48130_QuatrixFiled48130_QmlRegisterType()))
}

func (ptr *QuatrixFile) QmlRegisterType() int {
	return int(int32(C.QuatrixFiled48130_QuatrixFiled48130_QmlRegisterType()))
}

func QuatrixFile_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QuatrixFiled48130_QuatrixFiled48130_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QuatrixFile) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QuatrixFiled48130_QuatrixFiled48130_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QuatrixFile) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.QuatrixFiled48130___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QuatrixFile) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QuatrixFiled48130___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QuatrixFile) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QuatrixFiled48130___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QuatrixFile) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QuatrixFiled48130___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QuatrixFile) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QuatrixFiled48130___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QuatrixFile) __findChildren_newList2() unsafe.Pointer {
	return C.QuatrixFiled48130___findChildren_newList2(ptr.Pointer())
}

func (ptr *QuatrixFile) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QuatrixFiled48130___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QuatrixFile) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QuatrixFiled48130___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QuatrixFile) __findChildren_newList3() unsafe.Pointer {
	return C.QuatrixFiled48130___findChildren_newList3(ptr.Pointer())
}

func (ptr *QuatrixFile) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QuatrixFiled48130___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QuatrixFile) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QuatrixFiled48130___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QuatrixFile) __findChildren_newList() unsafe.Pointer {
	return C.QuatrixFiled48130___findChildren_newList(ptr.Pointer())
}

func (ptr *QuatrixFile) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QuatrixFiled48130___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QuatrixFile) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QuatrixFiled48130___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QuatrixFile) __children_newList() unsafe.Pointer {
	return C.QuatrixFiled48130___children_newList(ptr.Pointer())
}

func NewQuatrixFile(parent std_core.QObject_ITF) *QuatrixFile {
	tmpValue := NewQuatrixFileFromPointer(C.QuatrixFiled48130_NewQuatrixFile(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQuatrixFiled48130_DestroyQuatrixFile
func callbackQuatrixFiled48130_DestroyQuatrixFile(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QuatrixFile"); signal != nil {
		signal.(func())()
	} else {
		NewQuatrixFileFromPointer(ptr).DestroyQuatrixFileDefault()
	}
}

func (ptr *QuatrixFile) ConnectDestroyQuatrixFile(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QuatrixFile"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QuatrixFile", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QuatrixFile", f)
		}
	}
}

func (ptr *QuatrixFile) DisconnectDestroyQuatrixFile() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QuatrixFile")
	}
}

func (ptr *QuatrixFile) DestroyQuatrixFile() {
	if ptr.Pointer() != nil {
		C.QuatrixFiled48130_DestroyQuatrixFile(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QuatrixFile) DestroyQuatrixFileDefault() {
	if ptr.Pointer() != nil {
		C.QuatrixFiled48130_DestroyQuatrixFileDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQuatrixFiled48130_Event
func callbackQuatrixFiled48130_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQuatrixFileFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QuatrixFile) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QuatrixFiled48130_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQuatrixFiled48130_EventFilter
func callbackQuatrixFiled48130_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQuatrixFileFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QuatrixFile) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QuatrixFiled48130_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQuatrixFiled48130_ChildEvent
func callbackQuatrixFiled48130_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQuatrixFileFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QuatrixFile) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QuatrixFiled48130_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQuatrixFiled48130_ConnectNotify
func callbackQuatrixFiled48130_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQuatrixFileFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QuatrixFile) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QuatrixFiled48130_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQuatrixFiled48130_CustomEvent
func callbackQuatrixFiled48130_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewQuatrixFileFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QuatrixFile) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QuatrixFiled48130_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQuatrixFiled48130_DeleteLater
func callbackQuatrixFiled48130_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQuatrixFileFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QuatrixFile) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QuatrixFiled48130_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQuatrixFiled48130_Destroyed
func callbackQuatrixFiled48130_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackQuatrixFiled48130_DisconnectNotify
func callbackQuatrixFiled48130_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQuatrixFileFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QuatrixFile) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QuatrixFiled48130_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQuatrixFiled48130_ObjectNameChanged
func callbackQuatrixFiled48130_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQuatrixFiled48130_TimerEvent
func callbackQuatrixFiled48130_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQuatrixFileFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QuatrixFile) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QuatrixFiled48130_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type State_ITF interface {
	std_core.QObject_ITF
	State_PTR() *State
}

func (ptr *State) State_PTR() *State {
	return ptr
}

func (ptr *State) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *State) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromState(ptr State_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.State_PTR().Pointer()
	}
	return nil
}

func NewStateFromPointer(ptr unsafe.Pointer) (n *State) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(State)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *State:
			n = deduced

		case *std_core.QObject:
			n = &State{QObject: *deduced}

		default:
			n = new(State)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackStated48130_Constructor
func callbackStated48130_Constructor(ptr unsafe.Pointer) {
	this := NewStateFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackStated48130_Change
func callbackStated48130_Change(ptr unsafe.Pointer, hostname C.struct_Moc_PackedString, email C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "change"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(hostname), cGoUnpackString(email))
	}

}

func (ptr *State) ConnectChange(f func(hostname string, email string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "change") {
			C.Stated48130_ConnectChange(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "change"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "change", func(hostname string, email string) {
				signal.(func(string, string))(hostname, email)
				f(hostname, email)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "change", f)
		}
	}
}

func (ptr *State) DisconnectChange() {
	if ptr.Pointer() != nil {
		C.Stated48130_DisconnectChange(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "change")
	}
}

func (ptr *State) Change(hostname string, email string) {
	if ptr.Pointer() != nil {
		var hostnameC *C.char
		if hostname != "" {
			hostnameC = C.CString(hostname)
			defer C.free(unsafe.Pointer(hostnameC))
		}
		var emailC *C.char
		if email != "" {
			emailC = C.CString(email)
			defer C.free(unsafe.Pointer(emailC))
		}
		C.Stated48130_Change(ptr.Pointer(), C.struct_Moc_PackedString{data: hostnameC, len: C.longlong(len(hostname))}, C.struct_Moc_PackedString{data: emailC, len: C.longlong(len(email))})
	}
}

//export callbackStated48130_Hostnamed48130
func callbackStated48130_Hostnamed48130(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "hostname"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewStateFromPointer(ptr).HostnameDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *State) ConnectHostname(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hostname"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hostname", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hostname", f)
		}
	}
}

func (ptr *State) DisconnectHostname() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hostname")
	}
}

func (ptr *State) Hostname() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Stated48130_Hostnamed48130(ptr.Pointer()))
	}
	return ""
}

func (ptr *State) HostnameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Stated48130_HostnameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackStated48130_SetHostname
func callbackStated48130_SetHostname(ptr unsafe.Pointer, hostname C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setHostname"); signal != nil {
		signal.(func(string))(cGoUnpackString(hostname))
	} else {
		NewStateFromPointer(ptr).SetHostnameDefault(cGoUnpackString(hostname))
	}
}

func (ptr *State) ConnectSetHostname(f func(hostname string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setHostname"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setHostname", func(hostname string) {
				signal.(func(string))(hostname)
				f(hostname)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setHostname", f)
		}
	}
}

func (ptr *State) DisconnectSetHostname() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setHostname")
	}
}

func (ptr *State) SetHostname(hostname string) {
	if ptr.Pointer() != nil {
		var hostnameC *C.char
		if hostname != "" {
			hostnameC = C.CString(hostname)
			defer C.free(unsafe.Pointer(hostnameC))
		}
		C.Stated48130_SetHostname(ptr.Pointer(), C.struct_Moc_PackedString{data: hostnameC, len: C.longlong(len(hostname))})
	}
}

func (ptr *State) SetHostnameDefault(hostname string) {
	if ptr.Pointer() != nil {
		var hostnameC *C.char
		if hostname != "" {
			hostnameC = C.CString(hostname)
			defer C.free(unsafe.Pointer(hostnameC))
		}
		C.Stated48130_SetHostnameDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: hostnameC, len: C.longlong(len(hostname))})
	}
}

//export callbackStated48130_HostnameChanged
func callbackStated48130_HostnameChanged(ptr unsafe.Pointer, hostname C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "hostnameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(hostname))
	}

}

func (ptr *State) ConnectHostnameChanged(f func(hostname string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hostnameChanged") {
			C.Stated48130_ConnectHostnameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hostnameChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hostnameChanged", func(hostname string) {
				signal.(func(string))(hostname)
				f(hostname)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hostnameChanged", f)
		}
	}
}

func (ptr *State) DisconnectHostnameChanged() {
	if ptr.Pointer() != nil {
		C.Stated48130_DisconnectHostnameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hostnameChanged")
	}
}

func (ptr *State) HostnameChanged(hostname string) {
	if ptr.Pointer() != nil {
		var hostnameC *C.char
		if hostname != "" {
			hostnameC = C.CString(hostname)
			defer C.free(unsafe.Pointer(hostnameC))
		}
		C.Stated48130_HostnameChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: hostnameC, len: C.longlong(len(hostname))})
	}
}

//export callbackStated48130_Email
func callbackStated48130_Email(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "email"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewStateFromPointer(ptr).EmailDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *State) ConnectEmail(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "email"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "email", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "email", f)
		}
	}
}

func (ptr *State) DisconnectEmail() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "email")
	}
}

func (ptr *State) Email() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Stated48130_Email(ptr.Pointer()))
	}
	return ""
}

func (ptr *State) EmailDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Stated48130_EmailDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackStated48130_SetEmail
func callbackStated48130_SetEmail(ptr unsafe.Pointer, email C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setEmail"); signal != nil {
		signal.(func(string))(cGoUnpackString(email))
	} else {
		NewStateFromPointer(ptr).SetEmailDefault(cGoUnpackString(email))
	}
}

func (ptr *State) ConnectSetEmail(f func(email string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setEmail"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setEmail", func(email string) {
				signal.(func(string))(email)
				f(email)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setEmail", f)
		}
	}
}

func (ptr *State) DisconnectSetEmail() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEmail")
	}
}

func (ptr *State) SetEmail(email string) {
	if ptr.Pointer() != nil {
		var emailC *C.char
		if email != "" {
			emailC = C.CString(email)
			defer C.free(unsafe.Pointer(emailC))
		}
		C.Stated48130_SetEmail(ptr.Pointer(), C.struct_Moc_PackedString{data: emailC, len: C.longlong(len(email))})
	}
}

func (ptr *State) SetEmailDefault(email string) {
	if ptr.Pointer() != nil {
		var emailC *C.char
		if email != "" {
			emailC = C.CString(email)
			defer C.free(unsafe.Pointer(emailC))
		}
		C.Stated48130_SetEmailDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: emailC, len: C.longlong(len(email))})
	}
}

//export callbackStated48130_EmailChanged
func callbackStated48130_EmailChanged(ptr unsafe.Pointer, email C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "emailChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(email))
	}

}

func (ptr *State) ConnectEmailChanged(f func(email string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "emailChanged") {
			C.Stated48130_ConnectEmailChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "emailChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "emailChanged", func(email string) {
				signal.(func(string))(email)
				f(email)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "emailChanged", f)
		}
	}
}

func (ptr *State) DisconnectEmailChanged() {
	if ptr.Pointer() != nil {
		C.Stated48130_DisconnectEmailChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "emailChanged")
	}
}

func (ptr *State) EmailChanged(email string) {
	if ptr.Pointer() != nil {
		var emailC *C.char
		if email != "" {
			emailC = C.CString(email)
			defer C.free(unsafe.Pointer(emailC))
		}
		C.Stated48130_EmailChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: emailC, len: C.longlong(len(email))})
	}
}

//export callbackStated48130_Cookie
func callbackStated48130_Cookie(ptr unsafe.Pointer) C.uintptr_t {
	if signal := qt.GetSignal(ptr, "cookie"); signal != nil {
		oP := signal.(func() *http.Cookie)()
		qt.RegisterTemp(unsafe.Pointer(oP), oP)
		return C.uintptr_t(uintptr(unsafe.Pointer(oP)))
	}
	oP := NewStateFromPointer(ptr).CookieDefault()
	qt.RegisterTemp(unsafe.Pointer(oP), oP)
	return C.uintptr_t(uintptr(unsafe.Pointer(oP)))
}

func (ptr *State) ConnectCookie(f func() *http.Cookie) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "cookie"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "cookie", func() *http.Cookie {
				signal.(func() *http.Cookie)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cookie", f)
		}
	}
}

func (ptr *State) DisconnectCookie() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "cookie")
	}
}

func (ptr *State) Cookie() *http.Cookie {
	if ptr.Pointer() != nil {
		oP := unsafe.Pointer(uintptr(C.Stated48130_Cookie(ptr.Pointer())))
		if oI, ok := qt.ReceiveTemp(oP); ok {
			oD := oI.(*http.Cookie)
			return oD
		}

	}
	var out *http.Cookie
	return out
}

func (ptr *State) CookieDefault() *http.Cookie {
	if ptr.Pointer() != nil {
		oP := unsafe.Pointer(uintptr(C.Stated48130_CookieDefault(ptr.Pointer())))
		if oI, ok := qt.ReceiveTemp(oP); ok {
			oD := oI.(*http.Cookie)
			return oD
		}

	}
	var out *http.Cookie
	return out
}

//export callbackStated48130_SetCookie
func callbackStated48130_SetCookie(ptr unsafe.Pointer, cookie C.uintptr_t) {
	var cookieD *http.Cookie
	if cookieI, ok := qt.ReceiveTemp(unsafe.Pointer(uintptr(cookie))); ok {
		qt.UnregisterTemp(unsafe.Pointer(uintptr(cookie)))
		cookieD = cookieI.(*http.Cookie)
	}
	if signal := qt.GetSignal(ptr, "setCookie"); signal != nil {
		signal.(func(*http.Cookie))(cookieD)
	} else {
		NewStateFromPointer(ptr).SetCookieDefault(cookieD)
	}
}

func (ptr *State) ConnectSetCookie(f func(cookie *http.Cookie)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setCookie"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setCookie", func(cookie *http.Cookie) {
				signal.(func(*http.Cookie))(cookie)
				f(cookie)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setCookie", f)
		}
	}
}

func (ptr *State) DisconnectSetCookie() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setCookie")
	}
}

func (ptr *State) SetCookie(cookie *http.Cookie) {
	if ptr.Pointer() != nil {
		qt.RegisterTemp(unsafe.Pointer(cookie), cookie)
		C.Stated48130_SetCookie(ptr.Pointer(), C.uintptr_t(uintptr(unsafe.Pointer(cookie))))
	}
}

func (ptr *State) SetCookieDefault(cookie *http.Cookie) {
	if ptr.Pointer() != nil {
		qt.RegisterTemp(unsafe.Pointer(cookie), cookie)
		C.Stated48130_SetCookieDefault(ptr.Pointer(), C.uintptr_t(uintptr(unsafe.Pointer(cookie))))
	}
}

//export callbackStated48130_CookieChanged
func callbackStated48130_CookieChanged(ptr unsafe.Pointer, cookie C.uintptr_t) {
	var cookieD *http.Cookie
	if cookieI, ok := qt.ReceiveTemp(unsafe.Pointer(uintptr(cookie))); ok {
		cookieD = cookieI.(*http.Cookie)
	}
	if signal := qt.GetSignal(ptr, "cookieChanged"); signal != nil {
		signal.(func(*http.Cookie))(cookieD)
	}

}

func (ptr *State) ConnectCookieChanged(f func(cookie *http.Cookie)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "cookieChanged") {
			C.Stated48130_ConnectCookieChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "cookieChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "cookieChanged", func(cookie *http.Cookie) {
				signal.(func(*http.Cookie))(cookie)
				f(cookie)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cookieChanged", f)
		}
	}
}

func (ptr *State) DisconnectCookieChanged() {
	if ptr.Pointer() != nil {
		C.Stated48130_DisconnectCookieChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "cookieChanged")
	}
}

func (ptr *State) CookieChanged(cookie *http.Cookie) {
	if ptr.Pointer() != nil {
		qt.RegisterTemp(unsafe.Pointer(cookie), cookie)
		C.Stated48130_CookieChanged(ptr.Pointer(), C.uintptr_t(uintptr(unsafe.Pointer(cookie))))
	}
}

func State_QRegisterMetaType() int {
	return int(int32(C.Stated48130_Stated48130_QRegisterMetaType()))
}

func (ptr *State) QRegisterMetaType() int {
	return int(int32(C.Stated48130_Stated48130_QRegisterMetaType()))
}

func State_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Stated48130_Stated48130_QRegisterMetaType2(typeNameC)))
}

func (ptr *State) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Stated48130_Stated48130_QRegisterMetaType2(typeNameC)))
}

func State_QmlRegisterType() int {
	return int(int32(C.Stated48130_Stated48130_QmlRegisterType()))
}

func (ptr *State) QmlRegisterType() int {
	return int(int32(C.Stated48130_Stated48130_QmlRegisterType()))
}

func State_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Stated48130_Stated48130_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *State) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Stated48130_Stated48130_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *State) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.Stated48130___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *State) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Stated48130___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *State) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Stated48130___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *State) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Stated48130___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *State) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Stated48130___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *State) __findChildren_newList2() unsafe.Pointer {
	return C.Stated48130___findChildren_newList2(ptr.Pointer())
}

func (ptr *State) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Stated48130___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *State) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Stated48130___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *State) __findChildren_newList3() unsafe.Pointer {
	return C.Stated48130___findChildren_newList3(ptr.Pointer())
}

func (ptr *State) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Stated48130___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *State) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Stated48130___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *State) __findChildren_newList() unsafe.Pointer {
	return C.Stated48130___findChildren_newList(ptr.Pointer())
}

func (ptr *State) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Stated48130___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *State) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Stated48130___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *State) __children_newList() unsafe.Pointer {
	return C.Stated48130___children_newList(ptr.Pointer())
}

func NewState(parent std_core.QObject_ITF) *State {
	tmpValue := NewStateFromPointer(C.Stated48130_NewState(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackStated48130_DestroyState
func callbackStated48130_DestroyState(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~State"); signal != nil {
		signal.(func())()
	} else {
		NewStateFromPointer(ptr).DestroyStateDefault()
	}
}

func (ptr *State) ConnectDestroyState(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~State"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~State", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~State", f)
		}
	}
}

func (ptr *State) DisconnectDestroyState() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~State")
	}
}

func (ptr *State) DestroyState() {
	if ptr.Pointer() != nil {
		C.Stated48130_DestroyState(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *State) DestroyStateDefault() {
	if ptr.Pointer() != nil {
		C.Stated48130_DestroyStateDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackStated48130_Event
func callbackStated48130_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewStateFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *State) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Stated48130_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackStated48130_EventFilter
func callbackStated48130_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewStateFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *State) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Stated48130_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackStated48130_ChildEvent
func callbackStated48130_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewStateFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *State) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Stated48130_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackStated48130_ConnectNotify
func callbackStated48130_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewStateFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *State) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Stated48130_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackStated48130_CustomEvent
func callbackStated48130_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewStateFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *State) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Stated48130_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackStated48130_DeleteLater
func callbackStated48130_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewStateFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *State) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Stated48130_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackStated48130_Destroyed
func callbackStated48130_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackStated48130_DisconnectNotify
func callbackStated48130_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewStateFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *State) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Stated48130_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackStated48130_ObjectNameChanged
func callbackStated48130_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackStated48130_TimerEvent
func callbackStated48130_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewStateFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *State) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Stated48130_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type FileHelper_ITF interface {
	std_core.QObject_ITF
	FileHelper_PTR() *FileHelper
}

func (ptr *FileHelper) FileHelper_PTR() *FileHelper {
	return ptr
}

func (ptr *FileHelper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *FileHelper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromFileHelper(ptr FileHelper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.FileHelper_PTR().Pointer()
	}
	return nil
}

func NewFileHelperFromPointer(ptr unsafe.Pointer) (n *FileHelper) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(FileHelper)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *FileHelper:
			n = deduced

		case *std_core.QObject:
			n = &FileHelper{QObject: *deduced}

		default:
			n = new(FileHelper)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackFileHelperd48130_Constructor
func callbackFileHelperd48130_Constructor(ptr unsafe.Pointer) {
	this := NewFileHelperFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackFileHelperd48130_Metadata
func callbackFileHelperd48130_Metadata(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "metadata"); signal != nil {
		signal.(func())()
	}

}

func (ptr *FileHelper) ConnectMetadata(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "metadata"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "metadata", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "metadata", f)
		}
	}
}

func (ptr *FileHelper) DisconnectMetadata() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "metadata")
	}
}

func (ptr *FileHelper) Metadata() {
	if ptr.Pointer() != nil {
		C.FileHelperd48130_Metadata(ptr.Pointer())
	}
}

//export callbackFileHelperd48130_Changed
func callbackFileHelperd48130_Changed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *FileHelper) ConnectChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "changed") {
			C.FileHelperd48130_ConnectChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "changed"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "changed", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "changed", f)
		}
	}
}

func (ptr *FileHelper) DisconnectChanged() {
	if ptr.Pointer() != nil {
		C.FileHelperd48130_DisconnectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "changed")
	}
}

func (ptr *FileHelper) Changed() {
	if ptr.Pointer() != nil {
		C.FileHelperd48130_Changed(ptr.Pointer())
	}
}

func FileHelper_QRegisterMetaType() int {
	return int(int32(C.FileHelperd48130_FileHelperd48130_QRegisterMetaType()))
}

func (ptr *FileHelper) QRegisterMetaType() int {
	return int(int32(C.FileHelperd48130_FileHelperd48130_QRegisterMetaType()))
}

func FileHelper_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.FileHelperd48130_FileHelperd48130_QRegisterMetaType2(typeNameC)))
}

func (ptr *FileHelper) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.FileHelperd48130_FileHelperd48130_QRegisterMetaType2(typeNameC)))
}

func FileHelper_QmlRegisterType() int {
	return int(int32(C.FileHelperd48130_FileHelperd48130_QmlRegisterType()))
}

func (ptr *FileHelper) QmlRegisterType() int {
	return int(int32(C.FileHelperd48130_FileHelperd48130_QmlRegisterType()))
}

func FileHelper_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.FileHelperd48130_FileHelperd48130_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *FileHelper) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.FileHelperd48130_FileHelperd48130_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *FileHelper) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.FileHelperd48130___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *FileHelper) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.FileHelperd48130___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *FileHelper) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.FileHelperd48130___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *FileHelper) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.FileHelperd48130___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FileHelper) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FileHelperd48130___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *FileHelper) __findChildren_newList2() unsafe.Pointer {
	return C.FileHelperd48130___findChildren_newList2(ptr.Pointer())
}

func (ptr *FileHelper) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.FileHelperd48130___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FileHelper) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FileHelperd48130___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *FileHelper) __findChildren_newList3() unsafe.Pointer {
	return C.FileHelperd48130___findChildren_newList3(ptr.Pointer())
}

func (ptr *FileHelper) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.FileHelperd48130___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FileHelper) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FileHelperd48130___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *FileHelper) __findChildren_newList() unsafe.Pointer {
	return C.FileHelperd48130___findChildren_newList(ptr.Pointer())
}

func (ptr *FileHelper) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.FileHelperd48130___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FileHelper) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FileHelperd48130___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *FileHelper) __children_newList() unsafe.Pointer {
	return C.FileHelperd48130___children_newList(ptr.Pointer())
}

func NewFileHelper(parent std_core.QObject_ITF) *FileHelper {
	tmpValue := NewFileHelperFromPointer(C.FileHelperd48130_NewFileHelper(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackFileHelperd48130_DestroyFileHelper
func callbackFileHelperd48130_DestroyFileHelper(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~FileHelper"); signal != nil {
		signal.(func())()
	} else {
		NewFileHelperFromPointer(ptr).DestroyFileHelperDefault()
	}
}

func (ptr *FileHelper) ConnectDestroyFileHelper(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~FileHelper"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~FileHelper", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~FileHelper", f)
		}
	}
}

func (ptr *FileHelper) DisconnectDestroyFileHelper() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~FileHelper")
	}
}

func (ptr *FileHelper) DestroyFileHelper() {
	if ptr.Pointer() != nil {
		C.FileHelperd48130_DestroyFileHelper(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *FileHelper) DestroyFileHelperDefault() {
	if ptr.Pointer() != nil {
		C.FileHelperd48130_DestroyFileHelperDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackFileHelperd48130_Event
func callbackFileHelperd48130_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileHelperFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *FileHelper) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileHelperd48130_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackFileHelperd48130_EventFilter
func callbackFileHelperd48130_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileHelperFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *FileHelper) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileHelperd48130_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackFileHelperd48130_ChildEvent
func callbackFileHelperd48130_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewFileHelperFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *FileHelper) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.FileHelperd48130_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackFileHelperd48130_ConnectNotify
func callbackFileHelperd48130_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewFileHelperFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *FileHelper) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.FileHelperd48130_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackFileHelperd48130_CustomEvent
func callbackFileHelperd48130_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewFileHelperFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *FileHelper) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.FileHelperd48130_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackFileHelperd48130_DeleteLater
func callbackFileHelperd48130_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewFileHelperFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *FileHelper) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.FileHelperd48130_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackFileHelperd48130_Destroyed
func callbackFileHelperd48130_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackFileHelperd48130_DisconnectNotify
func callbackFileHelperd48130_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewFileHelperFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *FileHelper) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.FileHelperd48130_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackFileHelperd48130_ObjectNameChanged
func callbackFileHelperd48130_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackFileHelperd48130_TimerEvent
func callbackFileHelperd48130_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewFileHelperFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *FileHelper) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.FileHelperd48130_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type FileModel_ITF interface {
	std_core.QAbstractListModel_ITF
	FileModel_PTR() *FileModel
}

func (ptr *FileModel) FileModel_PTR() *FileModel {
	return ptr
}

func (ptr *FileModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *FileModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromFileModel(ptr FileModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.FileModel_PTR().Pointer()
	}
	return nil
}

func NewFileModelFromPointer(ptr unsafe.Pointer) (n *FileModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(FileModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *FileModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &FileModel{QAbstractListModel: *deduced}

		default:
			n = new(FileModel)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackFileModeld48130_Constructor
func callbackFileModeld48130_Constructor(ptr unsafe.Pointer) {
	this := NewFileModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackFileModeld48130_Roles
func callbackFileModeld48130_Roles(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roles"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__roles_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__roles_newList())
		for k, v := range NewFileModelFromPointer(ptr).RolesDefault() {
			tmpList.__roles_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *FileModel) ConnectRoles(f func() map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "roles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "roles", func() map[int]*std_core.QByteArray {
				signal.(func() map[int]*std_core.QByteArray)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "roles", f)
		}
	}
}

func (ptr *FileModel) DisconnectRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "roles")
	}
}

func (ptr *FileModel) Roles() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.FileModeld48130_Roles(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

func (ptr *FileModel) RolesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.FileModeld48130_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackFileModeld48130_SetRoles
func callbackFileModeld48130_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setRoles"); signal != nil {
		signal.(func(map[int]*std_core.QByteArray))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	} else {
		NewFileModelFromPointer(ptr).SetRolesDefault(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	}
}

func (ptr *FileModel) ConnectSetRoles(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRoles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", func(roles map[int]*std_core.QByteArray) {
				signal.(func(map[int]*std_core.QByteArray))(roles)
				f(roles)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", f)
		}
	}
}

func (ptr *FileModel) DisconnectSetRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRoles")
	}
}

func (ptr *FileModel) SetRoles(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *FileModel) SetRolesDefault(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackFileModeld48130_RolesChanged
func callbackFileModeld48130_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "rolesChanged"); signal != nil {
		signal.(func(map[int]*std_core.QByteArray))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i, v := range tmpList.__rolesChanged_roles_keyList() {
				out[v] = tmpList.__rolesChanged_roles_atList(v, i)
			}
			return out
		}(roles))
	}

}

func (ptr *FileModel) ConnectRolesChanged(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rolesChanged") {
			C.FileModeld48130_ConnectRolesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rolesChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", func(roles map[int]*std_core.QByteArray) {
				signal.(func(map[int]*std_core.QByteArray))(roles)
				f(roles)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", f)
		}
	}
}

func (ptr *FileModel) DisconnectRolesChanged() {
	if ptr.Pointer() != nil {
		C.FileModeld48130_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *FileModel) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackFileModeld48130_Files
func callbackFileModeld48130_Files(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "files"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__files_newList())
			for _, v := range signal.(func() []*QuatrixFile)() {
				tmpList.__files_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__files_newList())
		for _, v := range NewFileModelFromPointer(ptr).FilesDefault() {
			tmpList.__files_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *FileModel) ConnectFiles(f func() []*QuatrixFile) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "files"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "files", func() []*QuatrixFile {
				signal.(func() []*QuatrixFile)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "files", f)
		}
	}
}

func (ptr *FileModel) DisconnectFiles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "files")
	}
}

func (ptr *FileModel) Files() []*QuatrixFile {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*QuatrixFile {
			out := make([]*QuatrixFile, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__files_atList(i)
			}
			return out
		}(C.FileModeld48130_Files(ptr.Pointer()))
	}
	return make([]*QuatrixFile, 0)
}

func (ptr *FileModel) FilesDefault() []*QuatrixFile {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*QuatrixFile {
			out := make([]*QuatrixFile, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__files_atList(i)
			}
			return out
		}(C.FileModeld48130_FilesDefault(ptr.Pointer()))
	}
	return make([]*QuatrixFile, 0)
}

//export callbackFileModeld48130_SetFiles
func callbackFileModeld48130_SetFiles(ptr unsafe.Pointer, files C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setFiles"); signal != nil {
		signal.(func([]*QuatrixFile))(func(l C.struct_Moc_PackedList) []*QuatrixFile {
			out := make([]*QuatrixFile, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setFiles_files_atList(i)
			}
			return out
		}(files))
	} else {
		NewFileModelFromPointer(ptr).SetFilesDefault(func(l C.struct_Moc_PackedList) []*QuatrixFile {
			out := make([]*QuatrixFile, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setFiles_files_atList(i)
			}
			return out
		}(files))
	}
}

func (ptr *FileModel) ConnectSetFiles(f func(files []*QuatrixFile)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFiles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFiles", func(files []*QuatrixFile) {
				signal.(func([]*QuatrixFile))(files)
				f(files)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFiles", f)
		}
	}
}

func (ptr *FileModel) DisconnectSetFiles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFiles")
	}
}

func (ptr *FileModel) SetFiles(files []*QuatrixFile) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_SetFiles(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__setFiles_files_newList())
			for _, v := range files {
				tmpList.__setFiles_files_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *FileModel) SetFilesDefault(files []*QuatrixFile) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_SetFilesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__setFiles_files_newList())
			for _, v := range files {
				tmpList.__setFiles_files_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackFileModeld48130_FilesChanged
func callbackFileModeld48130_FilesChanged(ptr unsafe.Pointer, files C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "filesChanged"); signal != nil {
		signal.(func([]*QuatrixFile))(func(l C.struct_Moc_PackedList) []*QuatrixFile {
			out := make([]*QuatrixFile, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__filesChanged_files_atList(i)
			}
			return out
		}(files))
	}

}

func (ptr *FileModel) ConnectFilesChanged(f func(files []*QuatrixFile)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "filesChanged") {
			C.FileModeld48130_ConnectFilesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "filesChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "filesChanged", func(files []*QuatrixFile) {
				signal.(func([]*QuatrixFile))(files)
				f(files)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filesChanged", f)
		}
	}
}

func (ptr *FileModel) DisconnectFilesChanged() {
	if ptr.Pointer() != nil {
		C.FileModeld48130_DisconnectFilesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "filesChanged")
	}
}

func (ptr *FileModel) FilesChanged(files []*QuatrixFile) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_FilesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__filesChanged_files_newList())
			for _, v := range files {
				tmpList.__filesChanged_files_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func FileModel_QRegisterMetaType() int {
	return int(int32(C.FileModeld48130_FileModeld48130_QRegisterMetaType()))
}

func (ptr *FileModel) QRegisterMetaType() int {
	return int(int32(C.FileModeld48130_FileModeld48130_QRegisterMetaType()))
}

func FileModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.FileModeld48130_FileModeld48130_QRegisterMetaType2(typeNameC)))
}

func (ptr *FileModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.FileModeld48130_FileModeld48130_QRegisterMetaType2(typeNameC)))
}

func FileModel_QmlRegisterType() int {
	return int(int32(C.FileModeld48130_FileModeld48130_QmlRegisterType()))
}

func (ptr *FileModel) QmlRegisterType() int {
	return int(int32(C.FileModeld48130_FileModeld48130_QmlRegisterType()))
}

func FileModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.FileModeld48130_FileModeld48130_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *FileModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.FileModeld48130_FileModeld48130_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *FileModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.FileModeld48130_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *FileModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *FileModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.FileModeld48130_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *FileModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.FileModeld48130_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *FileModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *FileModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.FileModeld48130_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *FileModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.FileModeld48130_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *FileModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *FileModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.FileModeld48130_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *FileModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.FileModeld48130___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *FileModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.FileModeld48130___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *FileModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.FileModeld48130___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *FileModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.FileModeld48130___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *FileModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.FileModeld48130___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *FileModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.FileModeld48130___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *FileModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.FileModeld48130___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *FileModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.FileModeld48130___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *FileModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *FileModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.FileModeld48130___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *FileModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.FileModeld48130___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *FileModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.FileModeld48130___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *FileModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.FileModeld48130___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *FileModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.FileModeld48130___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *FileModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.FileModeld48130___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *FileModel) __roleNames_newList() unsafe.Pointer {
	return C.FileModeld48130___roleNames_newList(ptr.Pointer())
}

func (ptr *FileModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.FileModeld48130___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *FileModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.FileModeld48130___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *FileModel) __itemData_newList() unsafe.Pointer {
	return C.FileModeld48130___itemData_newList(ptr.Pointer())
}

func (ptr *FileModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.FileModeld48130___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *FileModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.FileModeld48130___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *FileModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.FileModeld48130___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *FileModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.FileModeld48130___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *FileModel) __match_newList() unsafe.Pointer {
	return C.FileModeld48130___match_newList(ptr.Pointer())
}

func (ptr *FileModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.FileModeld48130___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *FileModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.FileModeld48130___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *FileModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.FileModeld48130_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *FileModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *FileModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.FileModeld48130_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *FileModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.FileModeld48130_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *FileModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *FileModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.FileModeld48130_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *FileModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.FileModeld48130___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *FileModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.FileModeld48130___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *FileModel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.FileModeld48130___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *FileModel) __findChildren_newList2() unsafe.Pointer {
	return C.FileModeld48130___findChildren_newList2(ptr.Pointer())
}

func (ptr *FileModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.FileModeld48130___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *FileModel) __findChildren_newList3() unsafe.Pointer {
	return C.FileModeld48130___findChildren_newList3(ptr.Pointer())
}

func (ptr *FileModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.FileModeld48130___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *FileModel) __findChildren_newList() unsafe.Pointer {
	return C.FileModeld48130___findChildren_newList(ptr.Pointer())
}

func (ptr *FileModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.FileModeld48130___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *FileModel) __children_newList() unsafe.Pointer {
	return C.FileModeld48130___children_newList(ptr.Pointer())
}

func (ptr *FileModel) __roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.FileModeld48130___roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *FileModel) __roles_newList() unsafe.Pointer {
	return C.FileModeld48130___roles_newList(ptr.Pointer())
}

func (ptr *FileModel) __roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roles_keyList_atList(i)
			}
			return out
		}(C.FileModeld48130___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *FileModel) __setRoles_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.FileModeld48130___setRoles_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *FileModel) __setRoles_roles_newList() unsafe.Pointer {
	return C.FileModeld48130___setRoles_roles_newList(ptr.Pointer())
}

func (ptr *FileModel) __setRoles_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setRoles_roles_keyList_atList(i)
			}
			return out
		}(C.FileModeld48130___setRoles_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *FileModel) __rolesChanged_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.FileModeld48130___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *FileModel) __rolesChanged_roles_newList() unsafe.Pointer {
	return C.FileModeld48130___rolesChanged_roles_newList(ptr.Pointer())
}

func (ptr *FileModel) __rolesChanged_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____rolesChanged_roles_keyList_atList(i)
			}
			return out
		}(C.FileModeld48130___rolesChanged_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *FileModel) __files_atList(i int) *QuatrixFile {
	if ptr.Pointer() != nil {
		tmpValue := NewQuatrixFileFromPointer(C.FileModeld48130___files_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __files_setList(i QuatrixFile_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___files_setList(ptr.Pointer(), PointerFromQuatrixFile(i))
	}
}

func (ptr *FileModel) __files_newList() unsafe.Pointer {
	return C.FileModeld48130___files_newList(ptr.Pointer())
}

func (ptr *FileModel) __setFiles_files_atList(i int) *QuatrixFile {
	if ptr.Pointer() != nil {
		tmpValue := NewQuatrixFileFromPointer(C.FileModeld48130___setFiles_files_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __setFiles_files_setList(i QuatrixFile_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___setFiles_files_setList(ptr.Pointer(), PointerFromQuatrixFile(i))
	}
}

func (ptr *FileModel) __setFiles_files_newList() unsafe.Pointer {
	return C.FileModeld48130___setFiles_files_newList(ptr.Pointer())
}

func (ptr *FileModel) __filesChanged_files_atList(i int) *QuatrixFile {
	if ptr.Pointer() != nil {
		tmpValue := NewQuatrixFileFromPointer(C.FileModeld48130___filesChanged_files_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FileModel) __filesChanged_files_setList(i QuatrixFile_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130___filesChanged_files_setList(ptr.Pointer(), PointerFromQuatrixFile(i))
	}
}

func (ptr *FileModel) __filesChanged_files_newList() unsafe.Pointer {
	return C.FileModeld48130___filesChanged_files_newList(ptr.Pointer())
}

func (ptr *FileModel) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.FileModeld48130_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *FileModel) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *FileModel) ____roles_keyList_newList() unsafe.Pointer {
	return C.FileModeld48130_____roles_keyList_newList(ptr.Pointer())
}

func (ptr *FileModel) ____setRoles_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.FileModeld48130_____setRoles_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *FileModel) ____setRoles_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_____setRoles_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *FileModel) ____setRoles_roles_keyList_newList() unsafe.Pointer {
	return C.FileModeld48130_____setRoles_roles_keyList_newList(ptr.Pointer())
}

func (ptr *FileModel) ____rolesChanged_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.FileModeld48130_____rolesChanged_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *FileModel) ____rolesChanged_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_____rolesChanged_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *FileModel) ____rolesChanged_roles_keyList_newList() unsafe.Pointer {
	return C.FileModeld48130_____rolesChanged_roles_keyList_newList(ptr.Pointer())
}

func NewFileModel(parent std_core.QObject_ITF) *FileModel {
	tmpValue := NewFileModelFromPointer(C.FileModeld48130_NewFileModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackFileModeld48130_DestroyFileModel
func callbackFileModeld48130_DestroyFileModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~FileModel"); signal != nil {
		signal.(func())()
	} else {
		NewFileModelFromPointer(ptr).DestroyFileModelDefault()
	}
}

func (ptr *FileModel) ConnectDestroyFileModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~FileModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~FileModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~FileModel", f)
		}
	}
}

func (ptr *FileModel) DisconnectDestroyFileModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~FileModel")
	}
}

func (ptr *FileModel) DestroyFileModel() {
	if ptr.Pointer() != nil {
		C.FileModeld48130_DestroyFileModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *FileModel) DestroyFileModelDefault() {
	if ptr.Pointer() != nil {
		C.FileModeld48130_DestroyFileModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackFileModeld48130_DropMimeData
func callbackFileModeld48130_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *FileModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackFileModeld48130_Index
func callbackFileModeld48130_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewFileModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *FileModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.FileModeld48130_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackFileModeld48130_Sibling
func callbackFileModeld48130_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewFileModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *FileModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.FileModeld48130_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackFileModeld48130_Flags
func callbackFileModeld48130_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewFileModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *FileModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.FileModeld48130_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackFileModeld48130_InsertColumns
func callbackFileModeld48130_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *FileModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackFileModeld48130_InsertRows
func callbackFileModeld48130_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *FileModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackFileModeld48130_MoveColumns
func callbackFileModeld48130_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *FileModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackFileModeld48130_MoveRows
func callbackFileModeld48130_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *FileModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackFileModeld48130_RemoveColumns
func callbackFileModeld48130_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *FileModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackFileModeld48130_RemoveRows
func callbackFileModeld48130_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *FileModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackFileModeld48130_SetData
func callbackFileModeld48130_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *FileModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackFileModeld48130_SetHeaderData
func callbackFileModeld48130_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *FileModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackFileModeld48130_SetItemData
func callbackFileModeld48130_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewFileModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *FileModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackFileModeld48130_Submit
func callbackFileModeld48130_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *FileModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackFileModeld48130_ColumnsAboutToBeInserted
func callbackFileModeld48130_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackFileModeld48130_ColumnsAboutToBeMoved
func callbackFileModeld48130_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackFileModeld48130_ColumnsAboutToBeRemoved
func callbackFileModeld48130_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackFileModeld48130_ColumnsInserted
func callbackFileModeld48130_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackFileModeld48130_ColumnsMoved
func callbackFileModeld48130_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackFileModeld48130_ColumnsRemoved
func callbackFileModeld48130_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackFileModeld48130_DataChanged
func callbackFileModeld48130_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackFileModeld48130_FetchMore
func callbackFileModeld48130_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewFileModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *FileModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackFileModeld48130_HeaderDataChanged
func callbackFileModeld48130_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(std_core.Qt__Orientation, int, int))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackFileModeld48130_LayoutAboutToBeChanged
func callbackFileModeld48130_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackFileModeld48130_LayoutChanged
func callbackFileModeld48130_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackFileModeld48130_ModelAboutToBeReset
func callbackFileModeld48130_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackFileModeld48130_ModelReset
func callbackFileModeld48130_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackFileModeld48130_ResetInternalData
func callbackFileModeld48130_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewFileModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *FileModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.FileModeld48130_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackFileModeld48130_Revert
func callbackFileModeld48130_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewFileModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *FileModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.FileModeld48130_RevertDefault(ptr.Pointer())
	}
}

//export callbackFileModeld48130_RowsAboutToBeInserted
func callbackFileModeld48130_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackFileModeld48130_RowsAboutToBeMoved
func callbackFileModeld48130_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackFileModeld48130_RowsAboutToBeRemoved
func callbackFileModeld48130_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackFileModeld48130_RowsInserted
func callbackFileModeld48130_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackFileModeld48130_RowsMoved
func callbackFileModeld48130_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackFileModeld48130_RowsRemoved
func callbackFileModeld48130_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackFileModeld48130_Sort
func callbackFileModeld48130_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, std_core.Qt__SortOrder))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewFileModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *FileModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackFileModeld48130_RoleNames
func callbackFileModeld48130_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__roleNames_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewFileModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *FileModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.FileModeld48130_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackFileModeld48130_ItemData
func callbackFileModeld48130_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__itemData_newList())
			for k, v := range signal.(func(*std_core.QModelIndex) map[int]*std_core.QVariant)(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__itemData_newList())
		for k, v := range NewFileModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *FileModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.FileModeld48130_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackFileModeld48130_MimeData
func callbackFileModeld48130_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData(signal.(func([]*std_core.QModelIndex) *std_core.QMimeData)(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewFileModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewFileModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *FileModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.FileModeld48130_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackFileModeld48130_Buddy
func callbackFileModeld48130_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewFileModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *FileModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.FileModeld48130_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackFileModeld48130_Parent
func callbackFileModeld48130_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewFileModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *FileModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.FileModeld48130_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackFileModeld48130_Match
func callbackFileModeld48130_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__match_newList())
			for _, v := range signal.(func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewFileModelFromPointer(NewFileModelFromPointer(nil).__match_newList())
		for _, v := range NewFileModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *FileModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewFileModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.FileModeld48130_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackFileModeld48130_Span
func callbackFileModeld48130_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize(signal.(func(*std_core.QModelIndex) *std_core.QSize)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewFileModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *FileModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.FileModeld48130_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackFileModeld48130_MimeTypes
func callbackFileModeld48130_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewFileModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *FileModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.FileModeld48130_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackFileModeld48130_Data
func callbackFileModeld48130_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(*std_core.QModelIndex, int) *std_core.QVariant)(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewFileModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *FileModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.FileModeld48130_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackFileModeld48130_HeaderData
func callbackFileModeld48130_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewFileModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *FileModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.FileModeld48130_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackFileModeld48130_SupportedDragActions
func callbackFileModeld48130_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewFileModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *FileModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.FileModeld48130_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackFileModeld48130_SupportedDropActions
func callbackFileModeld48130_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewFileModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *FileModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.FileModeld48130_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackFileModeld48130_CanDropMimeData
func callbackFileModeld48130_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *FileModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackFileModeld48130_CanFetchMore
func callbackFileModeld48130_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *FileModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackFileModeld48130_HasChildren
func callbackFileModeld48130_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *FileModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackFileModeld48130_ColumnCount
func callbackFileModeld48130_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewFileModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *FileModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.FileModeld48130_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackFileModeld48130_RowCount
func callbackFileModeld48130_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewFileModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *FileModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.FileModeld48130_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackFileModeld48130_Event
func callbackFileModeld48130_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *FileModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackFileModeld48130_EventFilter
func callbackFileModeld48130_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFileModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *FileModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FileModeld48130_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackFileModeld48130_ChildEvent
func callbackFileModeld48130_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewFileModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *FileModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackFileModeld48130_ConnectNotify
func callbackFileModeld48130_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewFileModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *FileModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackFileModeld48130_CustomEvent
func callbackFileModeld48130_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewFileModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *FileModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackFileModeld48130_DeleteLater
func callbackFileModeld48130_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewFileModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *FileModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.FileModeld48130_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackFileModeld48130_Destroyed
func callbackFileModeld48130_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackFileModeld48130_DisconnectNotify
func callbackFileModeld48130_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewFileModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *FileModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackFileModeld48130_ObjectNameChanged
func callbackFileModeld48130_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackFileModeld48130_TimerEvent
func callbackFileModeld48130_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewFileModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *FileModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.FileModeld48130_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
