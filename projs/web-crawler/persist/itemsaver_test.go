package persist

import (
    "context"
    "encoding/json"
    "go-web-crawler/engine"
    "go-web-crawler/model"
    "testing"
)

func TestSave(t *testing.T) {
    expected := engine.Item {
        Url:            "http://album.zhenai.com/u/108906739",
        Type:           "zhenai",
        Id:             "108906739",
        Payload: model.Profile {
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
        },
    }


    // TODO: Start ElasticSearch using Docker first
    client, err := elastic.NewClient(elastic.SetSniff(false))
    if err != nil {
        panic(err)
    }

    err = save(expected)
    if err != nil {
        panic(err)
    }

    response, err := client.Get().Index("dating_profile").Type(expected.Type).Id(expected.Id).Do(context.Background())
    if err != nil {
        panic(err)
    }

    var actual engine.Item
    json.Unmarshal(*response.Source, &actual)
    actualProfile, _ := model.FromJsonObj(actual.Payload)
    actual.Payload = actualProfile

    if actual != expected {
        t.Errorf("got %v; expected %v", actual, expected)
    }

}
