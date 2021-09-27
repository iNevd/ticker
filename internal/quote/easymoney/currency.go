package easymoney

import (
	c "github.com/achannarasappa/ticker/internal/common"
	"github.com/go-resty/resty/v2"
)

// GetCurrencyRates retrieves the currency rates to convert from each currency for the given symbols to the target currency
func GetCurrencyRates(client resty.Client, symbols []string, targetCurrency string) (c.CurrencyRates, error) {
	return c.CurrencyRates{}, nil // 固定不获取汇率信息
}
