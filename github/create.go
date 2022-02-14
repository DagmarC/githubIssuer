package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// POST /repos/{owner}/{repo}/issues

/*
curl \
  -X POST \
  -H "Accept: application/vnd.github.v3+json" \
  https://api.github.com/repos/octocat/hello-world/issues \
  -d '{"title":"title"}'
*/

func CreateRepoIssues(repo, title, body, token string) (*Issue, error) {
	q := fmt.Sprintf("/repos/%v/issues", repo)

	issueReq := &IssueReq{
		Title: title,
		Body:  body,
	}

	buf, err := encodeIssueRequest(issueReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseGithubURL+q, buf)
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
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to edit issue: %s", resp.Status)
	}

	defer resp.Body.Close()

	var issue Issue
	if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func encodeIssueRequest(req *IssueReq) (io.ReadWriter, error) {
	var buf io.ReadWriter
	if req != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(req)
		if err != nil {
			return nil, err
		}
	}
	return buf, nil
}
