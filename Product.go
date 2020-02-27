package keepa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	u "github.com/syncfuture/go/util"
)

type GetProductQuery struct {
	ASIN    string
	Code    string
	Update  string
	History string
	Stats   string
	Offers  string
	Buybox  string
	Rating  string
	FBAfees string
	Rental  string
}

type ProductResults struct {
	Timestamp          int64         `json:"timestamp"`
	TokensLeft         int           `json:"tokensLeft"`
	RefillIn           int           `json:"refillIn"`
	RefillRate         int           `json:"refillRate"`
	TokenFlowReduction float32       `json:"tokenFlowReduction"`
	TokensConsumed     int           `json:"tokensConsumed"`
	ProcessingTimeInMs int           `json:"processingTimeInMs"`
	Products           []*ProductDTO `json:"products"`
	ItemHeight         int           `json:"itemHeight"`
	ItemLength         int           `json:"itemLength"`
	ItemWidth          int           `json:"itemWidth"`
	ItemWeight         int           `json:"itemWeight"`
	Stats              string        `json:"stats"`
	OffersSuccessful   bool          `json:"offersSuccessful"`
	G                  int           `json:"g"`
}

type ProductDTO struct {
	// CSV                             [][]int `json:"csv,omitempty"`
	Categories                      []int         `json:"categories,omitempty"`
	ImagesCSV                       string        `json:"imagesCSV,omitempty"`
	Manufacturer                    string        `json:"manufacturer,omitempty"`
	Title                           string        `json:"title,omitempty"`
	LastUpdate                      int           `json:"lastUpdate,omitempty"`
	LastPriceChange                 int           `json:"lastPriceChange,omitempty"`
	RootCategory                    int           `json:"rootCategory,omitempty"`
	ProductType                     int           `json:"productType,omitempty"`
	ParentAsin                      string        `json:"parentAsin,omitempty"`
	VariationCSV                    string        `json:"variationCSV,omitempty"`
	ASIN                            string        `json:"asin,omitempty"`
	DomainId                        int           `json:"domainId,omitempty"`
	Type                            string        `json:"type,omitempty"`
	HasReviews                      bool          `json:"hasReviews,omitempty"`
	MPN                             string        `json:"mpnstring,omitempty"`
	TrackingSince                   int           `json:"trackingSince,omitempty"`
	Brand                           string        `json:"brandstring,omitempty"`
	Label                           string        `json:"label,omitempty"`
	Department                      string        `json:"department,omitempty"`
	Publisher                       string        `json:"publisher,omitempty"`
	ProductGroup                    string        `json:"productGroupstring,omitempty"`
	PartNumber                      string        `json:"partNumber,omitempty"`
	Studio                          string        `json:"studio,omitempty"`
	Genre                           string        `json:"genre,omitempty"`
	Model                           string        `json:"model,omitempty"`
	Color                           string        `json:"color,omitempty"`
	Size                            string        `json:"size,omitempty"`
	Edition                         string        `json:"edition,omitempty"`
	Platform                        string        `json:"platform,omitempty"`
	Format                          string        `json:"format,omitempty"`
	PackageHeight                   int           `json:"packageHeight,omitempty"`
	PackageLength                   int           `json:"packageLength,omitempty"`
	PackageWidth                    int           `json:"packageWidth,omitempty"`
	PackageWeight                   int           `json:"packageWeight,omitempty"`
	PackageQuantity                 int           `json:"packageQuantityint,omitempty"`
	IsAdultProduct                  bool          `json:"isAdultProduct,omitempty"`
	IsEligibleForTradeIn            bool          `json:"isEligibleForTradeIn,omitempty"`
	IsEligibleForSuperSaverShipping bool          `json:"isEligibleForSuperSaverShipping,omitempty"`
	Offers                          []interface{} `json:"offers,omitempty"`
	BuyBoxSellerIdHistory           []string      `json:"buyBoxSellerIdHistory,omitempty"`
	IsRedirectASIN                  bool          `json:"isRedirectASIN,omitempty"`
	IsSNS                           bool          `json:"isSNS,omitempty"`
	Author                          string        `json:"author,omitempty"`
	Binding                         string        `json:"binding,omitempty"`
	NumberOfItems                   int           `json:"numberOfItems,omitempty"`
	NumberOfPages                   int           `json:"numberOfPages,omitempty"`
	PublicationDate                 int           `json:"publicationDate,omitempty"`
	ReleaseDate                     int           `json:"releaseDate,omitempty"`
	Languages                       [][]string    `json:"languages,omitempty"`
	LastRatingUpdate                int           `json:"lastRatingUpdate,omitempty"`
	EbayListingIds                  []float64     `json:"ebayListingIds,omitempty"`
	LastEbayUpdate                  int           `json:"lastEbayUpdate,omitempty"`
	EanList                         []string      `json:"eanList,omitempty"`
	UpcList                         []string      `json:"upcList,omitempty"`
	LiveOffersOrder                 []int         `json:"liveOffersOrder,omitempty"`
	FrequentlyBoughtTogether        []string      `json:"frequentlyBoughtTogether,omitempty"`
	Features                        []string      `json:"features,omitempty"`
	Description                     string        `json:"description,omitempty"`
	HazardousMaterialType           int           `json:"hazardousMaterialType,omitempty"`
	Promotions                      []interface{} `json:"promotions,omitempty"`
	NewPriceIsMAP                   bool          `json:"newPriceIsMAP,omitempty"`
	Coupon                          []int         `json:"coupon,omitempty"`
	AvailabilityAmazon              int           `json:"availabilityAmazon,omitempty"`
	ListedSince                     int           `json:"listedSince,omitempty"`
	FbaFees                         interface{}   `json:"fbaFees,omitempty"`
	Variations                      []interface{} `json:"variations,omitempty"`
}

func (x *KeepaClient) GetProduct(query *GetProductQuery) (r *ProductResults, err error) {
	r = new(ProductResults)
	var url string
	if query.ASIN != "" {
		url = fmt.Sprintf("%s/product?key=%s&domain=%d&asin=%s", x.Config.BaseURL, x.Config.AccessKey, x.Config.Domain, query.ASIN)
	} else if query.Code != "" {
		url = fmt.Sprintf("%s/product?key=%s&domain=%d&code=%s", x.Config.BaseURL, x.Config.AccessKey, x.Config.Domain, query.Code)
	} else {
		panic("ASIN and Code cannot both be empty")
	}

	if query.Update != "" {
		url += "&update=" + query.Update
	}
	if query.History != "" {
		url += "&history=" + query.History
	}
	if query.Stats != "" {
		url += "&stats=" + query.Stats
	}
	if query.Buybox != "" {
		url += "&buybox=" + query.Buybox
	}
	if query.Rating != "" {
		url += "&rating=" + query.Rating
	}
	if query.FBAfees != "" {
		url += "&fbafees=" + query.FBAfees
	}
	if query.Rental != "" {
		url += "&rental=" + query.Rental
	}

	resp, err := http.Get(url)
	if u.LogError(err) {
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if u.LogError(err) {
		return
	}

	err = json.Unmarshal(bytes, r)
	if u.LogError(err) {
		return
	}

	return
}
