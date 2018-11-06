package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andygrunwald/go-jira"
	"github.com/fatih/color"
)

var (
	// myIssues  = flag.Bool("my", false, "Show my open issues, sorted by time_created")
	// listTodos = flag.String("list", "", "List all todos from current dir")
	cfg = readConfig()
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

func listMyOpenIssues(client *jira.Client) {
	issues := getJiraIssues("assignee = currentUser() AND resolution = Unresolved AND status != Review ORDER BY created DESC")
	for _, issue := range issues {
		green := color.New(color.FgGreen)
		blue := color.New(color.FgBlue)
		green.Printf("%v/browse/%v ", cfg.URL, issue.Key)
		fmt.Print("[")
		blue.Printf("%v", issue.Fields.Status.Name)
		fmt.Print("] ")
		fmt.Printf("%v\n", issue.Fields.Summary)
	}
}

func usage() {
	// TODO(#9): implement a map for options instead of println'ing them all there
	fmt.Printf("snitch [opt]\n" +
		"\tlist: lists all todos of a dir recursively\n" +
		"\treport <owner/repo> [issue-body]: reports all todos of a dir recursively as GitHub issues\n" +
		"\tpurge <owner/repo>: removes all of the reported TODOs that refer to closed issues\n")
}

func getJiraIssues(request string) []jira.Issue {
	client := getJiraAPI()
	issues, _, err := client.Issue.Search(request, nil)
	if err != nil {
		panic(err)
	}
	return issues
}

func getJiraAPI() *jira.Client {

	tp := jira.BasicAuthTransport{
		Username: cfg.Login,
		Password: cfg.Password,
	}

	client, err := jira.NewClient(tp.Client(), cfg.URL)
	if err != nil {
		panic(err)
	}
	return client
}

func main() {

	if os.Args < 1 {



	switch os.Args[1]{
		case :
	}}


}
