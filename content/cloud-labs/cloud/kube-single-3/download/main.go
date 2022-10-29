package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

func download(fullURLFile string) {

	// Build fileName from fullPath
	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d\n", fileName, size)

}

func main() {
	readFile, err := os.Open("../index.md")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	wg := sync.WaitGroup{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(line, "![](https://i.imgur.com/") {
			line = strings.Replace(line, "![](", "", -1)
			line = strings.Replace(line, ")", "", -1)
			fmt.Println(line)
			wg.Add(1)
			go func(line string) {
				defer wg.Done()
				download(line)
			}(line)
		}
	}
	wg.Wait()
	readFile.Close()
}
