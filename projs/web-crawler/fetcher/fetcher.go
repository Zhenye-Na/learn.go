package fetcher

import (
    "bufio"
    "fmt"
    "golang.org/x/net/html/charset"
    "golang.org/x/text/encoding"
    "golang.org/x/text/encoding/unicode"
    "golang.org/x/text/transform"
    "io/ioutil"
    "log"
    "net/http"
    "time"
)

// Determine the encoding method of web page
func determineDecoding(r *bufio.Reader) encoding.Encoding {
    // peek first 1024 bytes
    bytes, err := r.Peek(1024)
    if err != nil {
        log.Printf("Fetcher error: %v", err)
        return unicode.UTF8
    }

    // determine encoding method
    e, _, _ := charset.DetermineEncoding(bytes, "")
    return e
}

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetcher(url string) ([]byte, error) {
    <- rateLimiter
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    } else {
        // defer close response
        defer response.Body.Close()

        if response.StatusCode != http.StatusOK {
            return nil, fmt.Errorf("wrong status code: %d", response.StatusCode)
        }

        // decode
        bodyReader := bufio.NewReader(response.Body)
        e := determineDecoding(bodyReader)
        utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

        return ioutil.ReadAll(utf8Reader)
    }
}