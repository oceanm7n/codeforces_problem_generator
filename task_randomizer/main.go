package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func scrape_data() []string {

	c := colly.NewCollector(
		colly.AllowedDomains(
			"https://codeforces.com",
			"https://codeforces.com/problemset",
			"http://codeforces.com/problemset",
			"codeforces.com/problemset",
			"http://codeforces.com/",
			"codeforces.com",
		),
	)

	links := make([]string, 0)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
		return
	})

	c.OnResponse(func(r *colly.Response) {

	})

	c.OnHTML(".id", func(e *colly.HTMLElement) {
		l := e.ChildAttrs("a", "href")
		links = append(links, l[0])
	})

	page := 1
	for page <= 79 {
		link := fmt.Sprint("https://codeforces.com/problemset/page/", page)
		c.Visit(link)
		page += 1
	}
	save_to_csv(links)
	return links
}

func save_to_csv(data []string) (fName string) {

	fName = "problem_set.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, link := range data {
		l := []string{link}
		writer.Write(l)
	}

	fmt.Println("Successfully saved scraped data to", fName)
	return fName
}

func stringInSlice(a byte, list []string) bool {
	for _, b := range list {
		if b[0] == a {
			return true
		}
	}
	return false
}

func get_random_problem(data []string, c string) string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(data))
	guess := data[r]
	if c == "any" {
		return "https://codeforces.com" + guess
	}
	if stringInSlice(c[0], []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "a", "b", "c", "d", "e", "f", "g", "h", "i"}) {
		if c[0] == guess[len(guess)-1] {
			return "https://codeforces.com" + guess
		} else {
			return get_random_problem(data, c)
		}
	} else {
		return "Please try fixing the task complexity argument when executing a binary."
	}
}

func get_data_from_file(filename string, dir string) (data []string) {

	f, err := os.Open(dir + "\\" + filename)
	if err != nil {
		log.Printf("Cannot open file %q: %s\n", filename, err)
		log.Fatalf("Please use --scrape argument to scrape the problem set first.")
	}
	defer f.Close()

	csv_reader := csv.NewReader(f)
	lines, err := csv_reader.ReadAll()
	if err != nil {
		log.Fatalf("Cannot read file %q: %s\n", filename, err)
	}

	data = make([]string, 0)

	for _, line := range lines {
		data = append(data, line[0])
	}

	return data
}

func get_run_directory() string {
	pwd, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	temp_dir := strings.Split(pwd, "\\")
	temp_dir = temp_dir[0 : len(temp_dir)-1]
	dir := strings.Join(temp_dir, "\\")
	return dir
}

func main() {

	dir := get_run_directory()
	args := ReadArgs()

	ARG_SCRAPE := args.GetArg("--scrape")
	ARG_FILE := args.GetArg("-p")
	ARG_COMPLEXITY := args.GetArg("-c")

	var data []string

	if ARG_SCRAPE.exists {
		data = scrape_data()
	} else {
		if ARG_FILE.exists {
			data = get_data_from_file(ARG_FILE.value, dir)
		} else {
			ARG_FILE.value = "problem_set.csv"
			data = get_data_from_file(ARG_FILE.value, dir)
		}
	}

	var link string

	if ARG_COMPLEXITY.exists {
		link = get_random_problem(data, ARG_COMPLEXITY.value)
	} else {
		link = get_random_problem(data, "any")
	}

	fmt.Println(link)

}
