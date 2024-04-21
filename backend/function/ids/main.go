package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Wiki struct {
	Title string
	URL   string
}

func main() {
	start := time.Now()

	scanner := bufio.NewScanner(os.Stdin)

	// Display a message to prompt user input
	fmt.Print("Enter a string: ")

	// Read input from the user
	scanner.Scan()

	// Get the string from the input
	inputString := strings.TrimSpace(scanner.Text())

	scanner.Scan()

	//Kayak King_Zhou_of_Shang
	targetTitle := strings.TrimSpace(scanner.Text())

	targetURL := getWikiURL(targetTitle)

	// found, path := IDS(inputString, targetTitle, 3, []Wiki{})

	found := false
	depth := 0
	var path []Wiki

	for !found {

		found, path = IDS(inputString, targetURL, depth, []Wiki{})
		fmt.Println("Depth: ", depth)

		depth++

		if depth > 3 {
			break
		}
	}

	if found {
		fmt.Println("Path found:")
		printPath(path)
	} else {
		fmt.Println("Path not found within the depth limit")
	}

	elapsed := time.Since(start) // Hitung waktu yang telah berlalu
	fmt.Printf("Runtime: %s\n", elapsed)
}

func isValidWikiLink(href string) bool {
	return strings.HasPrefix(href, "/wiki/") && !strings.Contains(href, ":") && !strings.Contains(href, "Main_Page") && !strings.Contains(href, "#")
}

func scrapeWikiPage(inputString string) []Wiki {
	var links []Wiki

	url := "https://en.wikipedia.org/wiki/" + strings.ReplaceAll(inputString, " ", "_")
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil
		// log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && isValidWikiLink(href) {
			links = append(links, Wiki{
				Title: s.Text(),
				URL:   "https://en.wikipedia.org" + href,
			})
		}
	})

	return links
}

func IDS(sourceTitle string, targetURL string, depth int, path []Wiki) (bool, []Wiki) {
	printPath(path)

	if getWikiURL(sourceTitle) == targetURL {
		return true, append(path, Wiki{Title: sourceTitle, URL: "https://en.wikipedia.org/wiki/" + sourceTitle})
	}

	if depth == 0 {
		return false, nil
	}

	links := scrapeWikiPage(sourceTitle)

	for _, tempLink := range links {
		if tempLink.URL == targetURL {
			return true, append(path, Wiki{Title: tempLink.Title, URL: tempLink.URL})
		}

		found, newPath := IDS(tempLink.Title, targetURL, depth-1, append(path, Wiki{Title: tempLink.Title, URL: tempLink.URL}))
		if found {
			return true, newPath
		}
	}

	return false, nil
}

func printPath(path []Wiki) {
	for i, title := range path {
		if i > 0 {
			fmt.Print(" -> ")
		}
		fmt.Print(title.Title)
	}
	fmt.Println()
}

func getWikiURL(title string) string {
	url := "https://en.wikipedia.org/wiki/" + strings.ReplaceAll(title, " ", "_")
	return url
}
