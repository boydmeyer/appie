package appie

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Appie struct {}

type Product struct {
	ID string
	Name string
	Price string
}

func New() (*Appie, error) {
	return &Appie{}, nil
}

func (a *Appie) GetProduct(id string) (*Product, error) {
	url, _ := a.generateUrl(id)
	c := colly.NewCollector()

	p := Product{ID: id}

	// Name
	c.OnHTML(".product-card-header_title__11vVs", func(e *colly.HTMLElement) {
		p.Name = e.Text
	})

	// Price
	c.OnHTML("div.product-card-hero-price_now__PlF9u", func(e *colly.HTMLElement) {
		p.Price = e.Text
	})

	c.Visit(url)
	return &p, nil
}

func (a *Appie) generateUrl(id string) (string, error) {
	return fmt.Sprintf("https://www.ah.nl/producten/product/wi%s", id), nil
}

