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
	Id          string `json:"id"`
	Width       uint32 `json:"width"`
	Height      uint32 `json:"height"`
	KeyInstance Key
}

func registerKey(id string, windowsKey hotkey.Key) {
	RegisteredKeys[id] = Key{Id: id, WindowsKey: windowsKey}
}

func Initialize() {
	registerKey("P", hotkey.KeyP)
	registerHotKey("P", 1920, 1080)
}

func registerHotKey(keyId string, width uint32, height uint32) {
	key, ok := RegisteredKeys[keyId]
	if !ok {
		panic("Cannot find key: " + keyId)
	}
	hotKey := HotKey{
		Id:          key.Id,
		Width:       width,
		Height:      height,
		KeyInstance: key,
	}
	RegisteredHotKeys[keyId] = hotKey
	if ok {
		//Check os
		go run(&hotKey)
	}
}

func unregisterHotKey(key string) {

}

func run(key *HotKey) {
	hk := hotkey.New([]hotkey.Modifier{}, key.KeyInstance.WindowsKey)
	err := hk.Register()
	if err != nil {
		return
	}
	for true {
		<-hk.Keydown()
		fmt.Println("Penis")
	}
	err = hk.Unregister()
	panic("Could not unregister hotkey: " + key.Id)
}
