package main

import (
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

func Gurl(url string) {
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background())
	if strings.HasPrefix(proxy, "socks5:") {
		opts := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.ProxyServer(proxy),
		)
		allocCtx, cancel = chromedp.NewExecAllocator(context.Background(), opts...)
		defer cancel()
	}
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
