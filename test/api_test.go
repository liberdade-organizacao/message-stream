package test

import (
    "testing"
    "net/http"
    "net/url"
    "io/ioutil"
    "strings"
)

const (
    rootUrl = "http://localhost:8000"
)

func TestRootApiCall(t *testing.T) {
    response, err := http.PostForm(rootUrl + "/", url.Values{
        "who": {"Joe"},
    })

    if err != nil {
        t.Log("Error should be nil")
        t.Log(err)
        t.Fail()
    }

    rawData, err := ioutil.ReadAll(response.Body)
    data := string(rawData[:])
    if !strings.Contains(data, "nil") {
        t.Log("Index post failed, response was: ", data)
        t.Fail()
    }
}

func TestSetup(t *testing.T) {
    response, err := http.PostForm(rootUrl + "/setup", url.Values{})

    if err != nil {
        t.Log("Error should be nil")
        t.Log(err)
        t.Fail()
    }

    rawData, err := ioutil.ReadAll(response.Body)
    data := string(rawData[:])
    if data != "ok" {
        t.Log("setup failed, response was: ", data)
        t.Fail()
    }
}
