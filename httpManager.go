package main 

import(
	"net/http"
	"io/ioutil"
)

type httpManager struct {
	url string
	contentPage []byte
	logManager *logManager
}

func (httpManager *httpManager) getContentPage(){
	if len(httpManager.url) > 0 {
		resp, err := http.Get(httpManager.url)

		if err != nil {
			httpManager.logManager.writeLog("error "+err.Error(),"error")
		} else {
			httpManager.readGetResult(resp)
		}
	} else {
		httpManager.logManager.writeLog("You must defined an url","error")
	}
}

func (httpManager *httpManager) readGetResult(resp *http.Response){

	defer resp.Body.Close()
	
	html, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		httpManager.logManager.writeLog("Something bad happen","error")
	}else {
		httpManager.contentPage = html
	}
}