package main

import (
	"errors"
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
	data, _ := ioutil.ReadAll(resp.Body)
	if string(data) != "Error:Error converting document, make sure the conversion tool is installed and that correct user permissions are applied to the SWF Path directory<br/><br/>Click <a href='http://flowpaper.com/docs_php.jsp'>here</a> for more information on configuring FlowPaper with PH" {
		os.WriteFile(fmt.Sprintf("./workdir/page%v.jpg", page), data, 0644)
		return nil
	} else {
		return errors.New("Not found")
	}

}

func main() {
	for i := 1; i <= 500; i++ {
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
