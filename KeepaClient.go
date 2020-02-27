package keepa

import "github.com/syncfuture/keepa/config"

type KeepaClient struct {
	Config *config.KeepaConfig
}

func NewClient(config *config.KeepaConfig) (r *KeepaClient) {
	r = new(KeepaClient)
	r.Config = config
	return r
}
