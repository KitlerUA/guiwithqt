package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	app.SetApplicationDisplayName("Quatrix Express")

	var view = quick.NewQQuickView(nil)

	initFiles(view)
	initHostname(view)
	initLogin(view)
	initHelper(view)

	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)

	view.Show()

	gui.QGuiApplication_Exec()
}

type Hostname struct {
	core.QObject

	//notice: used to receive signals
	_ func(hostname string) bool `slot:"find"`
}

//notice: this function can be called before or after SetSource
func initHostname(view *quick.QQuickView) {

	//notice: New+Objectname
	hostname := NewHostname(nil)

	view.RootContext().SetContextProperty("goHostname", hostname)

	//notice: Connect+Slotname
	hostname.ConnectFind(func(hostname string) bool {
		r, err := http.NewRequest("GET", fmt.Sprintf("https://%s.quatrix.it/api/1.0/account/metadata", hostname), nil)
		if err != nil {
			log.Printf("failed to create new request: %s", err)
			return false
		}
		c := http.Client{}
		response, err := c.Do(r)
		if err != nil {
			log.Printf("failed to do request: %s", err)
			return false
		}

		if response.StatusCode != http.StatusOK {
			log.Printf("hostname(%s) not found: %s", hostname, response.Status)
			return false
		}

		log.Printf("hostname(%s) found", hostname)

		//todo: we return to this later
		state.SetHostname(hostname)
		log.Printf("now hostname = %s", state.SetHostname)

		return true
	})
}

type State struct {
	core.QObject

	_ string       `property:"hostname"`
	_ string       `property:"email"`
	_ *http.Cookie `property:"cookie"`

	//notice: used to emit signal
	_ func(hostname, email string) `signal:"change"`
}

//! even if you access state's fields in go yoy have to use New func
var state = NewState(nil)

type Login struct {
	core.QObject

	_ func(username, password string) bool `slot:"login"`
}

func initLogin(view *quick.QQuickView) {
	login := NewLogin(nil)
	view.RootContext().SetContextProperty("goLogin", login)

	//notice: one more thing: you din't connect signal to function, because it's not a function
	login.ConnectLogin(func(username, password string) bool {
		r, err := http.NewRequest("GET", fmt.Sprintf("https://%s.quatrix.it/api/1.0/session/login", state.Hostname()), nil)
		if err != nil {
			log.Printf("failed to create new request: %s", err)
			return false
		}

		r.SetBasicAuth(username, password)

		c := http.Client{}
		response, err := c.Do(r)
		if err != nil {
			log.Printf("failed to do request: %s", err)
			return false
		}

		if response.StatusCode != http.StatusOK {
			log.Printf("login(%s) not found: %s", username, response.Status)
			return false
		}

		if len(response.Cookies()) > 0 {
			state.SetCookie(response.Cookies()[0])
		}

		log.Printf("login(%s) found", username)

		return true
	})
}

const (
	FileName = int(core.Qt__UserRole) + 1<<iota
	MTime
)

var files *FileModel

type FileModel struct {
	core.QAbstractListModel

	//notice: say heil to the (king) CONSTRUCTOR
	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*QuatrixFile           `property:"files"`
}

type QuatrixFile struct {
	core.QObject

	_ string `property:"fileName"`
	_ string `property:"mTime"`
}

func init() {
	//! to use QuatrixFile as base type for List, you should register it before FileModel
	QuatrixFile_QRegisterMetaType()
}

func (m *FileModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		FileName: core.NewQByteArray2("fileName", len("fileName")),
		MTime:    core.NewQByteArray2("mTime", len("mTime")),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)
}

func (m *FileModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Files()) {
		return core.NewQVariant()
	}

	var p = m.Files()[index.Row()]

	switch role {
	case FileName:
		{
			return core.NewQVariant14(p.FileName())
		}

	case MTime:
		{
			return core.NewQVariant14(p.MTime())
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *FileModel) rowCount(parent *core.QModelIndex) int {
	log.Printf("len files: %d", len(m.Files()))
	return len(m.Files())
}

func (m *FileModel) columnCount(parent *core.QModelIndex) int {
	return 2
}

func (m *FileModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

//! you have to create Files before SetSource, because it's not simple object, but list-model
func initFiles(view *quick.QQuickView) {
	files = NewFileModel(nil)
	view.RootContext().SetContextProperty("FilesModel", files)
}

type FileHelper struct {
	core.QObject

	_ func() `slot:"metadata"`
	_ func() `signal:"changed"`
}

func initHelper(view *quick.QQuickView) {
	helper := NewFileHelper(nil)
	view.RootContext().SetContextProperty("goHelper", helper)

	helper.ConnectMetadata(func() {
		r, err := http.NewRequest("GET", fmt.Sprintf("https://%s.quatrix.it/api/1.0/file/metadata", state.Hostname()), nil)
		if err != nil {
			log.Printf("failed to create new request: %s", err)
			return
		}

		r.AddCookie(state.Cookie())

		c := http.Client{}
		response, err := c.Do(r)
		if err != nil {
			log.Printf("failed to do request: %s", err)
			return
		}

		if response.StatusCode != http.StatusOK {
			log.Printf("failed to get metadata: %s", response.Status)
			return
		}

		if len(response.Cookies()) > 0 {
			state.SetCookie(response.Cookies()[0])
		}

		var metadata Metadata

		rawBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("read body failed: %s", err)
			return
		}
		defer response.Body.Close()

		err = json.Unmarshal(rawBody, &metadata)
		if err != nil {
			log.Printf("unmarshal failed: %s", err)
			return
		}

		log.Printf("raw metadata: %s", rawBody)

		log.Printf("metadata: %+v", metadata)

		var qFiles []*QuatrixFile
		for _, f := range metadata.Content {
			qFile := NewQuatrixFile(nil)
			qFile.SetFileName(f.Name)
			qFile.SetMTime(time.Unix(int64(f.Created), 0).Format("2006-01-02"))
			qFiles = append(qFiles, qFile)
		}

		files.SetFiles(qFiles)

		//notice: just emit signal
		helper.Changed()

		log.Printf("files set: %+v", qFiles)
		return

	})
}

type Metadata struct {
	Operations int     `json:"operations"`
	ModifiedMs float64 `json:"modified_ms"`
	Name       string  `json:"name"`
	Created    int     `json:"created"`
	Modified   int     `json:"modified"`
	SubType    string  `json:"sub_type"`
	Content    []struct {
		Operations int     `json:"operations"`
		ModifiedMs float64 `json:"modified_ms"`
		Name       string  `json:"name"`
		Created    int     `json:"created"`
		Modified   int     `json:"modified"`
		SubType    string  `json:"sub_type"`
		ParentID   string  `json:"parent_id"`
		Gid        int     `json:"gid"`
		Size       int     `json:"size"`
		Type       string  `json:"type"`
		ID         string  `json:"id"`
		UID        int     `json:"uid"`
	} `json:"content"`
	ParentID string `json:"parent_id"`
	Gid      int    `json:"gid"`
	Size     int    `json:"size"`
	Type     string `json:"type"`
	ID       string `json:"id"`
	UID      int    `json:"uid"`
}
