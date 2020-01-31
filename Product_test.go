package keepa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProduct(t *testing.T) {
	client := NewClient(&KeepaOptions{
		BaseURL:   "https://api.keepa.com",
		Domain:    1,
		AccessKey: "148rdq9hof8ip4gkalbn6tcato69c06nnvqachkdqldenjfrtp8kkanhb1id6sk0",
	})

	r, err := client.GetProduct(&GetProductQuery{
		ASIN: "B076TH98FX",
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, r.Products)
}
