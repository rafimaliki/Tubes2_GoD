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
	Title  string
	URL    string
	Parent *Wiki
}

// Struktur untuk merepresentasikan queue
type Queue struct {
	items []Wiki
}

// Fungsi untuk menambahkan elemen ke dalam queue
func (q *Queue) Enqueue(item Wiki) {
	q.items = append(q.items, item)
}

// Fungsi untuk menghapus elemen dari depan queue dan mengembalikan nil jika queue kosong
func (q *Queue) Dequeue() {
	if len(q.items) > 0 {
		q.items = q.items[1:]
	}
}

// Fungsi untuk mengambil elemen dari depan queue tanpa menghapusnya dan mengembalikan nil jika queue kosong
func (q *Queue) Front() Wiki {
	if len(q.items) == 0 {
		return Wiki{}
	}
	return q.items[0]
}

// Fungsi untuk mengecek apakah queue kosong
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Fungsi mengembalikan semua link yang ada pada 1 halaman Wikipedia
func getAllLinks(inputWiki Wiki) []Wiki {
	res, err := http.Get(inputWiki.URL)

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
					Parent: &inputWiki,
					Title:  s.Text(),
					URL:    "https://en.wikipedia.org" + href,
				})
			}
		}
	})

	return links
}

func isValidWikiLink(href string) bool {
	// Check if the link starts with "/wiki/" and not contains any ":" which indicates non-article links
	return strings.HasPrefix(href, "/wiki/") && !strings.Contains(href, ":")
}

func getTitleFromLink(link string) (string, error) {
	// Melakukan request HTTP ke link
	res, err := http.Get(link)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Memeriksa status code
	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Membuat dokumen HTML dari body response
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}
	// Mendapatkan judul dari halaman HTML
	title := doc.Find("title").Text()

	temp := strings.TrimSuffix(title, " - Wikipedia")

	return strings.TrimSpace(temp), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Menampilkan pesan untuk meminta input dari pengguna
	fmt.Print("Masukkan judul awal wikipedia: ")

	// Membaca input dari pengguna
	scanner.Scan()

	// Mengambil string dari input yang telah dibaca
	string_awal := scanner.Text()

	// Menampilkan pesan untuk meminta input dari pengguna
	fmt.Print("Masukkan judul akhir wikipedia: ")

	// Membaca input dari pengguna
	scanner.Scan()

	// Mengambil string dari input yang telah dibaca
	string_akhir := scanner.Text()

	// link judul
	link_awal := "https://en.wikipedia.org/wiki/" + strings.ReplaceAll(string_awal, " ", "_")

	// link akhir
	link_akhir := "https://en.wikipedia.org/wiki/" + strings.ReplaceAll(string_akhir, " ", "_")

	// mencari judul dari link awal
	title_awal, err := getTitleFromLink(link_awal)
	if err != nil {
		log.Fatal(err)
	}

	// mencari judul dari link akhir
	title_akhir, err := getTitleFromLink(link_akhir)
	if err != nil {
		log.Fatal(err)
	}

	wiki_awal := Wiki{
		Parent: nil,
		Title:  title_awal,
		URL:    link_awal,
	}

	// Membuat queue baru
	queue := Queue{}

	// Menambahkan judul awal ke dalam queue
	queue.Enqueue(wiki_awal)
	// Print judul awal dan akhir
	fmt.Println("Judul awal: " + title_awal)
	fmt.Println("Judul akhir: " + title_akhir)

	fmt.Println("Mencari jalur dari " + title_awal + " ke " + title_akhir + "...")
	for {
		list_links := getAllLinks(queue.Front())
		for i := 0; i < len(list_links); i++ {
			queue.Enqueue(list_links[i])
		}
		// fmt.Println(queue.Front().Title)
		if queue.Front().Title == title_akhir || queue.Front().URL == link_akhir {
			break
		} else {
			queue.Dequeue()
		}
	}

	result := queue.Front()
	list_result := []Wiki{result} // Inisialisasi slice dengan satu elemen
	for result.Parent != nil {
		result = *result.Parent                              // Mengambil parent berikutnya
		list_result = append([]Wiki{result}, list_result...) // Menambahkan parent ke slice
	}
	println()
	println("Jalur dari " + title_awal + " ke " + title_akhir + " adalah:")
	for i := 0; i < len(list_result); i++ {
		fmt.Println(list_result[i].Title)
		fmt.Println(list_result[i].URL)
	}
}
