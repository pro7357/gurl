package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

var (
	userJS string
	proxy  string
)

func init() {
	flag.StringVar(&userJS, "j", `[...new Set([... document.links].map(n => n.href))].join("\n")`, "the JS to run on each page")
	flag.StringVar(&proxy, "p", os.Getenv("HTTPS_PROXY"), "proxy setting")

	log.SetOutput(os.Stdout)
	log.SetFlags(0)
}

func main() {
	flag.Parse()

	urls := flag.Args()
	if len(urls) != 0 {
		for _, url := range urls {
			Gurl(url)
		}
	} else {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			url := sc.Text()
			Gurl(url)
		}
	}
}
