# Go-Scraper
> A simple web scraper written in Go, just to try the language.

The goal of this project is to learn and practice Go, also known as Golang, a strong and versatile language
fresh from some of googles finest, Robert Griesemer, Rob Pike, and Ken Thompson - According to [Wikipedia](https://en.wikipedia.org/wiki/Go_(programming_language).

Go's version of threading (go-routines) interest me, and the garbage collection seems impressive. So I want try and improve my grasp of Go
by building a web scraper with the following specifications:

1. Allow the user to input a URL to scrape.
2. Fetch the HTML content of the specified URL.
3. Parse the HTML to extract relevant information (e.g., titles, links, images).
4. Display the extracted information in the console.

This will all be via command line, but my hope is to make something more flexible.

For web-scraping, I am going to use Colly, and open source tool meant for scrapers,
crawlers, and spiders. It seems simple to use, and has very good documentation.
