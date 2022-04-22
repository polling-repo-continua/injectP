package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func main() {
	payload := os.Args[1]
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		u, err := url.Parse(sc.Text())
		if err != nil {
			continue
		}
		regex, _ := regexp.MatchString("[?]", u.String())
		if regex == true {
			inject(u.String(), payload)
		}
	}

}

func inject(link, payload string) {
	urlList := []string{}
	paramLink := strings.Split(link, "?")[1]

	regex2, _ := regexp.MatchString("&", paramLink)
	if regex2 == true {
		paramslinks := strings.Split(paramLink, "&")
		for _, prmt := range paramslinks {
			key := strings.Split(prmt, "=")[0]
			injectedParam := key + "=" + payload
			line := strings.ReplaceAll(link, prmt, injectedParam)
			urlList = append(urlList, line)

		}

	} else {
		key2 := strings.Split(paramLink, "=")[0]
		injectedParam2 := key2 + "=" + payload
		line := strings.ReplaceAll(link, paramLink, injectedParam2)
		urlList = append(urlList, line)
	}
	for _, p := range urlList {
		fmt.Println(p)
	}
}
