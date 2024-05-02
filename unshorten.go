package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	errTooManyRedirects = errors.New("too many redirects")

	version = "<dev>"
)

func efprintf(f string, a ...any) {
	_, _ = fmt.Fprintf(os.Stderr, f, a...)
}

func usage() {
	efprintf("Usage: %s [OPTIONS] URL\n", filepath.Base(os.Args[0]))
	efprintf("Unshorten the given URL, printing each redirect followed along the way.\n\n")
	efprintf("Options:\n")
	flag.PrintDefaults()
	efprintf("\nGitHub:\n  https://github.com/cdzombak/unshorten\n")
	efprintf("\nAuthor: Chris Dzombak <https://www.dzombak.com>\n")
}

func main() {
	quiet := flag.Bool("quiet", false, "Run quietly; only display the final URL.")
	maxRedirects := flag.Int("max-redirects", 10, "Maximum number of redirects to follow.")
	printVersion := flag.Bool("version", false, "Print version and exit.")
	flag.Usage = usage
	flag.Parse()
	if *printVersion {
		fmt.Println(version)
		os.Exit(0)
	}
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
		CheckRedirect: func(req *http.Request, _ []*http.Request) error {
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
