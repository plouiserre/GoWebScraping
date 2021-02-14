package main

import (
	"time"
	"fmt"
)

func main(){
	startProgram := time.Now()
	
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

	endProgram := time.Since(startProgram)

	msg := fmt.Sprintf("duration: %s", endProgram)
	
	log.writeLog(msg,"info")

	log.writeLog("*---------------End process---------------*","info")
}
