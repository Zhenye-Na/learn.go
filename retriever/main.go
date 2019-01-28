package main

import "fmt"

type Retriever interface {
	Get (str string) string
}

func download(r Retriever) string {
	return r.Get("www.imooc.com")
}

func main() {
	var r Retriever
	r = moock.Retriever{"Nothing to download here ..."}
	fmt.Println(download(r))
}
