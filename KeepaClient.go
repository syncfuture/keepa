package keepa

type KeepaClient struct {
	Options *KeepaOptions
}

func NewClient(options *KeepaOptions) (r *KeepaClient) {
	r = new(KeepaClient)
	r.Options = options
	return r
}
