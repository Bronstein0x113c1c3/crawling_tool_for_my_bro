package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

func crawling(page int) {
	resp, err := http.Get(fmt.Sprintf("http://tailieuso.tlu.edu.vn/flowpaper/services/view.php?doc=5547919287785293543991158337515674687&format=png&page=%v", page) + "&subfolder=55%2F47%2F91%2F")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	os.WriteFile(fmt.Sprintf("./workdir/page%v.png", page), data, 0644)
	wg.Done()
}

func main() {
	for i := 1; i <= 173; i++ {
		wg.Add(1)
		go crawling(i)
	}
	wg.Wait()
}
