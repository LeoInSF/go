package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/politics.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/opinions.xml
</loc>
</sitemap>
*/

// type Location struct {
// 	Loc string `xml:"loc"`
// }

/*
[5] type == array
[]int == slice
*/

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

func main() {
	var s SitemapIndex
	var n News
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	newsMap := make(map[string]NewsMap)

	for _, Location := range s.Locations {
		fmt.Println(string(Location))

		tmpResp, err := http.Get(strings.TrimSpace(Location))

		if err != nil {
			panic(err)
		}
		bytes, _ := ioutil.ReadAll(tmpResp.Body)
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Keywords {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}

		for idx, data := range newsMap {
			fmt.Println("\n\n\n", idx)
			fmt.Println("\n", data.Keyword)
			fmt.Println("\n", data.Location)
		}
	}
}
