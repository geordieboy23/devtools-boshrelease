package main

import (
	"net/http"
	"os"
)

const (
	JENKINS_HOST string = "JENKINS_HOST"
)

var jenkinsHostUrl string
var httpClient *http.Client

func init() {
	jenkinsHostUrl = os.Getenv(JENKINS_HOST)

	if jenkinsHostUrl == "" {
		panic("JENKINS_HOST is empty")
	}
}

func main() {

}
