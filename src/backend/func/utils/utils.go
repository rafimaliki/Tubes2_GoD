package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

var (
    NUM_PARALLELISM int = 100
    NUM_CPU int = runtime.NumCPU()
    mu sync.Mutex
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

func FormatDuration(duration time.Duration) string {
    hours := int(duration.Hours())
    minutes := int(duration.Minutes()) % 60
    seconds := int(duration.Seconds()) % 60
    milliseconds := int(duration.Milliseconds()) % 1000

    durationStr := ""

    if hours > 0 {
        durationStr += fmt.Sprintf("%dh ", hours)
    }
    if minutes > 0 {
        durationStr += fmt.Sprintf("%dm ", minutes)
    }
    if seconds > 0 {
        durationStr += fmt.Sprintf("%ds ", seconds)
    }
    if milliseconds > 0 {
        durationStr += fmt.Sprintf("%dms", milliseconds)
    }
    durationStr = strings.TrimSpace(durationStr)

    return durationStr
}

func IsValidWiki(title string) bool {

	url := "https://en.wikipedia.org/wiki/" + title
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func ConstructURL(title string) string {
	return "https://en.wikipedia.org/wiki/" + title
}

func Scrap(urls []string) ([]WikiTitle, bool) {

    var wikis []WikiTitle

    // print urls
    fmt.Println("Scraping", len(urls), urls)

    // setting colly collector
    c := colly.NewCollector(
        colly.AllowedDomains("en.wikipedia.org"), colly.Async(true), colly.CacheDir("./cache"),
    )
    c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: NUM_PARALLELISM})
    
    // init queue
    q, err := queue.New(NUM_CPU,
        &queue.InMemoryQueueStorage{MaxSize: NUM_PARALLELISM},
    )

    // return jika gagal init queue
    if err != nil {
        return nil, true
    }

    // isi queue dengan url
    for _, url := range urls {
        q.AddURL("https://en.wikipedia.org/wiki/" + url)
    }

    // callback onHTML
    c.OnHTML("a[href^='/wiki/']:not([href*=':']):not([href='/wiki/Main_Page'])", func(e *colly.HTMLElement) {
        link,_ := url.PathUnescape(strings.TrimPrefix(e.Attr("href"), "/wiki/"))
        if idx := strings.Index(link, "#"); idx != -1 {
            link = link[:idx]
        }
        parent, _ := url.PathUnescape(strings.TrimPrefix(e.Request.URL.String(), "https://en.wikipedia.org/wiki/"))

        // lock
        mu.Lock()
        defer mu.Unlock()
        
        wikis = append(wikis, WikiTitle{link, parent})
    })

    // run gocollyqueue
    err = q.Run(c)

    // return jika gagal run gocollyqueue
    if err != nil {
        return nil, true
    }

    // tunggu semua request selesai
    c.Wait()

    return wikis, false
}