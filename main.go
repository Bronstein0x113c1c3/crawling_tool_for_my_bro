package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

func crawling(page int) error {
	resp, err := http.Get(fmt.Sprintf("http://tailieuso.tlu.edu.vn/flowpaper/services/view.php?doc=57346364846697234711889407769857573585&format=jpg&page=%v", page) + "&subfolder=57%2F34%2F63%2F")
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	os.WriteFile(fmt.Sprintf("./workdir/page%v.jpg", page), data, 0644)
	return err
}

func main() {
	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go func(page int) {
			if err := crawling(page); err != nil {
				log.Println("Error at page ", page)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
