package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

func UnicodeIndex(str string, subbeg string, subend string) []string {

	var result []string
	var begidx int
	var endidx int

	constr := str

	for {
		begidx = strings.Index(constr, subbeg)
		if begidx > 0{
			constr = constr[begidx+len(subbeg):]
			endidx = strings.Index(constr, subend)
			if endidx > 0 {
				item := constr[:endidx]
				result = append(result, item)
				constr = constr[endidx+len(subend):]
			} else {
				break
			}
		}else{
			break
		}
	}
	return result
}

func downimage (imagPath string, name string, wg *sync.WaitGroup)  {
	defer wg.Done()

	resp, err := http.Get(imagPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	out, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(out, bytes.NewReader(body))

	fmt.Println("image download done .", imagPath[:10], "......")
}

func main() {

	file, err := os.Open("langue.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	items := UnicodeIndex(string(bytes), "data-src=\"", "\"")
	wg := sync.WaitGroup{}
	for i, item := range items {
		wg.Add(1)
		go downimage(item, fmt.Sprintf("%d.png", i), &wg)
	}
	wg.Wait()
}