package main

type PullRequest struct {
	author     string
	url        string
	open       bool
	projet     string
	approvals  int
	fromBranch string
	toBranch   string
}

var samplePullRequest PullRequest

var pullRequests []PullRequest

func loadSamplePRs() {
	samplePullRequest.author = "Renato"
	samplePullRequest.url = "http://blue.man"
	samplePullRequest.open = true
	samplePullRequest.projet = "jeitto"
	samplePullRequest.approvals = 1
	samplePullRequest.fromBranch = "homolog"
	samplePullRequest.toBranch = "master"
}

func getOpenPRs() []PullRequest {
	return pullRequests
}
