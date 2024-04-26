package bfs

import (
	"fmt"
	"strings"
	"time"

	"backend/func/utils"
)


func BREADTH_FIRST_SEARCH(source, target string) ([]utils.Wiki, int, bool) {

    start_time := time.Now()

    // init map parent
    parent := make(map[string]string)

    // init queue untuk bfs
    queue := []string{source}

    // counter
    curr_degree := 0
    prev_degree_count := 1
    next_degree_count := 0
    
    found := false

    for !found && len(queue) > 0 {
        
        num_scrap := min(prev_degree_count, utils.NUM_PARALLELISM) // menentukan jumlah scrap untuk parallelism scrapping
        curr_batch := []string{} // batch untuk scrap

		for i := 0; i < num_scrap; i++ {
			curr_batch = append(curr_batch, queue[i])
		}

        queue = queue[num_scrap:]
		wikis, err := utils.Scrap(curr_batch)

        if !err {
            prev_parent_len := len(parent)
            
            // cari target di wikis
            for _, wiki := range wikis {

                // jika wiki adalah target, break
                if wiki.Title == target {
                    found = true
                    parent[wiki.Title] = wiki.Parent
                    break
                }
                
                // jika wiki belum pernah dikunjungi, tambahkan ke parent dan queue
                if _, exists := parent[wiki.Title]; !exists {
                    parent[wiki.Title] = wiki.Parent
                    queue = append(queue, wiki.Title)
                }
            }
            
            // update counter
            new_links := len(parent) - prev_parent_len
            next_degree_count += new_links
            prev_degree_count -= num_scrap

            if prev_degree_count == 0 {
                curr_degree++
                prev_degree_count = next_degree_count
                next_degree_count = 0
            }

            // print status
            fmt.Printf("\033[32mPrevE    :\033[0m %d\n", prev_degree_count)
            fmt.Printf("\033[32mDeg      :\033[0m %d\n", curr_degree)
            fmt.Printf("\033[32mRead     :\033[0m %d\n", len(parent))
            fmt.Printf("\033[32mIn Queue :\033[0m %d\n", len(queue))
            fmt.Printf("\033[32mBacth no :\033[0m %d\n", ((len(parent)-len(queue)+1)/utils.NUM_PARALLELISM)+1)
            fmt.Printf("\033[32mTime     : \033[0m")
            fmt.Println(utils.FormatDuration(time.Since(start_time)))
        }
    }

    if !found {
        return nil, 0, true

    } else {
        path_title := []string{target}
        current := target

        // backtracking cari parent pakai map
        for current != source {
            current = parent[current]
            path_title = append([]string{current}, path_title...)
        }

        // constuct Wiki struct
        var path []utils.Wiki
        for _, title := range path_title {
            path = append(path, utils.Wiki{Title: strings.ReplaceAll(title, "_", " "), 
                                           URL  : utils.ConstructURL(title)})
        }

        return path, len(parent), false
    }
}


func EntryPoint(source_wiki, target_wiki string) ([]utils.Wiki, utils.Duration, int) {

    source_wiki = strings.ReplaceAll(source_wiki, " ", "_")
    target_wiki = strings.ReplaceAll(target_wiki, " ", "_")

    if (!(utils.IsValidWiki(source_wiki) && utils.IsValidWiki(target_wiki))) {
        return []utils.Wiki{}, utils.Duration{}, 0

    } else {
		start_time := time.Now()
		path, searched, err := BREADTH_FIRST_SEARCH(source_wiki, target_wiki)
		duration := utils.FormatDuration(time.Since(start_time))

		if err {
			return []utils.Wiki{}, utils.FormatDuration(0), 0
		}

		fmt.Println("\033[93mPath:\033[0m")
		for _, wiki := range path {
			fmt.Println(wiki)
		}
		fmt.Println("\033[93mSearched: \033[0m", searched)
		fmt.Println("\033[93mDuration: \033[0m", duration)

		return path, duration, searched
	} 
}
