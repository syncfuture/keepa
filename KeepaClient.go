package keepa

import (
	"github.com/syncfuture/keepa/protoc/keepaconfig"
)

type KeepaClient struct {
	Config *keepaconfig.KeepaConfig
}

func NewClient(config *keepaconfig.KeepaConfig) (r *KeepaClient) {
	r = new(KeepaClient)
	r.Config = config
	return r
}
