package parser

import (
    "go-web-crawler/engine"
    "regexp"
)

var (
    profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
    cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
    // re := regexp.MustCompile(cityRe)
    matches := profileRe.FindAllSubmatch(contents, -1)  // function returns [][][]byte

    result := engine.ParseResult{}

    for _, m := range matches {
        // result.Items    = append(result.Items, "User: " + string(m[2]))
        result.Requests = append(result.Requests, engine.Request{
            Url:        string(m[1]),
            ParserFunc: engine.NilParser,
        })
    }

    matches = cityUrlRe.FindAllSubmatch(contents, -1)  // function returns [][][]byte
    
    for _, m := range matches {
        // result.Items    = append(result.Items, "User: " + string(m[2]))
        result.Requests = append(result.Requests, engine.Request{
            Url:        string(m[1]),
            ParserFunc: ParseCity,
        })
    }
    return result
}