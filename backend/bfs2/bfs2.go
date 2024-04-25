package bfs2

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"strings"
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

func scrap(urls []string) ([]WikiTitle) {
    fmt.Println("Scraping", len(urls), urls, "...")
    var wikis []WikiTitle

    c := colly.NewCollector(
        colly.AllowedDomains("en.wikipedia.org"), colly.Async(true),colly.CacheDir("./cache"),
    )
    c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: limit})
    
    q, err := queue.New(22,
        &queue.InMemoryQueueStorage{MaxSize: 10000},
    )

    if err != nil {
        return nil
    }

    for _, url := range urls {
        q.AddURL("https://en.wikipedia.org/wiki/" + url)
    }

    c.OnHTML("a[href^='/wiki/']:not([href*=':']):not([href='/wiki/Main_Page'])", func(e *colly.HTMLElement) {
        link,_ := url.PathUnescape(strings.TrimPrefix(e.Attr("href"), "/wiki/"))
        if idx := strings.Index(link, "#"); idx != -1 {
            link = link[:idx]
        }
        parent, _ := url.PathUnescape(strings.TrimPrefix(e.Request.URL.String(), "https://en.wikipedia.org/wiki/"))
        wikis = append(wikis, WikiTitle{link, parent})
    })

    err = q.Run(c)

    if err != nil {
        return nil
    }

    c.Wait()

    return wikis
}

func findPath(source, target string) ([]Wiki, error) {

    fmt.Println("Finding path from", source, "to", target, "...")

    startTime := time.Now()

    parent := make(map[string]string)
    queue := []string{source}
    found := false

    currDegree := 0
    prevDegreeEntry := 1
    nextDegreeEntry := 0

    for !found{
        
        numScraping := min(prevDegreeEntry, limit)
        currentSet := []string{}

		for i := 0; i < numScraping; i++ {
			currentSet = append(currentSet, queue[i])
		}
        queue = queue[numScraping:]

		wikis:= scrap(currentSet)

        prevLen := len(parent)
        
        for _, wiki := range wikis {
            if wiki.Title == target {
                parent[wiki.Title] = wiki.Parent
                found = true
                // fmt.Printf("%d.  %s \033[32m→\033[0m %s\n", len(parent), wiki.Parent, wiki.Title)
                break
            }

            if _, ok := parent[wiki.Title]; !ok {
                parent[wiki.Title] = wiki.Parent
                queue = append(queue, wiki.Title)
                // fmt.Printf("%d.  %s \033[32m→\033[0m %s\n", len(parent), wiki.Parent, wiki.Title)
            }
        }

        newEntry := len(parent) - prevLen
        nextDegreeEntry += newEntry

        prevDegreeEntry -= numScraping
        if prevDegreeEntry == 0 {
            prevDegreeEntry = nextDegreeEntry
            nextDegreeEntry = 0
            currDegree++
        }

        fmt.Printf("\033[32mPrevE    :\033[0m %d\n", prevDegreeEntry)
        fmt.Printf("\033[32mDeg      :\033[0m %d\n", currDegree)
		fmt.Printf("\033[32mRead     :\033[0m %d\n", len(parent))
		fmt.Printf("\033[32mIn Queue :\033[0m %d\n", len(queue))
        fmt.Printf("\033[32mBacth no :\033[0m %d\n", ((len(parent)-len(queue)+1)/limit)+1)
        fmt.Printf("\033[32mTime     : \033[0m")
        fmt.Println(getDuration(time.Since(startTime)))
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
            path = append(path, Wiki{strings.ReplaceAll(title, "_", " "), getURL(title)})
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

        numCPUs := runtime.NumCPU()
        fmt.Println("Number of CPUs: ", numCPUs)
        runtime.GOMAXPROCS(numCPUs)

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
