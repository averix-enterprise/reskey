package backend

import (
	"fmt"
	"golang.design/x/hotkey"
)

var RegisteredKeys = make(map[string]Key)
var RegisteredHotKeys = make(map[string]HotKey)

type Key struct {
	Id         string `json:"id"`
	WindowsKey hotkey.Key
}

type HotKey struct {
	Id             string         `json:"id"`
	Width          uint32         `json:"width"`
	Height         uint32         `json:"height"`
	KeyInstance    Key            `json:"-"`
	HotKeyInstance *hotkey.Hotkey `json:"-"`
	Running        bool           `json:"-"`
}

func registerKey(id string, windowsKey hotkey.Key) {
	RegisteredKeys[id] = Key{Id: id, WindowsKey: windowsKey}
}

func Initialize() {
	registerKey("P", hotkey.KeyP)
	registerKey("U", hotkey.KeyU)

	loadHotKeysFromFile()
}

func loadHotKeysFromFile() {
	var data map[string]HotKey
	LoadOrCreateFile("hotkeys.json", &data, make(map[string]HotKey))
	for _, value := range data {
		hk := value
		RegisterHotKey(hk.Id, hk.Width, hk.Height, false)
	}
}

func saveHotKeysIntoFile() {
	WriteToFile("hotkeys.json", RegisteredHotKeys)
}

func RegisterHotKey(keyId string, width uint32, height uint32, save bool) {
	key, ok := RegisteredKeys[keyId]
	if !ok {
		panic("Cannot find key: " + keyId)
	}
	hotKey := HotKey{
		Id:             key.Id,
		Width:          width,
		Height:         height,
		KeyInstance:    key,
		Running:        true,
		HotKeyInstance: hotkey.New([]hotkey.Modifier{}, key.WindowsKey),
	}
	RegisteredHotKeys[keyId] = hotKey

	err := hotKey.HotKeyInstance.Register()
	if err != nil {
		return
	}
	if ok {
		go run(&hotKey)
	}
	if save {
		saveHotKeysIntoFile()
	}
}

func run(hk *HotKey) {
	for hk.Running {
		<-hk.HotKeyInstance.Keydown()
		fmt.Printf("Exec: %s, Running: %t\n", hk.Id, hk.Running)
	}
}

func UnregisterHotKey(key string) {
	hotKey, ok := RegisteredHotKeys[key]
	if !ok {
		return
	}
	hotKey.Running = false
	err := hotKey.HotKeyInstance.Unregister()
	if err != nil {
		return
	}
	saveHotKeysIntoFile()
}
