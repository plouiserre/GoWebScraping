package main


func main(){
	httpManager := httpManager{
		//TODO must be in configuration file
		url : "https://www.jeuxvideo.com",
	}
	httpManager.getContentPage()

	contentManager := contentManager{
		contentPage : httpManager.contentPage,
	}

	contentManager.GetLinks()
}
