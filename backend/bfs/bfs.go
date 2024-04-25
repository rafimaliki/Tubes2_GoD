package bfs

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

type Wiki struct {
	Title string
	URL   string
}

type WikiTitle struct {
	Title string
	Parent string
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


var cache = sync.Map{}


func scrap(url string, ch chan<-[]WikiTitle, wg *sync.WaitGroup){

    defer wg.Done()  

	fmt.Println("Scraping", url, "...")


    if links, ok := cache.Load(url); ok {
        ch <- links.([]WikiTitle)
        return
    }

    c := colly.NewCollector(
        colly.AllowedDomains("en.wikipedia.org"),
    )

    var wikis []WikiTitle

    c.OnHTML("a[href^='/wiki/']:not([href*=':']):not([href*='%']):not([href*='#']):not([href='/wiki/Main_Page'])", func(e *colly.HTMLElement) {
        link := strings.TrimPrefix(e.Attr("href"), "/wiki/")
        wikis = append(wikis, WikiTitle{link, url})
    })

    c.Visit("https://en.wikipedia.org/wiki/" + url)

    cache.Store(url, wikis)

	ch <- wikis
	fmt.Println("Scraping", url, "done.")
}

func findPath(source, target string) ([]Wiki, error) {
    parent := make(map[string]string)
    queue := []string{source}
    found := false
    var wg sync.WaitGroup
	ch := make(chan []WikiTitle)

    for !found && len(queue) > 0 {
        
        numScraping := min(len(queue), 4)
        currentSet := []string{}
        for i := 0; i < numScraping; i++ {
            current := queue[i]
            currentSet = append(currentSet, current)
            wg.Add(1)
            go scrap(current, ch, &wg)
            //time.Sleep(2 * time.Millisecond)
        }

        
        go func() {
            fmt.Println("Waiting scrap: ", currentSet)
            wg.Wait()
            fmt.Println("Scrapping finished: ", currentSet)
        }()

        queue = queue[numScraping:]

        wikis := <-ch
        for _, wiki := range wikis {
            if wiki.Title == target {
                parent[wiki.Title] = wiki.Parent
                found = true
                break
            }

            if _, ok := parent[wiki.Title]; !ok {
                parent[wiki.Title] = wiki.Parent
                queue = append(queue, wiki.Title)
                fmt.Printf("%d.  %s, \033[32m→\033[0m %s\n", len(parent), wiki.Parent, wiki.Title)
            }
        }
        
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    if found {
        pathTitle := []string{target}
        current := target
        for current != source {
            current = parent[current]
            pathTitle = append([]string{current}, pathTitle...)
        }

        var path []Wiki
        for _, title := range pathTitle {
            path = append(path, Wiki{title, getURL(title)})
        }

        return path, nil
    }
    return nil, fmt.Errorf("target Wiki not found")
}

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