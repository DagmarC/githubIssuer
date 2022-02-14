// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/DagmarC/githubIssuer/github"
)

var ri = flag.Bool("ri", false, "Get the repo issues.")
var repo = flag.String("repo", "DagmarC/gopl-solutions", "Get the repo issues.")

var ci = flag.Bool("ci", false, "Create the repo issue.")
var title = flag.String("title", " Title", "Create new repo issue.")
var body = flag.String("body", "Default Body", "Body of the issue.")

var st = flag.Bool("st", false, "Search the terms.")

var token = flag.String("token", "", "github auth token")

//!+
func main() {

	flag.Parse()

	if *ri {
		resultR, err := github.GetRepoIssues(*repo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("------Repository issues------")
		for _, item := range *resultR {
			fmt.Printf("#%-5d %9.9s %.55s %.55s\n",
				item.Number, item.User.Login, item.Title, item.User)
		}
	}
	if *st {
		resultS, err := github.SearchIssues(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("------Searchd issues------")
		fmt.Printf("%d issues:\n", resultS.TotalCount)
		for _, item := range resultS.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}

	if *ci {
		resp, err := github.CreateRepoIssues(*repo, *title, *body, *token)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("------Created issue------")
		fmt.Println(*resp)
	}
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/

// go run . -ci=true --token="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" --title="post test" --body="body go test" --repo="DagmarC/gopl-solutions"