package backend

import (
	"fmt"
	"github.com/google/uuid"
	"golang.design/x/hotkey"
	"reflect"
)

var RegisteredKeys = make(map[string]Key)
var RegisteredMods = make(map[string]Mod)
var RegisteredHotKeys = make(map[string]HotKey)

type Key struct {
	Id         string `json:"id"`
	WindowsKey hotkey.Key
}

type Mod struct {
	Id         string `json:"id"`
	WindowsKey hotkey.Modifier
}

type HotKey struct {
	Id             string         `json:"id"`
	Key            string         `json:"key"`
	Modifiers      []string       `json:"modifiers"`
	Width          uint32         `json:"width"`
	Height         uint32         `json:"height"`
	KeyInstance    Key            `json:"-"`
	HotKeyInstance *hotkey.Hotkey `json:"-"`
	Running        bool           `json:"-"`
}

func registerKey(id string, windowsKey hotkey.Key) {
	RegisteredKeys[id] = Key{Id: id, WindowsKey: windowsKey}
}

func registerModifier(id string, windowsKey hotkey.Modifier) {
	RegisteredMods[id] = Mod{Id: id, WindowsKey: windowsKey}
}

func Initialize() {
	registerKey("P", hotkey.KeyP)
	registerKey("U", hotkey.KeyU)
	registerModifier("Shift", hotkey.ModShift)

	loadHotKeysFromFile()
}

func loadHotKeysFromFile() {
	var data map[string]HotKey
	LoadOrCreateFile("hotkeys.json", &data, make(map[string]HotKey))
	for _, value := range data {
		hk := value
		newHotKey := RegisterHotKey(hk.Key, hk.Modifiers, hk.Width, hk.Height, false)
		newHotKey.Id = hk.Id
		StartHotKey(newHotKey)
	}
}

func saveHotKeysIntoFile() {
	WriteToFile("hotkeys.json", RegisteredHotKeys)
}

func StartHotKey(hk *HotKey) bool {
	if hk.Key == "None" {
		return true
	}
	hk.Running = true

	var modifier []hotkey.Modifier
	if len(hk.Modifiers) > 0 {
		for _, value := range hk.Modifiers {
			modifier = append(modifier, RegisteredMods[value].WindowsKey)
		}

	}
	hk.HotKeyInstance = hotkey.New(modifier, RegisteredKeys[hk.Key].WindowsKey)
	err := hk.HotKeyInstance.Register()
	if err != nil {
		fmt.Println(err)
		return false
	}
	go run(hk)
	return true
}

func IsKeyAvailable(key string, modifiers []string) bool {
	if key == "None" {
		return true
	}
	_, ok := RegisteredKeys[key]
	if !ok {
		return false
	}
	if len(modifiers) > 0 {
		for _, value := range modifiers {
			_, ok = RegisteredMods[value]
			if !ok {
				return false
			}
		}
	}

	for _, value := range RegisteredHotKeys {
		if value.Key != key {
			continue
		}
		if reflect.DeepEqual(value.Modifiers, modifiers) {
			return false
		}
	}
	return true
}

func RegisterHotKey(key string, modifiers []string, width uint32, height uint32, save bool) *HotKey {
	if !IsKeyAvailable(key, modifiers) {
		return nil
	}
	hotKey := HotKey{
		Id:        uuid.New().String(),
		Width:     width,
		Height:    height,
		Modifiers: modifiers,
		Key:       key,
	}
	RegisteredHotKeys[hotKey.Id] = hotKey
	if save {
		saveHotKeysIntoFile()
	}
	return &hotKey
}

func run(hk *HotKey) {
	for hk.Running {
		<-hk.HotKeyInstance.Keydown()
		fmt.Printf("Exec: %s, Running: %t\n", hk.Id, hk.Running)
	}
}

func ChangeKeys(id string, key string, modifiers []string) bool {
	hk, ok := RegisteredHotKeys[id]
	if !ok {
		return false
	}
	StopHotKey(&hk)
	hk.Key = key
	hk.Modifiers = modifiers
	return StartHotKey(&hk)
}

func UnregisterHotKey(id string) bool {
	hotKey, ok := RegisteredHotKeys[id]
	if !ok {
		return false
	}
	StopHotKey(&hotKey)
	delete(RegisteredHotKeys, id)
	return true
}

func StopHotKey(hotKey *HotKey) bool {
	if hotKey.Key == "None" {
		return true
	}
	hotKey.Running = false
	err := hotKey.HotKeyInstance.Unregister()
	if err != nil {
		return false
	}
	return true
}
