package main 

import (
	"time"
	"fmt"
)

type websiteManager struct {
	allLinks []string
	log *logManager
	conf Configuration
	index int
}

type pagesAnalysed struct {
	count int
}

func (websiteManager *websiteManager) scrollWebSite (website string){
	i := 0

	websiteManager.allLinks = append(websiteManager.allLinks, website)

	pagesAnalysed := make(chan pagesAnalysed)

	for i< len(websiteManager.allLinks){		
		go websiteManager.analyzeWebPage(website, pagesAnalysed)
		numberPagedAnalysed := <- pagesAnalysed
		i = numberPagedAnalysed.count
	}
}

func (websiteManager *websiteManager) analyzeWebPage(website string, analyser chan pagesAnalysed){
		startAnalyzePage := time.Now()
		
		urlPage := websiteManager.allLinks[websiteManager.index]

		fmt.Println("Début traitement page ", urlPage)
		
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
		
		fmt.Println("fin traitement page ", urlPage)
		
		websiteManager.index ++
		analyser <-  pagesAnalysed {
			count : websiteManager.index, 
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