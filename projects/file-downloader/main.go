package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// downloader method

func DownloaderMTD(url, destDir string) error {
	fileName := filepath.Base(url)
	filepath := filepath.Join(destDir, fileName)

	out, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer out.Close()

	fmt.Println("Downloading .....")
	fmt.Println("")

	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		os.Remove(filepath)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		os.Remove(filepath)
		return fmt.Errorf("Bad status %s", res.Status)
	}

	if !(filepath[len(filepath)-4:] == ".jpg" || filepath[len(filepath)-4:] == ".png" || filepath[len(filepath)-5:] == ".jpeg") {
		os.Remove(filepath)
		return fmt.Errorf("This file is not supported for download")
	}

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Downloading of %s took %s amount of time.\n", fileName, time.Since(start))
	return nil

}

func SequentialDownloader(urls []string, destDir string) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	start := time.Now()
	for _, url := range urls {
		if err := DownloaderMTD(url, destDir); err != nil {
			fmt.Println("Error downloading", err)
			continue
		}
	}
	fmt.Printf("Total download time took %d for this urls : %s\n\n ", time.Since(start).Milliseconds(), urls)

	return nil
}

type Result struct {
	Url      string
	FileName string
	Size     int64
	Duration time.Duration
	Error    error
}

func ConcurrentDownloader(urls []string, destDir string, maxConcurrent int) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	result := make(chan Result)

	var wg sync.WaitGroup

	limiter := make(chan struct{}, maxConcurrent)

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			limiter <- struct{}{}
			defer func() {
				<-limiter
			}()

			start := time.Now()
			filename := filepath.Base(url)
			fmt.Printf("file name printed here %s\n", filename)
			filepath := filepath.Join(destDir, filename)

			out, err := os.Create(filepath)
			if err != nil {
				result <- Result{Url: url, Error: err}
				return
			}
			defer out.Close()

			res, err := http.Get(url)
			if err != nil {
				result <- Result{Url: url, Error: err}
				return
			}
			defer res.Body.Close()
			if res.StatusCode != http.StatusOK {
				result <- Result{Url: url, Error: fmt.Errorf("Bad status %s", res.Status)}
				return
			}

			if !(filepath[len(filepath)-4:] == ".jpg" || filepath[len(filepath)-4:] == ".png" || filepath[len(filepath)-5:] == ".jpeg") {
				result <- Result{Url: url, Error: fmt.Errorf("This file is not supported for download")}
				return
			}

			n, err := io.Copy(out, res.Body)
			if err != nil {
				result <- Result{Url: url, Error: err}
				return
			}

			timeSince := time.Since(start)

			result <- Result{
				Url:      url,
				FileName: filename,
				Size:     n,
				Duration: timeSince,
				Error:    nil,
			}

		}(url)
	}

	go func() {
		wg.Wait()
		close(result)
	}()
	var totalSize int64
	var errors []error

	start := time.Now()

	for res := range result {
		if res.Error != nil {
			fmt.Printf("Error downloading... %s : %s\n", res.Url, res.Error.Error())
			errors = append(errors, res.Error)
		} else {
			totalSize += res.Size
			fmt.Printf("Download this file :%s of this file size %d at this duration : %s\n ", res.FileName, res.Size, res.Duration)
		}
	}
	startedSince := time.Since(start)
	fmt.Printf("Total downloaded size: %d in %s\n", totalSize, startedSince)

	if len(errors) > 0 {
		return fmt.Errorf("Some downloads failed")
	}

	return nil
}

func main() {

	var wg sync.WaitGroup

	urls := []string{
		"https://www.gstatic.com/webp/gallery3/2.png",
		"https://www.gstatic.com/webp/gallery.png",
		"https://www.gstatic.com/webp/gallery3/1.png",
		"https://www.gstatic.com/webp/gallery3/3.png",
		"https://www.gstatic.com/webp/galleryß.png",
		"https://www.gstatic.com/webp/gallery3/4.png",
		"https://www.gstatic.com/webp/gallery3/5.png",
	}
	// run the concurrent downloader and the sequential downloader at the same time to compare the time taken by both methods using go routines
	wg.Add(1)

	go func() {
		defer wg.Done()
		er := SequentialDownloader(urls, "./sequential")
		if er != nil {
			fmt.Println("error downloading the file", er)
			return
		}
		log.Println("Sequential Download Done!!")
	}()

	wg.Add(2)
	go func() {
		defer wg.Done()
		er := ConcurrentDownloader(urls, "./concurrent", len(urls))
		if er != nil {
			fmt.Println("eror downloading the file", er)
			return
		}
		log.Println("Concurrent Download Done!!")
	}()

	wg.Wait()

	log.Println("Done!!")
}
