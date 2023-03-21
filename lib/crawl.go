package lib

import (
	// "net/http"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func Crawl(url string) {
	
	type NikeProduct struct {
		Title 					string 
		Category 				string
		Price 					string
		DiscountedPrice 		string
		Colors 					[]string
	}

	crawler := colly.NewCollector()
	nikeProduct := NikeProduct{}

	crawler.OnHTML("h1#pdp_product_title", func(element *colly.HTMLElement) {
		nikeProduct.Title = element.Text
	})

	crawler.OnHTML("h2.headline-5", func(element *colly.HTMLElement) {
		nikeProduct.Category = element.Text
	})
	
	crawler.OnHTML("div.product-price__wrapper", func(element *colly.HTMLElement) {
		
		childNodes := element.DOM.Children().Nodes
		if len(childNodes) == 2 {
			nikeProduct.Price = element.DOM.FindNodes(childNodes[0]).Text()
			discounted_price := element.DOM.FindNodes(childNodes[1]).Text()
			nikeProduct.DiscountedPrice = strings.Replace(discounted_price, "Discounted from", "", 1)
		} else {
			nikeProduct.Price = element.DOM.FindNodes(childNodes[0]).Text()
			nikeProduct.DiscountedPrice = "0"
		}
	})

	crawler.OnHTML("div.colorway-container", func(element *colly.HTMLElement) {
		src, is_exist := element.DOM.Find("img").Attr("src")
		if !is_exist {
			log.Fatal("is not exist")
		}
		nikeProduct.Colors = append(nikeProduct.Colors, src)
	})

	crawler.OnHTML("div#pdp-6-up button", func(element *colly.HTMLElement) {
		pics := element.DOM.Find("picture")
		picNodes := pics.Nodes
		if len(picNodes) == 2 {
			src, is_exist := element.DOM.FindNodes(picNodes[1]).Attr("src")
			if !is_exist {
				fmt.Println("is not exist")
			}
			fmt.Println(src)
		}
	})

	crawler.Visit(url)

}

func UrlProcessing(url string) map[string]interface{} {
	var urlArray []string 
	
	urlArray = strings.Split(url, "/")[4:]
	urlArray = urlArray[:len(urlArray)-1]
	firstElementSplited := strings.Split(urlArray[0], "-")
	category := firstElementSplited[len(firstElementSplited)-2]
	fmt.Println("category: ", category)

	data := make(map[string]interface{}, 0)
	return data
}
