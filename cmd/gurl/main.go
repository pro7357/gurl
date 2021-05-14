package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/pro7357/gurl"
)

func main() {
	flag.Parse()

	urls := flag.Args()
	if len(urls) != 0 {
		for _, url := range urls {
			gurl.Gurl(url)
		}
	} else {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			url := sc.Text()
			gurl.Gurl(url)
		}
	}
}
