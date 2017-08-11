package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Nothing to download")
		return
	}

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy: %v\n", err)
			os.Exit(1)
		}

		resp.Body.Close()
	}
}
