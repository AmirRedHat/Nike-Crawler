package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type NikeProduct struct {
	Title 					string 
	Category 				string
	Price 					string
	DiscountedPrice 		string
	Colors 					[]string
	Images					[]string
	Sizes 					[]string
}

func Crawl(url string) NikeProduct {

	crawler := colly.NewCollector()

	var title, category, price, discounted_price string
	var colors, images, sizes []string

	// // fetching title 
	// crawler.OnHTML("h1#pdp_product_title", func(element *colly.HTMLElement) {
	// 	title = element.Text
	// })

	// // fetching category
	// crawler.OnHTML("h2.headline-5", func(element *colly.HTMLElement) {
	// 	category = element.Text
	// })
	
	// // fetching price
	// crawler.OnHTML("div.product-price__wrapper", func(element *colly.HTMLElement) {
		
	// 	childNodes := element.DOM.Children().Nodes
	// 	if len(childNodes) == 2 {
	// 		price = element.DOM.FindNodes(childNodes[0]).Text()
	// 		discounted_price = element.DOM.FindNodes(childNodes[1]).Text()
	// 		discounted_price = strings.Replace(discounted_price, "Discounted from", "", 1)
	// 	} else {
	// 		price = element.DOM.FindNodes(childNodes[0]).Text()
	// 		discounted_price = "$0"
	// 	}
	// })

	// fetching colors
	crawler.OnHTML("div.colorway-container", func(element *colly.HTMLElement) {
		src, is_exist := element.DOM.Find("img").Attr("src")
		if !is_exist {
			log.Fatal("is not exist")
		}
		colors = append(colors, src)
	})

	// fetching images
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

	// fetching sizes
	crawler.OnHTML("script#__NEXT_DATA__", func(element *colly.HTMLElement) {
		var sepURL []string 
		var productCode string
		var productFileName string
		var result map[string]interface{}
		var productData map[string]interface{}
		
		sepURL = strings.Split(url, "/")
		productCode = sepURL[len(sepURL)-1]
		json.Unmarshal([]byte(element.Text), &result)
		productData = result["props"].(map[string]interface{})["pageProps"].(map[string]interface{})["initialState"].(map[string]interface{})["Threads"].(map[string]interface{})["products"].(map[string]interface{})[productCode].(map[string]interface{})
		
		// save product data in file
		jsonProductData, _ := json.Marshal(productData)
		productFileName = fmt.Sprintf("./%s.json", productCode)
		ioutil.WriteFile(productFileName,  jsonProductData, 0644)
		
		title = productData["fullTitle"].(string)
		category = productData["subTitle"].(string)
		price = fmt.Sprintf("%f", productData["fullPrice"].(float64))
		discounted_price = fmt.Sprintf("%f", productData["currentPrice"].(float64))

		sizeList := productData["skus"].([]interface{})		
		for i := range sizeList {
			size := sizeList[i].(map[string]interface{})["localizedSize"].(string)
			sizes = append(sizes, size)
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
	nikeProduct.Sizes 				= sizes

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

func TestNike() {


	// rawJson := `{
	// 	"packet":{
	// 		 "hostname":"host1",
	// 		 "pid":"123435",
	// 		 "processname":"process",
	// 		 "type":"partial"}
	// 	}`

	// ioutil.WriteFile("./test2.json", []byte(rawJson), 0644)

	content, err := ioutil.ReadFile("./test.json")
	if err != nil {
		fmt.Println("reading file error")
		log.Fatal(err)
	}

	fmt.Println("file opened successfully")

	var res map[string]interface{}
	
	errr := json.Unmarshal(content, &res)
	if errr != nil {
		log.Fatal(errr)
	}
	result := res["props"].(map[string]interface{})["pageProps"].(map[string]interface{})["initialState"].(map[string]interface{})["Threads"].(map[string]interface{})["products"].(map[string]interface{})["DV0804-200"].(map[string]interface{})["skus"].([]interface{})
	for i := range result {
		fmt.Println(result[i].(map[string]interface{})["nikeSize"])
	}


	// var data map[string]interface{}
	// jsonerr := json.NewDecoder(file).Decode(&data)
	// if jsonerr != nil {
	// 	fmt.Println("converting json error")
	// 	log.Fatal(jsonerr)
	// }

	// fmt.Println(data)
}
