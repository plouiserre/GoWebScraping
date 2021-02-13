package main


func main(){
	//TODO must be in configuration file
	confManager := Configuration{}
	conf := confManager.readConf()

	log := new(logManager)
	log.logFile = conf.LogFile

	log.initLog()

	log.writeLog("*---------------Start process---------------*","info")

	websiteManager := websiteManager{
		log : log,
		conf : conf,
	}
	websiteManager.scrollWebSite(conf.WebSiteAnalyse)

	log.writeLog("*---------------End process---------------*","info")
}
