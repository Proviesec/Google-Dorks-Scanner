import (
	"context"
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
)

type Result struct {
	Rank int `json:"rank"`
	URL string `json:"url"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type SearchOptions struct {
	CountryCode string
	LanguageCode string
	Limit int
	Start int
	// Default: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"
	UserAgent string
	ProxyAddr string
}

var googleUrls = map[string]string{
	"com": "https://www.google.com/search?q=",
 	"de": "https://www.google.de/search?q=",
	"uk": "https://www.google.co.uk/search?q=",
	"fr": "https://www.google.fr/search?q=",
}
