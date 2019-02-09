package parser

import (
    "go-web-crawler/model"
    "io/ioutil"
    "testing"
)

func TestParseProfile(t *testing.T) {
    contents, err := ioutil.ReadFile("profile_test_data.html")
    if err != nil {
        panic(err)
    }

    result := ParseProfile(contents)  // "安静的雪"

    if len(result.Items) != 1 {
        t.Errorf("Items should contain only 1 element; but got %v",  result.Items)

    }

    profile := result.Items[0].(model.Profile)
    expected := model.Profile{
        Age:            34,
        Height:         162,
        Weight:         57,
        Income:         "3001-5000元",
        Gender:         "女",
        Constellations: "牡羊座",
        Occupation:     "人事/行政",
        Marriage:       "离异",
        House:          "已购房",
        Education:      "大学本科",
        Car:            "未购车",
    }

    if expected != profile {
        t.Errorf("Expected %v; but got %v", expected, profile)
    }
}