// Code generated by "./generator ./session/org.desktopspec.applicationmanager1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package applicationmanager1

import (
	"errors"
	"unsafe"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Application interface {
	application // interface org.desktopspec.ApplicationManager1.Application
	proxy.Object
}

type objectApplication struct {
	interfaceApplication // interface org.desktopspec.ApplicationManager1.Application
	proxy.ImplObject
}

func NewApplication(conn *dbus.Conn, path dbus.ObjectPath) (Application, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectApplication)
	obj.ImplObject.Init_(conn, "org.desktopspec.ApplicationManager1", path)
	return obj, nil
}

type application interface {
	GoLaunch(flags dbus.Flags, ch chan *dbus.Call, action string, fields []string, options map[string]dbus.Variant) *dbus.Call
	Launch(flags dbus.Flags, action string, fields []string, options map[string]dbus.Variant) (dbus.ObjectPath, error)
	GoSendToDesktop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	SendToDesktop(flags dbus.Flags) (bool, error)
	GoRemoveFromDesktop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RemoveFromDesktop(flags dbus.Flags) (bool, error)
	Categories() proxy.PropStringArray
	X_linglong() proxy.PropBool
	X_Flatpak() proxy.PropBool
	InstalledTime() proxy.PropUint64
	NoDisplay() proxy.PropBool
	MimeTypes() proxy.PropStringArray
	Actions() proxy.PropStringArray
	AutoStart() proxy.PropBool
	LastLaunchedTime() proxy.PropUint64
	Instances() proxy.PropObjectPathArray
	ID() proxy.PropString
	Terminal() proxy.PropBool
	ScaleFactor() proxy.PropDouble
	IsOnDesktop() proxy.PropBool
}

type interfaceApplication struct{}

func (v *interfaceApplication) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceApplication) GetInterfaceName_() string {
	return "org.desktopspec.ApplicationManager1.Application"
}

// method Launch

func (v *interfaceApplication) GoLaunch(flags dbus.Flags, ch chan *dbus.Call, action string, fields []string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Launch", flags, ch, action, fields, options)
}

func (*interfaceApplication) StoreLaunch(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *interfaceApplication) Launch(flags dbus.Flags, action string, fields []string, options map[string]dbus.Variant) (dbus.ObjectPath, error) {
	return v.StoreLaunch(
		<-v.GoLaunch(flags, make(chan *dbus.Call, 1), action, fields, options).Done)
}

// method SendToDesktop

func (v *interfaceApplication) GoSendToDesktop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SendToDesktop", flags, ch)
}

func (*interfaceApplication) StoreSendToDesktop(call *dbus.Call) (success bool, err error) {
	err = call.Store(&success)
	return
}

func (v *interfaceApplication) SendToDesktop(flags dbus.Flags) (bool, error) {
	return v.StoreSendToDesktop(
		<-v.GoSendToDesktop(flags, make(chan *dbus.Call, 1)).Done)
}

// method RemoveFromDesktop

func (v *interfaceApplication) GoRemoveFromDesktop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveFromDesktop", flags, ch)
}

func (*interfaceApplication) StoreRemoveFromDesktop(call *dbus.Call) (success bool, err error) {
	err = call.Store(&success)
	return
}

func (v *interfaceApplication) RemoveFromDesktop(flags dbus.Flags) (bool, error) {
	return v.StoreRemoveFromDesktop(
		<-v.GoRemoveFromDesktop(flags, make(chan *dbus.Call, 1)).Done)
}

// property Categories as

func (v *interfaceApplication) Categories() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "Categories",
	}
}

// property X_linglong b

func (v *interfaceApplication) X_linglong() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "X_linglong",
	}
}

// property X_Flatpak b

func (v *interfaceApplication) X_Flatpak() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "X_Flatpak",
	}
}

// property installedTime t

func (v *interfaceApplication) InstalledTime() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "installedTime",
	}
}

// property NoDisplay b

func (v *interfaceApplication) NoDisplay() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "NoDisplay",
	}
}

// property MimeTypes as

func (v *interfaceApplication) MimeTypes() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "MimeTypes",
	}
}

// property Actions as

