package main 

import (
	"time"
	"fmt"
)

type websiteManager struct {
	allLinks []string
	log *logManager
	conf Configuration
}

func (websiteManager *websiteManager) scrollWebSite (website string){
	i := 0

	websiteManager.allLinks = append(websiteManager.allLinks, website)

	for i< len(websiteManager.allLinks){	
		startAnalyzePage := time.Now()
		urlPage := websiteManager.allLinks[i]
		
		httpManager := httpManager{
			url : urlPage,
			logManager : websiteManager.log,
		}
		httpManager.getContentPage()
		
		saveDatas := saveDatas {
			url : urlPage,
			logManager : websiteManager.log,
			content : string(httpManager.contentPage),
			conf : websiteManager.conf,
		}

		saveDatas.saveWebPage()

		contentManager := contentManager{
			contentPage : httpManager.contentPage,
			conf : websiteManager.conf,
		}

		contentManager.GetLinks()
		for _, link := range  contentManager.links{
			linkStudy := website+link
			contains := websiteManager.containsLink(linkStudy)
			if contains == false {
				websiteManager.log.writeLog("nouveau lien enregistré "+linkStudy,"info")
				websiteManager.allLinks = append(websiteManager.allLinks, linkStudy)
			}
		}
		
		endAnalyze := time.Since(startAnalyzePage)
		
		msg := fmt.Sprintf("urlPage reading in : %s", endAnalyze)
	
		websiteManager.log.writeLog(msg,"info")

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