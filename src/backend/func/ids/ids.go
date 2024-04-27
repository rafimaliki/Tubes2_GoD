package ids

import (
	"backend/func/utils"
	"fmt"
	"net/url"
	"runtime"
	"strings"
	"time"
)

var (
	NUM_PARALLELISM int = 100
	NUM_CPU         int = runtime.NumCPU()
)

// func main() {
// 	path, elapsed, foundLink := EntryPointIDS("Hololive", "Shark")

// 	if len(path) > 0 {
// 		fmt.Println("Path found:")
// 		printPath(path)
// 	} else {
// 		fmt.Println("Path not found within the depth limit")
// 	}

// 	fmt.Println(elapsed)
// 	fmt.Println(foundLink)
// }

func EntryPointIDS(source_wiki, target_wiki string) ([]utils.Wiki, string, int) {
	var foundLinksCount int
	start := time.Now()
	targetURL := getWikiURL(target_wiki)
	foundCount := 0
	found := false
	depth := 0
	var err bool
	var path []utils.Wiki
	source := utils.Wiki{Title: source_wiki, URL: getWikiURL(source_wiki)}

	if !utils.IsValidWiki(source_wiki) {
		return []utils.Wiki{}, utils.FormatDuration(0), -1

	} else if !utils.IsValidWiki(target_wiki) {
		return []utils.Wiki{}, utils.FormatDuration(0), -2
	} else {

		for !found {
			// Tambahkan source ke awal path
			path = append([]utils.Wiki{source}, path...)

			found, path, foundCount, err = IDS(source_wiki, targetURL, depth, path)
			if err != false {
				return path, utils.FormatDuration(0), -3
			}
			fmt.Println("Depth: ", depth)

			depth++

			foundLinksCount = foundCount

			if depth > 6 {
				break
			}
		}

		elapsed := utils.FormatDuration(time.Since(start))
		return path, elapsed, foundLinksCount
	}

}

func IDS(sourceTitle string, targetURL string, depth int, path []utils.Wiki) (bool, []utils.Wiki, int, bool) {
	printPath(path)

	if getWikiURL(sourceTitle) == targetURL {
		return true, append(path, utils.Wiki{Title: sourceTitle, URL: "https://en.wikipedia.org/wiki/" + sourceTitle}), 0, false
	}

	if depth == 0 {
		return false, nil, 0, false
	}

	links, err := utils.Scrap([]string{sourceTitle}) // Mengubah pemanggilan fungsi Scrap dengan menyediakan slice string dari judul artikel
	foundCount := len(links)

	if err != false {
		return false, nil, 0, true
	}

	for _, tempLink := range links {
		// fmt.Println(getURL(tempLink.Title), " = ?", targetURL)
		if getWikiURL(tempLink.Title) == targetURL { // Memeriksa apakah judul artikel sama dengan URL target
			return true, append(path, utils.Wiki{Title: tempLink.Title, URL: "https://en.wikipedia.org/wiki/" + tempLink.Title}), foundCount, false
		}

		found, newPath, count, err := IDS(tempLink.Title, targetURL, depth-1, append(path, utils.Wiki{Title: tempLink.Title, URL: "https://en.wikipedia.org/wiki/" + tempLink.Title})) // Memperbaiki pemanggilan rekursif untuk menambahkan objek Wiki ke dalam path
		if found {
			foundCount += count
			return true, newPath, foundCount, err
		}
	}

	return false, nil, foundCount, err
}

func printPath(path []utils.Wiki) {
	for i, title := range path {
		if i > 0 {
			fmt.Print(" -> ")
		}
		fmt.Print(title.Title)
	}
	fmt.Println()
}

func getWikiURL(title string) string {
	url := "https://en.wikipedia.org/wiki/" + url.PathEscape(strings.ReplaceAll(title, " ", "_"))
	return url
}
