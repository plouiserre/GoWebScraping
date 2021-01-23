package main 

import(
	"fmt"
)

type websiteManager struct {
	allLinks []string
}

func (websiteManager *websiteManager) scrollWebSite (){
	httpManager := httpManager{
		//TODO must be in configuration file
		url : "https://www.jeuxvideo.com",
	}
	httpManager.getContentPage()

	contentManager := contentManager{
		contentPage : httpManager.contentPage,
	}

	contentManager.GetLinks()

	websiteManager.allLinks = append(websiteManager.allLinks, contentManager.links...)

	fmt.Println("liens enregistrés ", len(websiteManager.allLinks))
}

//1 - remplacer le main DID

//2 - stocker le résultat dans un array DID

//3 - passer à l'itération suivante de l'array