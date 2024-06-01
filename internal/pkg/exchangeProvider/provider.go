// Package exchange provides a client for interacting with NBP API
// It allows for dynamic currency exchange
package exchange

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pstano1/go-cart/internal/pkg"
	"go.uber.org/zap"
)

type IExchangeProvider interface {
	FetchNBPTable(tableName string) error
	GetExchangeRate(from, to string) (float32, error)
}

type ExchangeProvider struct {
	log           *zap.Logger
	baseURL       string
	exchangeRates map[string]float32
}

func NewProvider(logger *zap.Logger) IExchangeProvider {
	return &ExchangeProvider{
		log:           logger,
		baseURL:       "http://api.nbp.pl/api",
		exchangeRates: make(map[string]float32),
	}
}

// FetchNBPTable fetches current NBP exhcange rates for specified table
// then parses them & stores in memory
func (e *ExchangeProvider) FetchNBPTable(tableName string) error {
	e.log.Debug("Pulling exchange rates from NBP table",
		zap.String("table", tableName),
	)
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/exchangerates/tables/%s/last/1/?format=json", e.baseURL, tableName), nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		print(res.StatusCode)
		return pkg.ErrStatusNotOK
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var exchangesTable []pkg.NBPTable
	if err = json.Unmarshal(body, &exchangesTable); err != nil {
		return err
	}
	if len(exchangesTable) != 1 {
		return nil
	}
	for _, rate := range exchangesTable[0].Rates {
		e.exchangeRates[rate.CurrencyCode] = rate.AskRate
	}

	return nil
}

// GetExchangeRate returns exchnage rate for given currency pair
// (given it's provided by NBP)
func (e *ExchangeProvider) GetExchangeRate(from, to string) (float32, error) {
	if to == "PLN" {
		rate, ok := e.exchangeRates[from]
		if !ok {
			return -1, pkg.ErrBaseCurrencyNotAvailable
		}
		return rate, nil
	}
	fromRate, ok := e.exchangeRates[from]
	if !ok {
		return -1, pkg.ErrBaseCurrencyNotAvailable
	}
	toRate, ok := e.exchangeRates[to]
	if !ok {
		return -1, pkg.ErrCurrencyNotAvailable
	}
	return fromRate / toRate, nil
}
