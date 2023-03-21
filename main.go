package main

import (
	"localPackage/lib"
)

func main() {
	product := lib.Crawl("https://www.nike.com/t/air-max-plus-mens-shoes-x9G2xF/604133-139")
	lib.ToCSV(product, "./nike_product.json")
}