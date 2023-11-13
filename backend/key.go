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
	Id             string `json:"id"`
	Width          uint32 `json:"width"`
	Height         uint32 `json:"height"`
	KeyInstance    Key
	Channel        chan bool
	HotKeyInstance *hotkey.Hotkey
}

func registerKey(id string, windowsKey hotkey.Key) {
	RegisteredKeys[id] = Key{Id: id, WindowsKey: windowsKey}
}

func Initialize() {
	registerKey("P", hotkey.KeyP)
	registerKey("U", hotkey.KeyU)
	registerHotKey("P", 1920, 1080)
	registerHotKey("P", 1920, 1080)
}

func registerHotKey(keyId string, width uint32, height uint32) {
	key, ok := RegisteredKeys[keyId]
	if !ok {
		panic("Cannot find key: " + keyId)
	}
	hotKey := HotKey{
		Id:             key.Id,
		Width:          width,
		Height:         height,
		KeyInstance:    key,
		Channel:        make(chan bool),
		HotKeyInstance: hotkey.New([]hotkey.Modifier{}, key.WindowsKey),
	}
	RegisteredHotKeys[keyId] = hotKey

	err := hotKey.HotKeyInstance.Register()
	if err != nil {
		return
	}
	if ok {
		go func() {
			for {
				select {
				case <-hotKey.Channel:
					return
				default:
					<-hotKey.HotKeyInstance.Keydown()
					fmt.Println("Working")
				}
			}
		}()
	}
}

func unregisterHotKey(key string) {
	hotKey, ok := RegisteredHotKeys[key]
	if !ok {
		panic("Cannot find hotkey: " + key)
	}
	hotKey.Channel <- true
	err := hotKey.HotKeyInstance.Unregister()
	if err != nil {
		panic("Cannot unregister hotkey: " + key)
	}

}
