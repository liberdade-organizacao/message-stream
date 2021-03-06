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
    secret = "mysecretpassword"
)

func TestRootApiCall(t *testing.T) {
    response, err := http.PostForm(rootUrl + "/", url.Values{})

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
    response, err := http.PostForm(rootUrl + "/setup", url.Values{
        "key": {secret},
    })

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

func TestNewMessage(t *testing.T) {
    response, err := http.PostForm(rootUrl + "/create", url.Values{
        "kind": {"mail"},
        "content": {"that's an email, please send me"},
        "key": {secret},
    })

    if err != nil {
        t.Log("Error should be nil")
        t.Log(err)
        t.Fail()
    }

    rawData, err := ioutil.ReadAll(response.Body)
    data := string(rawData[:])
    if data != "ok" {
        t.Log("message creation failed, response was: ", data)
        t.Fail()
    }
}

func TestGetMessages(t *testing.T) {
    response, err := http.PostForm(rootUrl + "/read", url.Values{
        "kind": {"mail"},
        "offset": {"0"},
        "key": {secret},
    })

    if err != nil {
        t.Log("Error should be nil")
        t.Log(err)
        t.Fail()
    }

    rawData, err := ioutil.ReadAll(response.Body)
    data := string(rawData[:])
    if !strings.Contains(data, "---") || !strings.Contains(data, "...") {
        t.Log("failed getting messages, response was: ", data)
        t.Fail()
    }
}

func TestSecretValidation(t *testing.T) {
    response, err := http.PostForm(rootUrl + "/read", url.Values{
        "kind": {"mail"},
        "offset": {"0"},
        "key": {"something else that is not the secret"},
    })

    if err != nil {
        t.Log("Error should be nil")
        t.Log(err)
        t.Fail()
    }

    rawData, err := ioutil.ReadAll(response.Body)
    data := string(rawData[:])
    if !strings.Contains(data, "error") {
        t.Log("failed not getting messages, response was: ", data)
        t.Fail()
    }
}
