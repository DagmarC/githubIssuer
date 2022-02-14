package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PATCH /repos/{owner}/{repo}/issues/{issue_number}

func UpdateRepoIssues(repo, title, body, token string, number int) (*Issue, error) {
	q := fmt.Sprintf("/repos/%v/issues/%d", repo, number)

	issueReq := &IssueReq{
		Title: title,
		Body:  body,
	}

	buf, err := encodeIssueRequest(issueReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", baseGithubURL+q, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token " + token)
	req.Header.Set("Accept", mediaTypeV3)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to edit issue: %s", resp.Status)
	}

	defer resp.Body.Close()

	var issue Issue
	if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}

	return &issue, nil
}
