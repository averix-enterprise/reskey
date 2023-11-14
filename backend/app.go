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
