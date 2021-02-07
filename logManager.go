package main 

import(
	"log"
	"os"
)

type logManager struct{
	logFile string
	warningLogger *log.Logger
	infoLogger *log.Logger
	errorLogger *log.Logger
}

func (logManager *logManager) initLog (){
	file, err := os.OpenFile(logManager.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY,0666)
	if err != nil{
		log.Fatal(err)
	}

	logManager.infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    logManager.warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    logManager.errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func(logManager *logManager) writeLog(message string, logType string){
	if logType == "error"{
		logManager.errorLogger.Println(message)
	} else if logType == "warning"{
		logManager.warningLogger.Println(message)
	} else {
		logManager.infoLogger.Println(message)
	}
}