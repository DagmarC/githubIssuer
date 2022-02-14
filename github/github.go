// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 110.
//!+

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import "time"

const issuesURL = "https://api.github.com/search/issues"
const baseGithubURL = "https://api.github.com"
const mediaTypeV3 = "application/vnd.github.v3+json"
const defaultUser = "DagmarC"
const defaultUserURL = "https://github.com/DagmarC"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    // in Markdown format
	Assignee  *User
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// IssueRequest represents a request to create/edit an issue.
// It is separate from Issue above because otherwise Labels
// and Assignee fail to serialize to the correct JSON.
type IssueReq struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Labels    *[]string `json:"labels,omitempty"`
	State     string   `json:"state,omitempty"`
	Milestone *int      `json:"milestone,omitempty"`
	Assignees *[]string `json:"assignees,omitempty"`
}

//!-