func (v *interfaceApplication) Actions() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "Actions",
	}
}

// property AutoStart b

func (v *interfaceApplication) AutoStart() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "AutoStart",
	}
}

// property LastLaunchedTime t

func (v *interfaceApplication) LastLaunchedTime() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "LastLaunchedTime",
	}
}

// property Instances ao

func (v *interfaceApplication) Instances() proxy.PropObjectPathArray {
	return &proxy.ImplPropObjectPathArray{
		Impl: v,
		Name: "Instances",
	}
}

// property ID s

func (v *interfaceApplication) ID() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "ID",
	}
}

// property Terminal b

func (v *interfaceApplication) Terminal() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Terminal",
	}
}

// property ScaleFactor d

func (v *interfaceApplication) ScaleFactor() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "ScaleFactor",
	}
}

// property isOnDesktop b

func (v *interfaceApplication) IsOnDesktop() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "isOnDesktop",
	}
}

type Manager interface {
	manager // interface org.desktopspec.ApplicationManager1
	proxy.Object
}

type objectManager struct {
	interfaceManager // interface org.desktopspec.ApplicationManager1
	proxy.ImplObject
}

func NewManager(conn *dbus.Conn) Manager {
	obj := new(objectManager)
	obj.ImplObject.Init_(conn, "org.desktopspec.ApplicationManager1", "/org/desktopspec/ApplicationManager1")
	return obj
}

type manager interface {
	GoReloadApplications(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ReloadApplications(flags dbus.Flags) error
	GoIdentify(flags dbus.Flags, ch chan *dbus.Call, pidfd dbus.UnixFD) *dbus.Call
	Identify(flags dbus.Flags, pidfd dbus.UnixFD) (string, dbus.ObjectPath, map[string]map[string]dbus.Variant, error)
	GoAddUserApplication(flags dbus.Flags, ch chan *dbus.Call, desktop_file map[string]dbus.Variant, name string) *dbus.Call
	AddUserApplication(flags dbus.Flags, desktop_file map[string]dbus.Variant, name string) (string, error)
	List() proxy.PropObjectPathArray
}

type interfaceManager struct{}

func (v *interfaceManager) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceManager) GetInterfaceName_() string {
	return "org.desktopspec.ApplicationManager1"
}

// method ReloadApplications

func (v *interfaceManager) GoReloadApplications(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReloadApplications", flags, ch)
}

func (v *interfaceManager) ReloadApplications(flags dbus.Flags) error {
	return (<-v.GoReloadApplications(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Identify

func (v *interfaceManager) GoIdentify(flags dbus.Flags, ch chan *dbus.Call, pidfd dbus.UnixFD) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Identify", flags, ch, pidfd)
}

func (*interfaceManager) StoreIdentify(call *dbus.Call) (id string, instance dbus.ObjectPath, application_instance_info map[string]map[string]dbus.Variant, err error) {
	err = call.Store(&id, &instance, &application_instance_info)
	return
}

func (v *interfaceManager) Identify(flags dbus.Flags, pidfd dbus.UnixFD) (string, dbus.ObjectPath, map[string]map[string]dbus.Variant, error) {
	return v.StoreIdentify(
		<-v.GoIdentify(flags, make(chan *dbus.Call, 1), pidfd).Done)
}

// method addUserApplication

func (v *interfaceManager) GoAddUserApplication(flags dbus.Flags, ch chan *dbus.Call, desktop_file map[string]dbus.Variant, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".addUserApplication", flags, ch, desktop_file, name)
}

func (*interfaceManager) StoreAddUserApplication(call *dbus.Call) (app_id string, err error) {
	err = call.Store(&app_id)
	return
}

func (v *interfaceManager) AddUserApplication(flags dbus.Flags, desktop_file map[string]dbus.Variant, name string) (string, error) {
	return v.StoreAddUserApplication(
		<-v.GoAddUserApplication(flags, make(chan *dbus.Call, 1), desktop_file, name).Done)
}

// property List ao

func (v *interfaceManager) List() proxy.PropObjectPathArray {
	return &proxy.ImplPropObjectPathArray{
		Impl: v,
		Name: "List",
	}
}

type MimeManager interface {
	mimeManager // interface org.desktopspec.MimeManager1
	proxy.Object
}

type objectMimeManager struct {
	interfaceMimeManager // interface org.desktopspec.MimeManager1
	proxy.ImplObject
}

func NewMimeManager(conn *dbus.Conn) MimeManager {
	obj := new(objectMimeManager)
	obj.ImplObject.Init_(conn, "org.desktopspec.ApplicationManager1", "/org/desktopspec/ApplicationManager1/MimeManager1")
	return obj
}

type mimeManager interface {
	GoQueryDefaultApplication(flags dbus.Flags, ch chan *dbus.Call, content string) *dbus.Call
	QueryDefaultApplication(flags dbus.Flags, content string) (string, dbus.ObjectPath, error)
	GoSetDefaultApplication(flags dbus.Flags, ch chan *dbus.Call, defaultApps map[string]string) *dbus.Call
	SetDefaultApplication(flags dbus.Flags, defaultApps map[string]string) error
	GoUnsetDefaultApplication(flags dbus.Flags, ch chan *dbus.Call, mimeTypes []string) *dbus.Call
	UnsetDefaultApplication(flags dbus.Flags, mimeTypes []string) error
	GoListApplications(flags dbus.Flags, ch chan *dbus.Call, mimeType string) *dbus.Call
	ListApplications(flags dbus.Flags, mimeType string) (map[dbus.ObjectPath]map[string]map[string]dbus.Variant, error)
}

type interfaceMimeManager struct{}

func (v *interfaceMimeManager) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceMimeManager) GetInterfaceName_() string {
	return "org.desktopspec.MimeManager1"
}

