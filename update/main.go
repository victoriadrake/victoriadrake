package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
)

func makeReadme(filename string) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://victoria.dev/index.xml")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}
	// Get the freshest item
	rssItem := feed.Items[0]

	// Unwrap Markdown content
	content, err := ioutil.ReadFile("static.md")
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
		return err
	}
	stringyContent := string(content)

	date := time.Now().Format("2 Jan 2006")

	// Whisk together static and dynamic content until stiff peaks form
	blog := "- ✨ Read my latest blog post: **[" + rssItem.Title + "](" + rssItem.Link + ")**"
	updated := "Last updated by magic on " + date + "."
	data := fmt.Sprintf("%s%s\n\n%s\n", stringyContent, blog, updated)

	// Prepare file with a light coating of os
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Bake at n bytes per second until golden brown
	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func main() {

	makeReadme("../README.md")

}
