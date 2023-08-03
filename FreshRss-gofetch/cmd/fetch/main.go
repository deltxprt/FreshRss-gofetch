package main

import (
	"flag"
	"github.com/go-git/go-git/v5"
	"log"
	"os"
	"strings"
)

func readUrls(url string) ([]string, error) {
	var urls []string
	if strings.Contains(url, ";") {
		urls = strings.Split(url, ";")
	} else {
		urls = []string{url}
	}

	return urls, nil
}

func downloadExtensions(destPath string, url string) error {
	// Clones the repository into the worktree (fs) and stores all the .git
	_, err := git.PlainClone(destPath, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return err
	}
	files, err := os.ReadDir(destPath)

	if err != nil {
		return err
	}
	for _, file := range files {
		log.Println(file.Name())

		if !file.IsDir() {
			err = os.Remove(destPath + file.Name())
			if err != nil {
				return err
			}
		} else if file.Name() == ".git" {
			err = os.RemoveAll(destPath + file.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	var urls []string
	path := flag.String("path", "/extensions/", "path to the extensions")
	url := flag.String("url", "", "urls to the zip file of the extensions")
	flag.Parse()

	if *url == "" {
		log.Fatal("no url specified")
		return
	}

	urls, err := readUrls(*url)

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, url := range urls {
		err := downloadExtensions(*path, url)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
