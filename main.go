// Copyright 2022 Proviesec
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
func usage() {
	fmt.Fprint(os.Stderr, `Usage: Google Dork Scanner [url] [flag]
A google dorks scanner. The scanner sent the dorks to google to find security vulnerabilities. 
Flags:
    -v, --verbose  Print verbose logs to stderr.
`)
}

func main () {
	var url string
	var v bool
	flag.StringVar(&url, "url", "", "")
	flag.BoolVar(&v, "v", false, "")
	
	start := time.Now()
	var counter = 0
	var concurrency = 10
	for i:=1; i<=concurrency; i++ {
		go func() {
			time.Sleep(time.Millisecond * 50)
			counter++
		}()
	}
	time.Sleep(2 * time.Second)
	elapsed := time.Since(start)
	log.Println(elapsed)
}

var googleUrls = map[string]string{
	"com": "https://www.google.com/search?q=",
 	"de": "https://www.google.de/search?q=",
	"uk": "https://www.google.co.uk/search?q=",
	"fr": "https://www.google.fr/search?q=",
}
