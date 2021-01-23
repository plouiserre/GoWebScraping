package main 

import (
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
	contentManager.forbiddenExtension = []string{".png",".css",".woff",".json"}
	
	substract := regexp.MustCompile("href=\"(.*?)\"")
	resultHref := substract.FindAllString(content, -1)
	
	for _, href := range resultHref {
		contentManager.CheckLinks(href)
	}
}

func (contentManager *contentManager) CheckLinks(href string){
	isNotForbiddenExtension := true
	
	for _, extension := range contentManager.forbiddenExtension{
		isNotForbiddenExtension =  !strings.Contains(href, extension)
		
		if !isNotForbiddenExtension{
			break;
		}
	}

	if isNotForbiddenExtension{
		href = strings.Replace(href,"href=\"","",-1)
		href = strings.Replace(href,"\"","",-1)
		
		
		letters := []rune(href)
		if len(letters)>0{
			isLinkAutorized := letters[0] == '/'
			
			if isLinkAutorized{
				contentManager.links = append(contentManager.links, href)	
			}
		}
	}
}

