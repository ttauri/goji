package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/andygrunwald/go-jira"
	"github.com/fatih/color"
)

var (
	myIssues = flag.Bool("my", false, "Show my open issues, sorted by time_created")
)

// Config from file
type Config struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	URL      string `json:"jira_url"`
}

func readConfig() *Config {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	conf := new(Config)
	_ = json.Unmarshal(bytes, &conf)
	return conf
}

func main() {
	flag.Parse()
	cfg := readConfig()
	fmt.Println(cfg)
	tp := jira.BasicAuthTransport{
		Username: cfg.Login,
		Password: cfg.Password,
	}

	client, err := jira.NewClient(tp.Client(), cfg.URL)
	if err != nil {
		panic(err)
	}

	// Print My open issues
	if *myIssues == true {
		issues, _, err := client.Issue.Search("assignee = currentUser() AND resolution = Unresolved AND status != Review ORDER BY created DESC", nil)
		if err != nil {
			panic(err)
		}
		for _, issue := range issues {
			green := color.New(color.FgGreen)
			blue := color.New(color.FgBlue)
			green.Printf("https://pmc.acronis.com/browse/%v ", issue.Key)
			fmt.Print("[")
			blue.Printf("%v", issue.Fields.Status.Name)
			fmt.Print("] ")
			fmt.Printf("%v\n", issue.Fields.Summary)
		}
	}
}