// method queryDefaultApplication

func (v *interfaceMimeManager) GoQueryDefaultApplication(flags dbus.Flags, ch chan *dbus.Call, content string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".queryDefaultApplication", flags, ch, content)
}

func (*interfaceMimeManager) StoreQueryDefaultApplication(call *dbus.Call) (mimeType string, application dbus.ObjectPath, err error) {
	err = call.Store(&mimeType, &application)
	return
}

func (v *interfaceMimeManager) QueryDefaultApplication(flags dbus.Flags, content string) (string, dbus.ObjectPath, error) {
	return v.StoreQueryDefaultApplication(
		<-v.GoQueryDefaultApplication(flags, make(chan *dbus.Call, 1), content).Done)
}

// method setDefaultApplication

func (v *interfaceMimeManager) GoSetDefaultApplication(flags dbus.Flags, ch chan *dbus.Call, defaultApps map[string]string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".setDefaultApplication", flags, ch, defaultApps)
}

func (v *interfaceMimeManager) SetDefaultApplication(flags dbus.Flags, defaultApps map[string]string) error {
	return (<-v.GoSetDefaultApplication(flags, make(chan *dbus.Call, 1), defaultApps).Done).Err
}

// method unsetDefaultApplication

func (v *interfaceMimeManager) GoUnsetDefaultApplication(flags dbus.Flags, ch chan *dbus.Call, mimeTypes []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".unsetDefaultApplication", flags, ch, mimeTypes)
}

func (v *interfaceMimeManager) UnsetDefaultApplication(flags dbus.Flags, mimeTypes []string) error {
	return (<-v.GoUnsetDefaultApplication(flags, make(chan *dbus.Call, 1), mimeTypes).Done).Err
}

// method listApplications

func (v *interfaceMimeManager) GoListApplications(flags dbus.Flags, ch chan *dbus.Call, mimeType string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".listApplications", flags, ch, mimeType)
}

func (*interfaceMimeManager) StoreListApplications(call *dbus.Call) (applications_and_properties map[dbus.ObjectPath]map[string]map[string]dbus.Variant, err error) {
	err = call.Store(&applications_and_properties)
	return
}

func (v *interfaceMimeManager) ListApplications(flags dbus.Flags, mimeType string) (map[dbus.ObjectPath]map[string]map[string]dbus.Variant, error) {
	return v.StoreListApplications(
		<-v.GoListApplications(flags, make(chan *dbus.Call, 1), mimeType).Done)
}
