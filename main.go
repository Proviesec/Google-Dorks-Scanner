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
  var concurrency = 10
  for i:=0; i<*concurrency/2; i++ {
    wg.Add(1)
    go func() {
      time.Sleep(time.Millisecond * 50)
      wg.Done()
    }()
    wg.Wait()
  }
}

var googleUrls = map[string]string{
	"com": "https://www.google.com/search?q=",
  "de": "https://www.google.de/search?q=",
	"uk": "https://www.google.co.uk/search?q=",
	"fr": "https://www.google.fr/search?q=",
}
