package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Wiki struct {
	Title string
	URL   string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Menampilkan pesan untuk meminta input dari pengguna
	fmt.Print("Masukkan sebuah string: ")

	// Membaca input dari pengguna
	scanner.Scan()

	// Mengambil string dari input yang telah dibaca
	inputString := scanner.Text()

	res, err := http.Get("https://en.wikipedia.org/wiki/" + strings.ReplaceAll(inputString, " ", "_"))

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var links []Wiki

	// Menemukan semua elemen <a> dan mengambil atribut href-nya
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// Mendapatkan nilai dari atribut href
		href, exists := s.Attr("href")
		if exists {
			// Memeriksa apakah tautan mengarah ke artikel Wikipedia
			if isValidWikiLink(href) {
				// Menambahkan tautan ke array links
				links = append(links, Wiki{
					Title: s.Text(),
					URL:   "https://en.wikipedia.org" + href,
				})
			}
		}
	})

	fmt.Println("Tautan Wikipedia:")
	for _, link := range links {
		fmt.Printf("%s: %s\n", link.Title, link.URL)
	}
}

func isValidWikiLink(href string) bool {
	// Check if the link starts with "/wiki/" and not contains any ":" which indicates non-article links
	return strings.HasPrefix(href, "/wiki/") && !strings.Contains(href, ":")
}
