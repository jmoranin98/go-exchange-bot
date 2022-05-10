package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type CurrencyExchange struct {
	Buy  string
	Sell string
}

type kambistaResponse struct {
	TC kambistaTC `json:"tc"`
}

type kambistaTC struct {
	Bid float64 `json:"bid"`
	Ask float64 `json:"ask"`
}

func (c *CurrencyExchange) GetFormattedMessage() string {
	return fmt.Sprintf(`
	KAMBISTA 🚀
	VENTA -> %s
	COMPRA -> %s
	with ❤️ by resep.
	`, c.Sell, c.Buy)
}

func ScrapeExchange() (*CurrencyExchange, error) {
	return scrapeKambista()
}

func scrapeKambista() (*CurrencyExchange, error) {
	var kambistaRes kambistaResponse

	apiURL := "https://api.kambista.com/v1/exchange/calculates?originCurrency=USD&destinationCurrency=PEN&amount=1&active=S"
	res, err := http.Get(apiURL)

	if err != nil {
		return nil, errors.New("cannot get kambista exchange")
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&kambistaRes)
	if err != nil {
		return nil, err
	}

	return &CurrencyExchange{
		Buy:  fmt.Sprintf("%v", kambistaRes.TC.Bid),
		Sell: fmt.Sprintf("%v", kambistaRes.TC.Ask),
	}, nil
}
