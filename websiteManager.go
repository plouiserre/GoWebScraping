package main 

import(
	"fmt"
)

type websiteManager struct {
	allLinks []string
}

func (websiteManager *websiteManager) scrollWebSite (){
	i := 0
	//TODO must be in configuration file
	//webSiteAnalyse := "https://www.jeuxvideo.com"
	webSiteAnalyse := "https://vuejs.org/"
	websiteManager.allLinks = append(websiteManager.allLinks, webSiteAnalyse)

	for i< len(websiteManager.allLinks){	
		httpManager := httpManager{
			url : websiteManager.allLinks[i],
		}
		httpManager.getContentPage()

		contentManager := contentManager{
			contentPage : httpManager.contentPage,
		}

		contentManager.GetLinks()
		for _, link := range  contentManager.links{
			fmt.Println("lien étudié ", link)
			//websiteManager.allLinks = append(websiteManager.allLinks, contentManager.links...)
			contains := websiteManager.containsLink(link)
			if contains == false {
				websiteManager.allLinks = append(websiteManager.allLinks, link)
			}
		}
		fmt.Println("liens enregistrés ", len(websiteManager.allLinks))
		i ++
	}
}
func (websiteManager *websiteManager) containsLink(linkSearching string) bool{
	isPresent := false
	for _, link := range websiteManager.allLinks{
		if link == linkSearching{
			isPresent = true
			break
		}
	}
	return isPresent
}