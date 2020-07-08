package config

import (
	"github.com/syncfuture/advertise/protoc/adconfig"
)

// type KeepaConfig struct {
// 	ID        string
// 	UserID    string
// 	BaseURL   string
// 	AccessKey string
// 	Domain    int
// }

type IConfigGetter interface {
	Get(userID, configID string) (*adconfig.ADConfig, error)
}

type IConfigManager interface {
	Create(config *adconfig.ADConfig) error
	Update(config *adconfig.ADConfig) error
	Delete(userID, configID string) error
}
