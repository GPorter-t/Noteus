package initialize

import (
	"Noteus/global"
	"Noteus/storage"
)

func Storage() {
	systemCfg := global.GVA_CONFIG.System
	global.GVA_STORE = &storage.Storage{
		LRU: storage.NewLRUStore(systemCfg.LruMaxSize),
	}
}
