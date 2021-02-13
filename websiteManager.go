package main 

//TODO à supprimer 
type websiteManager struct {
	allLinks []string
	log *logManager
	conf Configuration
}

func (websiteManager *websiteManager) scrollWebSite (website string){
	i := 0

	websiteManager.allLinks = append(websiteManager.allLinks, website)

	for i< len(websiteManager.allLinks){	
		
		httpManager := httpManager{
			url : websiteManager.allLinks[i],
			logManager : websiteManager.log,
		}
		httpManager.getContentPage()
		
		saveDatas := saveDatas {
			url : websiteManager.allLinks[i],
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