package main

import (
	"bufio"
	"flag"
	"sync"
	"fmt"
	"time"
	"os"
	"strconv"
	"strings"
	"net/http"
	"log"
)

func main () {
	stat := time-now()
	var counter = 0
	var concurrency = 10
	for i:=1; i<=concurrency; i++ {
		go func() {
			time.Sleep(time.Millisecond * 50)
			counter++
		}()
	}
}

var googleUrls = map[string]string{
	"com": "https://www.google.com/search?q=",
 	"de": "https://www.google.de/search?q=",
	"uk": "https://www.google.co.uk/search?q=",
	"fr": "https://www.google.fr/search?q=",
}
