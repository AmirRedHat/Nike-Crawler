package lib

import (
	"fmt"
	"log"
	"strings"
	"io/ioutil"
	"encoding/json"
	"github.com/gocolly/colly"
)

type NikeProduct struct {
	Title 					string 
	Category 				string
	Price 					string
	DiscountedPrice 		string
	Colors 					[]string
	Images					[]string
}

func Crawl(url string) NikeProduct {

	crawler := colly.NewCollector()

	var title, category, price, discounted_price string
	var colors, images []string

	crawler.OnHTML("h1#pdp_product_title", func(element *colly.HTMLElement) {
		title = element.Text
	})

	crawler.OnHTML("h2.headline-5", func(element *colly.HTMLElement) {
		category = element.Text
	})
	
	crawler.OnHTML("div.product-price__wrapper", func(element *colly.HTMLElement) {
		
		childNodes := element.DOM.Children().Nodes
		if len(childNodes) == 2 {
			price = element.DOM.FindNodes(childNodes[0]).Text()
			discounted_price = element.DOM.FindNodes(childNodes[1]).Text()
			discounted_price = strings.Replace(discounted_price, "Discounted from", "", 1)
		} else {
			price = element.DOM.FindNodes(childNodes[0]).Text()
			discounted_price = "$0"
		}
	})

	crawler.OnHTML("div.colorway-container", func(element *colly.HTMLElement) {
		src, is_exist := element.DOM.Find("img").Attr("src")
		if !is_exist {
			log.Fatal("is not exist")
		}
		colors = append(colors, src)
	})

	crawler.OnHTML("div#pdp-6-up div.css-du206p", func(element *colly.HTMLElement) {
		pics := element.DOM.Find("picture")
		if len(pics.Nodes) == 2 {
			src, is_exist := element.DOM.FindNodes(pics.Nodes[1]).Find("img").Attr("src")
			if !is_exist {
				log.Fatal("is not exist")
			}
			images = append(images, src)
		}
	})

	crawler.Visit(url)
	
	nikeProduct := NikeProduct{}
	nikeProduct.Title 				= title
	nikeProduct.Category 			= category
	nikeProduct.Price 				= price
	nikeProduct.DiscountedPrice 	= discounted_price
	nikeProduct.Colors 				= colors
	nikeProduct.Images 				= images

	return nikeProduct

}

func ToCSV(nikeProduct NikeProduct, fileName string) { 
	jsonData, err := json.MarshalIndent([]NikeProduct{nikeProduct}, "", "")
	if err != nil {
		log.Fatal("unable to save this product")
	}
	ioutil.WriteFile(fileName, jsonData, 0644)
	fmt.Println("file saved!")
}
