package main

func main(){
	//TODO must be in configuration file
	//webSiteAnalyse := "https://www.jeuxvideo.com"
	webSiteAnalyse := "https://vuejs.org"
	logFile := "/Users/plouiserre/Desktop/logs.txt"


	log := new(logManager)
	log.logFile = logFile

	log.initLog()

	log.writeLog("*---------------Start process---------------*","info")

	websiteManager := websiteManager{
		log : log,
	}
	websiteManager.scrollWebSite(webSiteAnalyse)

	log.writeLog("*---------------End process---------------*","info")
}
