# githubIssuer
Tool that allows users create, read, update, and delete GitHub issues from the command line, invoking their preferred text editor when substantial text input is required.

flags:
ri bool -> get the repo issue
ci bool -> create the repo issue
st bool ->  github search terms with space e.g. "repo:golang/go is:open json decoder"
repo string -> githubUser/repository e.g DagmarC/gopol-solutions
title string (when ci is true)
body string (when ci is true)


Usage:
go run . -ci --token="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" --title="post test" --body="body go test" --repo="DagmarC/gopl-solutions"

go run . -ri --repo="DagmarC/gopl-solutions"

go run . -st repo:golang/go is:open json decoder

go run . -ui --token="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" --title="EDIIIT test" --body="body go test" --repo="DagmarC/gopl-solutions" --n=1
