package bfs

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type Wiki struct {
	Title string
	URL   string
}
type Duration struct {
    Hours       int
    Minutes     int
    Seconds     int
    Milliseconds int
}


func getDuration(duration time.Duration) Duration {
    hours := int(duration.Hours())
    minutes := int(duration.Minutes()) % 60
    seconds := int(duration.Seconds()) % 60
    milliseconds := int(duration.Nanoseconds() / 1e6) % 1000

    return Duration{
        Hours:       hours,
        Minutes:     minutes,
        Seconds:     seconds,
        Milliseconds: milliseconds,
    }
}

func isValidWiki(title string) bool {
	url := "https://en.wikipedia.org/wiki/" + title
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func scrap(url string) []string {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	var wikis []string
	visited := make(map[string]bool)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasPrefix(link, "/wiki/") && !strings.Contains(link, ":") && !strings.Contains(link, "%") && !strings.Contains(link, "#") && (link != "/wiki/Main_Page"){
			link = strings.TrimPrefix(link, "/wiki/")
			if !visited[link] {
				visited[link] = true
				wikis = append(wikis, link)
			}
		}
	})

	c.Visit("https://en.wikipedia.org/wiki/" + url)

	return wikis
}

func findPath(source, target string) ([]Wiki, error) {

	
	parent := make(map[string]string)
	queue := make([]string, 0, 250000)
	queue = append(queue, source)

	found := false

	for !found {
		current := queue[0]
		queue = queue[1:]

		wikis := scrap(current)

		for _, wiki := range wikis {
			if wiki == target {
				parent[wiki] = current
				found = true
				break
			}
			
			if _, ok := parent[wiki]; !ok {
				fmt.Printf("%d. \033[32mParent:\033[0m %s, \033[32mWiki:\033[0m %s\n", len(parent), current, wiki)
				parent[wiki] = current
				queue = append(queue, wiki)
			}
		}
	}

	if found {
		path_title := []string{target}
		current := target
		for current != source {
			current = parent[current]
			path_title = append([]string{current}, path_title...)
		}

		var path []Wiki
		for _, title := range path_title {
			path = append(path, Wiki{title, getURL(title)})
		}

		return path, nil
	} else {
		return nil, fmt.Errorf("target Wiki not found")
	}
}

// func getTitle(url string) string {
// 	return strings.TrimPrefix(url, "/wiki/")
// }

func getURL(title string) string {
	return "https://en.wikipedia.org/wiki/" + title
}

func Entrypoint(sourceTitle, targetTitle string) ([]Wiki, Duration, error) {

	if isValidWiki(sourceTitle) && isValidWiki(targetTitle) {

		startTime := time.Now()
		path, err := findPath(sourceTitle, targetTitle)
		duration := getDuration(time.Since(startTime))

		if err != nil {
			return nil, getDuration(0), err
		}

		fmt.Println("\033[93mPath:\033[0m")
		for _, wiki := range path {
			fmt.Println(wiki)
		}
		fmt.Println("\033[93mDuration:\033[0m")
		fmt.Println(duration)
		fmt.Println("")
		return path, duration, nil
	} else {
		return nil, getDuration(0), fmt.Errorf("invalid source or target Wiki")
	}
}
