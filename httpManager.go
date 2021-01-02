package main 

import(
	"net/http"
	"io/ioutil"
	"fmt"
)

type httpManager struct {
	url string
	contentPage []byte
}

func (httpManager *httpManager) getContentPage(){
	if len(httpManager.url) > 0 {
		resp, err := http.Get(httpManager.url)

		if err != nil {
			fmt.Println("Something bad happen")
		} else {
			httpManager.readGetResult(resp)
		}
	} else {
		fmt.Println("You must defined an url")
	}
}

func (httpManager *httpManager) readGetResult(resp *http.Response){

	defer resp.Body.Close()
	
	html, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Something bad happen")
	}else {
		httpManager.contentPage = html
	}
}