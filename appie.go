package appie

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Appie struct {}

type Product struct {
	ID string
	Name string
	Description string
	Price string
	Image string
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

	// Description
	c.OnHTML("div[data-testhook='product-summary'] p", func(e *colly.HTMLElement) {
		p.Description = e.Text
	})

	// Price
	c.OnHTML("div[data-testhook='price-amount']", func(e *colly.HTMLElement) {
		p.Price = e.Text
	})

	// Image
	c.OnHTML("img[data-testhook='product-image']", func(e *colly.HTMLElement) {
		p.Image = e.Attr("src")
	})

	c.Visit(url)
	return &p, nil
}

func (a *Appie) generateUrl(id string) (string, error) {
	return fmt.Sprintf("https://www.ah.nl/producten/product/wi%s", id), nil
}

