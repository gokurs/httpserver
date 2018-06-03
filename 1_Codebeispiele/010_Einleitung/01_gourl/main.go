package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var (
	flagURL       = flag.String("url", "", "the url to load")
	flagPrintBody = flag.Bool("body", false, "when set the http body is printed out")
	flagParseBody = flag.Bool("parse", false, "parses the body")
)

func main() {
	flag.Parse()
	resp, err := http.Get(*flagURL)
	if err != nil {
		fmt.Printf("error loading %s\n%s", *flagURL, err)
	}
	defer resp.Body.Close()
	switch {
	case *flagPrintBody:
		_, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading from response body: ", err)
		}

	case *flagParseBody:
		parseResponse(resp)

	default:
		responseOutput(resp)
	}

}

func responseOutput(resp *http.Response) {
	fmt.Fprintln(os.Stdout, "URL: ", *flagURL)
	fmt.Fprintln(os.Stdout, "Response:\n____________________")
	fmt.Fprintln(
		os.Stdout,
		formatResponseString(fmt.Sprintf("%#v", resp)),
	)
}

func formatResponseString(s string) string {
	s = strings.Replace(s, ", ", ",\n\n", -1)
	s = strings.Replace(s, ":", ":\n", -1)
	return s
}

func parseResponse(resp *http.Response) {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error parsing from response body: ", err)
	}
	var f func(*html.Node, string)
	skipData := map[string]bool{
		"script": true,
		"head":   true,
		"style":  true,
		"footer": true,
		"nav":    true,
	}
	f = func(n *html.Node, padding string) {
		switch d := strings.Trim(n.Data, " "); n.Type {
		case html.TextNode:
			if d == "" || strings.HasPrefix(d, "\n") {
				return
			}
			fmt.Fprintln(os.Stdout, padding, d)
		case html.ElementNode:
			if skipData[d] {
				return
			}
			padding = padding + n.Data + ">"

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, padding)
		}
	}
	f(doc, "")
}
