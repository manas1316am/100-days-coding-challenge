package main

import "fmt"

const GlobalLimit = 100
const MaxCacheSize int = 10 * GlobalLimit

const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

var cache map[string]string

// getting cache
func cacheGet(key string) string {
	return cache[key]
}

// setting cache
func cacheSet(key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

// getting book
func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

// setting book
func SetBook(isbn, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

// Getting CD
func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

// Setting CD
func SetCD(sku string, title string) {
	cacheSet(CacheKeyCD+sku, title)
}

func main() {
	cache = make(map[string]string)

	SetBook("1234-5678", "Get Ready To GO")
	SetCD("1234-4567", "Get Ready to audio-book")

	fmt.Println("Book:", GetBook("1234-5678"))
	fmt.Println("CD:", GetBook("1234-4567"))

}
