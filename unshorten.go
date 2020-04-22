package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var errTooManyRedirects = errors.New("too many redirects")

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] URL\n", os.Args[0])
	fmt.Printf("Unshorten the given URL, printing each redirect followed along the way.\n\n")
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
	fmt.Printf("\nIssues:\n  https://github.com/cdzombak/unshorten/issues/new\n")
	fmt.Printf("\nAuthor: Chris Dzombak <https://www.dzombak.com>\n")
}

func main() {
	quiet := flag.Bool("quiet", false, "Run quietly; only display the final URL.")
	maxRedirects := flag.Int("max-redirects", 10, "Maximum number of redirects to follow.")
	flag.Usage = usage
	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	url := flag.Args()[0]
	if !(strings.HasPrefix(strings.ToLower(url), "http://") || strings.HasPrefix(strings.ToLower(url), "https://")) {
		url = fmt.Sprintf("http://%s", url)
	}
	redirects := 0
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if redirects >= *maxRedirects {
				return errTooManyRedirects
			}
			redirects++
			if !*quiet {
				fmt.Println(req.URL)
			}
			return nil
		},
	}
	if !*quiet {
		fmt.Println(url)
	}
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		fmt.Printf("Error building initial request: %s\n", err)
		os.Exit(1)
	}
	resp, err := client.Do(req)
	if errors.Is(err, errTooManyRedirects) {
		fmt.Printf("Too many redirects: followed %d (customize with -max-redirects).\n", redirects)
		os.Exit(2)
	} else if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	if *quiet {
		fmt.Println(resp.Request.URL)
	}
}
