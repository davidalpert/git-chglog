package chglog

import (
	"fmt"
	agjira "github.com/andygrunwald/go-jira"
)

// JiraClient is an HTTP client for Jira
type JiraClient interface {
	GetJiraIssue(id string) (*agjira.Issue, error)
	Validate() error
}

type jiraClient struct {
	username string
	token    string
	url      string
}

// NewJiraClient returns an instance of JiraClient
func NewJiraClient(config *Config) JiraClient {
	return jiraClient{
		username: config.Options.JiraUsername,
		token:    config.Options.JiraToken,
		url:      config.Options.JiraURL,
	}
}

func (jira jiraClient) Validate() error {
	if jira.username == "" {
		return fmt.Errorf("username cannot be nil")
	}

	if jira.token == "" {
		return fmt.Errorf("token cannot be nil")
	}

	if jira.url == "" {
		return fmt.Errorf("url cannot be nil")
	}

	return nil
}

func (jira jiraClient) GetJiraIssue(id string) (*agjira.Issue, error) {
	tp := agjira.BasicAuthTransport{
		Username: jira.username,
		Password: jira.token,
	}
	client, err := agjira.NewClient(tp.Client(), jira.url)
	if err != nil {
		return nil, err
	}
	issue, _, err := client.Issue.Get(id, nil)
	return issue, err
}
