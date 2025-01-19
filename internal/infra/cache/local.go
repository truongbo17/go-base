package cache

import (
	"github.com/chenyahui/gin-cache/persist"
	"time"
)

func InitCacheLocal() *persist.MemoryStore {
	return persist.NewMemoryStore(1 * time.Minute)
}
