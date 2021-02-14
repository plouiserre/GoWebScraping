package main 

import (
	"strings"
	"os"
)

type saveDatas struct{
	completeNameFile string
	content string
	logManager *logManager
	url string
	conf Configuration
}

func (save *saveDatas)saveWebPage(){
	save.getFileName()

	f, errCreateFile := os.Create(save.completeNameFile)

	if errCreateFile != nil {
		save.logManager.writeLog("error "+errCreateFile.Error(),"error")
	}

	defer f.Close()

	_, errWriteFile := f.WriteString(save.content)

	if errWriteFile != nil {
		save.logManager.writeLog("error "+errWriteFile.Error(),"error")
	}
}

func (save *saveDatas) getFileName(){
	var nameFile string
	urlParts := strings.Split(save.url,"/")
	lastSectionURL := urlParts[len(urlParts)-1]	
	if len(lastSectionURL) >0 {
		nameFile = lastSectionURL
	} else {
		nameFile = urlParts[len(urlParts)-2]
	}
	save.completeNameFile = save.conf.SaveWebSite+"/"+nameFile
}