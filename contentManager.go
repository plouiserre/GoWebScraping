package main 

import (
	"fmt"
	"regexp"
	"strings"
)

type contentManager struct {
	contentPage []byte	
	links []string
	forbiddenExtension []string
}

func (contentManager *contentManager) GetLinks(){
	content := string(contentManager.contentPage)
	
	//TODO must be in configuration file
	contentManager.forbiddenExtension = []string{".png",".css",".woff"}
	
	substract := regexp.MustCompile("href=\"(.*?)\"")
	resultHref := substract.FindAllString(content, -1)
	
	for _, href := range resultHref {
		contentManager.CheckLinks(href)
	}

	occurenceHref := len(contentManager.links)
	fmt.Println("occurence of hrefContent ",occurenceHref)
}

func (contentManager *contentManager) CheckLinks(href string){
	isAutorized := true
	
	for _, extension := range contentManager.forbiddenExtension{
		isAutorized =  !strings.Contains(href, extension)
		if !isAutorized{
			break;
		}
	}

	if isAutorized{
		href = strings.Replace(href,"href=\"","",-1)
		href = strings.Replace(href,"\"","",-1)
		contentManager.links = append(contentManager.links, href)	
	}
}

