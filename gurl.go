package gurl

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/chromedp/chromedp"
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

func Gurl(url string) {
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background())
	if strings.HasPrefix(proxy, "socks5:") {
		opts := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.ProxyServer(proxy),
		)
		allocCtx, cancel = chromedp.NewExecAllocator(context.Background(), opts...)
	}
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Evaluate(userJS, &res),
	)
	if err != nil {
		log.Printf("error on %s : %s", url, err)
	}

	log.Println(res)
}
