package main

import (
	"net/http"
	"fmt"
	"time"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Nothing to download")
		return
	}

	start := time.Now()
	ch := make(chan string)
	contentCh := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchAll(url, ch, contentCh)
	}

	pages := make([]string, len(os.Args)-1)
	for i := range os.Args[1:] {
		fmt.Println(<-ch)
		pages[i] = <-contentCh
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	outputDir := "C:\\Users\\Gautham\\Projects\\Github\\GoStash\\TGPL\\chapter1\\outputs\\1.10\\"
	writeAll(outputDir, pages)
}

func writeAll(outputDir string, pages []string) {
	for i, page := range pages {
		ioutil.WriteFile(outputDir+strconv.Itoa(i)+".txt", []byte(page), 0644)
	}
}

func fetchAll(url string, ch chan<- string, contentCh chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("while fetching %s : %v\n", url, err)
		return
	}

	// Copy to ioutil.Discard to redirect to /dev/null.
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s : %v\n", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, len(content), url)
	contentCh <- string(content)
}
