/*

A command line client for the Giphy API

Installation

Just go get the command:

		go get -u github.com/peterhellberg/giphy/cmd/giphy

Configuration

The command line client can be used straight out of the box, but
there are also a few environment variables that you can use in order
to override the default configuration.

    Environment variable  | Default value
    ----------------------|--------------
    GIPHY_API_KEY         | dc6zaTOxFJmzC
    GIPHY_RATING          | g
    GIPHY_LIMIT           | 10
    GIPHY_BASE_URL_SCHEME | http
    GIPHY_BASE_URL_HOST   | api.giphy.com
    GIPHY_BASE_PATH       | /v1
    GIPHY_USER_AGENT      | giphy.go

Usage

The command line client consists of a few sub commands.

    Commands:
      search, s           [args]
      gif, id             [args]
      random, rand, r     [args]
      translate, trans, t [args]
      trending, trend, tr [args]

*/
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/peterhellberg/giphy"
	"flag"
	"os/exec"
	"runtime"
)

var command = flag.String("command", "", strings.Join([]string{
	"Commands:",
	"search, s           [args]",
	"gif, id             [args]",
	"random, rand, r     [args]",
	"translate, trans, t [args]",
	"trending, trend, tr [args]",
}, "\n\t"))

var render = flag.Bool("render", false, "Boolean flag to render the GIF in browser or not")


func main() {
	g := giphy.DefaultClient

	flag.Parse()

	switch *command {
	default:
		flag.PrintDefaults()
	case "search", "s":
		search(g, flag.Args())
	case "gif", "id":
		gif(g, flag.Args())
	case "random", "rand", "r":
		random(g, flag.Args())
	case "translate", "trans", "t":
		translate(g, flag.Args())
	case "trending", "trend", "tr":
		trending(g, flag.Args())
	}
}

func search(c *giphy.Client, args []string) {
	res, err := c.Search(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, d := range res.Data {
		fmt.Println(d.Images.Original.URL)

		renderInBrowser(d.Images.Original.URL)
	}
}

func renderInBrowser(URL string) {

	if *render {
		switch runtime.GOOS {
		case "darwin":
			_, err := exec.LookPath("open")

			if err == nil {
				exec.Command("open", URL).Output()
			} else {
				fmt.Println("Couldnt open in browser, missing package")
			}
		case "windows":
			_, err := exec.LookPath("start")

			if err == nil {
				exec.Command("start", URL).Output()
			} else {
				fmt.Println("Couldnt open in browser, missing package")
			}
		case "linux":
			_, err := exec.LookPath("xdg-open")

			if err == nil {
				exec.Command("xdg-open", URL).Output()
			} else {
				fmt.Println("Couldnt open in browser, missing package")
			}
		default:
			fmt.Println("Couldnt open in browser, missing package")
		}
	}

}

func gif(c *giphy.Client, args []string) {
	if len(args) == 0 {
		fmt.Println("missing Giphy id")
		os.Exit(1)
	}

	res, err := c.GIF(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(res.Data.Images.Original.URL)

	renderInBrowser(res.Data.Images.Original.URL)
}

func random(c *giphy.Client, args []string) {
	res, err := c.Random(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(res.Data.ImageOriginalURL)

	renderInBrowser(res.Data.ImageOriginalURL)
}

func translate(c *giphy.Client, args []string) {
	res, err := c.Translate(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(res.Data.Images.Original.URL)

	renderInBrowser(res.Data.Images.Original.URL)
}

func trending(c *giphy.Client, args []string) {
	res, err := c.Trending(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, d := range res.Data {
		fmt.Println(d.Images.Original.URL)

		renderInBrowser(d.Images.Original.URL)
	}
}
