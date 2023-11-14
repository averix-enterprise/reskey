package backend

import (
	"context"
)

type App struct {
	ctx context.Context
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetAllHotKeys() []HotKey {
	hotKeys := make([]HotKey, 0, len(RegisteredHotKeys))
	for _, val := range RegisteredHotKeys {
		hotKeys = append(hotKeys, val)
	}
	return hotKeys
}

func (a *App) AddHotKey() bool {
	return RegisterHotKey("None", make([]string, 0), 1920, 1080, true) != nil
}

func (a *App) UpdateHotKey(id string, key string, modifiers []string) bool {
	return ChangeKeys(id, key, modifiers)
}

func (a *App) DeleteHotKey(id string) bool {
	return UnregisterHotKey(id)
}
