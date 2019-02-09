package parser

import (
    "go-web-crawler/engine"
    "go-web-crawler/model"
    "regexp"
    "strconv"
)

var ageRe      = regexp.MustCompile(`<td><span class="label">年龄: </span>([\d]+)岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况: </span>([^<]+)</td>`)


func ParseProfile(contents []byte,  url string, name string) engine.ParseResult{

    profile := model.Profile{}
    profile.Name = name

    age, err := strconv.Atoi(extractString(contents, ageRe))
    if err != nil {
        profile.Age = age
    }
    profile.Marriage = extractString(contents, marriageRe)

    result := engine.ParseResult{
        Items: []engine.Item{
            Url:     "",
            Type:    "zhenai",
            Payload: profile,
        },
    }
    return result
}


func extractString(contents []byte, re *regexp.Regexp) string {
    match := re.FindSubmatch(contents)

    if len(match) >= 2 {
        return string(match[1])
    } else {
        return ""
    }
}