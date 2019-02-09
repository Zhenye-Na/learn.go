package main

import (
    "go-web-crawler/engine"
    "go-web-crawler/persist"
    "go-web-crawler/scheduler"
    "go-web-crawler/zhenai/parser"
)

// main function
func main() {
    e := engine.ConcurrentEngine{
        Scheduler:   &scheduler.QueuedScheduler{},
        WorkerCount: 10,
        ItemChan:    persist.ItemSaver(),
    }

    //e.Run(engine.Request{
    //    Url: "http://www.zhenai.com/zhenghun",
    //    ParserFunc: parser.ParseCityList,
    //})

    e.Run(engine.Request{
        Url: "http://www.zhenai.com/zhenghun/shanghai",
        ParserFunc: parser.ParseCity,
    })
}
