package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type CurrencyExchange struct {
	Buy  string
	Sell string
}

func (c *CurrencyExchange) GetFormattedMessage() string {
	return fmt.Sprintf(`
	KAMBISTA 🚀
	VENTA -> %s
	COMPRA -> %s
	with ❤️ by resep.
	`, c.Sell, c.Buy)
}

func ScrapeExchange() (CurrencyExchange, error) {
	return scrapeKambista()
}

func scrapeKambista() (CurrencyExchange, error) {
	fmt.Println("scraping kambista.com")
	baseURL := "https://kambista.com/"
	exchange := CurrencyExchange{}

	c := colly.NewCollector()

	c.OnHTML("strong[id=valcompra]", func(e *colly.HTMLElement) {
		text := e.Text
		exchange.Buy = text
	})

	c.OnHTML("strong[id=valventa]", func(e *colly.HTMLElement) {
		text := e.Text
		exchange.Sell = text
	})

	c.Visit(baseURL)

	return exchange, nil
}
