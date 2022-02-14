package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
curl \
  -H "Accept: application/vnd.github.v3+json" \
  https://api.github.com/issues

  GET /repos/{owner}/{repo}/issues

*/

func GetRepoIssues(repo string) (*[]*Issue, error) {

	q := "/repos/" + repo + "/issues"
	fmt.Println(baseGithubURL + q)

	req, err := http.NewRequest("GET", baseGithubURL+q, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", mediaTypeV3)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result []*Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
