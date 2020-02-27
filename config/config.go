package config

type KeepaConfig struct {
	ID        string
	UserID    string
	BaseURL   string
	AccessKey string
	Domain    int
}

type IConfigGetter interface {
	Get(userID, configID string) (*KeepaConfig, error)
}

type IConfigManager interface {
	Create(config *KeepaConfig) error
	Update(config *KeepaConfig) error
	Delete(userID, configID string) error
}
