package config

import (
	"github.com/syncfuture/keepa/protoc/keepaconfig"
)

// type KeepaConfig struct {
// 	ID        string
// 	UserID    string
// 	BaseURL   string
// 	AccessKey string
// 	Domain    int
// }

type IConfigGetter interface {
	Get(userID, configID string) (*keepaconfig.KeepaConfig, error)
}

type IConfigManager interface {
	Create(config *keepaconfig.KeepaConfig) error
	Update(config *keepaconfig.KeepaConfig) error
	Delete(userID, configID string) error
}
