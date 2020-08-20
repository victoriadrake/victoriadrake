package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
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
	blogItem := feed.Items[0]

	wc, err := fp.ParseURL("https://victoria.dev/wc/index.xml")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}
	// Add this much magic
	wcItem := wc.Items[0]

	rand.Seed(time.Now().UnixNano())
	whoList := []string{"🦄", "🐣", "🦊", "🦔", "🦡", "🐹", "🐷", "🌮", "🍝", "🥑", "💩"}
	who := rand.Intn(len(whoList))
	whatList := []string{"👍", "🎉", "💕", "🤷", "👏", "🙌"}
	what := rand.Intn(len(whatList))
	date := time.Now().Format("2 Jan 2006")

	// Whisk together static and dynamic content until stiff peaks form
	hello := "### Hello! I’m Victoria Drake. 👋\n\nI’m a software developer at 💜 and Director of Engineering at work. I build my skill stack in public and share open source knowledge through the " + wcItem.Description + " words I’ve written on [victoria.dev](https://victoria.dev). I hope to encourage people to learn openly and fearlessly, with wild child-like abandon.\n\nI spent summers as a kid hacking Neopets and coding a medieval multiplayer role play world through a teeny tiny terminal window. Sometime before that, I wrote mystery novels and lived a secret life as a super sleuth.\n\n- See my :octocat: **[GitHub Action cybersecurity tools](https://github.com/search?q=user%3Avictoriadrake+GitHub+Action+security)**"
	blog := "- This " + whoList[who] + " says they " + whatList[what] + " my latest blog post: **[" + blogItem.Title + "](" + blogItem.Link + ")**"
	updated := "Last updated by magic on " + date + "."
	data := fmt.Sprintf("%s\n%s\n\n%s\n", hello, blog, updated)

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
