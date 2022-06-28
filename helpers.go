package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/toqueteos/webbrowser"
)

func Get_run_directory() string {
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

func StringInSlice(a byte, list []string) bool {
	for _, b := range list {
		if b[0] == a {
			return true
		}
	}
	return false
}

func Print_result(link string, open bool) {
	fmt.Println("Random Codeforces problem selected, link:")
	fmt.Println(link)
	if !open {
		webbrowser.Open(link)
	}
}

func GetHelp() {
	help := `
This command line tool scrapes the list of tasks from Codeforces website (http://www.codeforces.com/). 

Usage: 
	-c  	 	Select task complexity. Ignoring this parameter will result in randomly selected task complexity-wise
	--scrape  	Scrapes data and saves to ./data folder as .csv file. Be sure you have a .csv file unless 
	-p 			Specify a binary executable location (without ./data)
	-d 			Disable browser redirecting to a randomed page

Examples:
	task_randomizer.exe --help
	Get help

	task_randomizer.exe --scrape 
	Must be run first time to get data

	task_randomizer.exe -c B
	Select a random %task%B problem and open in browser

	task_randomizer.exe -c B -d
	Select a random %task%B problem WITHOUT opening in browser

	`
	fmt.Println(help)
	os.Exit(1)
}
