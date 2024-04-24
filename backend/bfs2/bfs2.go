package bfs2

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

var limit int = 100

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

func scrap(urls []string, target string) ([]WikiTitle, error) {
    fmt.Println("Scraping", urls, "...")
    var wikis []WikiTitle
    parentLinks := sync.Map{}

    c := colly.NewCollector(
        colly.AllowedDomains("en.wikipedia.org"),
    )

    q, err := queue.New(limit,
        &queue.InMemoryQueueStorage{MaxSize: 10000},
    )
    if err != nil {
        return nil, err
    }

    for _, url := range urls {
        q.AddURL("https://en.wikipedia.org/wiki/" + url)
    }
    // targetFound := make(chan struct{}, 1)

    c.OnHTML("a[href^='/wiki/']:not([href*=':']):not([href*='%']):not([href*='#']):not([href*='ISO']):not([href='/wiki/Main_Page'])", func(e *colly.HTMLElement) {
        link := strings.TrimPrefix(e.Attr("href"), "/wiki/")
        parent := strings.TrimPrefix(e.Request.URL.String(), "https://en.wikipedia.org/wiki/")
        parentLinks.Store(link, parent)
        wikis = append(wikis, WikiTitle{link, parent})

        if link == target {
            fmt.Println("Target found")
            // targetFound <- struct{}{}
        }
    })


    // c.OnRequest(func(r *colly.Request) {
    //     select {
    //     case <-targetFound:
    //         r.Abort()
    //     default:

    //     }
    // })

    err = q.Run(c)
    if err != nil {
        return nil, err
    }

    
    //fmt.Println("Scraping", urls, "done.")
    return wikis, nil
}

func findPath(source, target string) ([]Wiki, error) {
    parent := make(map[string]string)
    queue := []string{source}
    found := false

    for !found{
        
        numScraping := min(len(queue), limit)
        currentSet := []string{}

		for i := 0; i < numScraping; i++ {
			currentSet = append(currentSet, queue[i])
		}

		wikis, _ := scrap(currentSet, target)
        queue = queue[numScraping:]

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

		fmt.Println(len(parent))
        
    }

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

func constructURL(title string) string {
    return strings.ReplaceAll(title, " ", "_")
}

func Entrypoint(sourceTitle, targetTitle string) ([]Wiki, Duration, error) {

	if isValidWiki(sourceTitle) && isValidWiki(targetTitle) {

		startTime := time.Now()
        sourceTitle = constructURL(sourceTitle)
        targetTitle = constructURL(targetTitle)
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
