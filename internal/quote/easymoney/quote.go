package easymoney

import (
	"fmt"
	c "github.com/achannarasappa/ticker/internal/common"
	"github.com/go-resty/resty/v2"
	"strconv"
)

// Response represents the container object from the API response
type Response struct {
	Data map[string]interface{} `json:"data"`
}

func (r Response) TransferToQuote() c.AssetQuote {
	atoi := func(i interface{}) int {
		ii, _ := strconv.Atoi(fmt.Sprint(i))
		return ii
	}

	assetQuote := c.AssetQuote{
		Name:   fmt.Sprint(r.Data["f58"]),
		Symbol: fmt.Sprint(r.Data["f57"]),
		Class:  c.AssetClassStock,
		Currency: c.Currency{
			FromCurrencyCode: "CNY",
		},
		QuotePrice: c.QuotePrice{
			Price:          float64(atoi(r.Data["f43"])) / 100,
			PricePrevClose: float64(atoi(r.Data["f60"])) / 100,
			PriceOpen:      float64(atoi(r.Data["f46"])) / 100,
			PriceDayHigh:   float64(atoi(r.Data["f44"])) / 100,
			PriceDayLow:    float64(atoi(r.Data["f45"])) / 100,
			Change:         float64(atoi(r.Data["f169"])) / 100,
			ChangePercent:  float64(atoi(r.Data["f170"])) / 100,
		},
		QuoteExtended: c.QuoteExtended{
			FiftyTwoWeekHigh: float64(atoi(r.Data["f174"])) / 100,
			FiftyTwoWeekLow:  float64(atoi(r.Data["f175"])) / 100,
			MarketCap:        float64(atoi(r.Data["116"])),
			Volume:           float64(atoi(r.Data["f47"])) / 100,
		},
		QuoteSource: c.QuoteSourceEastMoney,
		Exchange: c.Exchange{
			Name:                    "东方财富",
			Delay:                   0,
			State:                   c.ExchangeStateOpen,
			IsActive:                true,
			IsRegularTradingSession: true,
		},
		Meta: c.Meta{},
	}
	return assetQuote
}

// GetAssetQuotes issues a HTTP request to retrieve quotes from the API and process the response
func GetAssetQuotes(client resty.Client, symbols []string) func() []c.AssetQuote {
	return func() []c.AssetQuote {
		var assetQuotes []c.AssetQuote
		for _, symbol := range symbols {
			url := fmt.Sprintf("http://push2.eastmoney.com/api/qt/stock/get?fields=f12,f57,f58,f43,f44,f45,f46,f47,f48,f60,f162,f167,f168,f169,f170,f174,f175&secid=%s", symbol)
			res, err := client.R().
				SetResult(Response{}).
				Get(url)
			if err == nil {
				assetQuote := res.Result().(*Response).TransferToQuote()
				assetQuotes = append(assetQuotes, assetQuote)
			}
		}
		return assetQuotes
	}
}
