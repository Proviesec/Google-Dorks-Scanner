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
