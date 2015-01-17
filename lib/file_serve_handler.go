package lib

import (
	"log"
	"net/http"
	"strconv"
)

const ServeFilePath = "/serve"

func ServeFile(response http.ResponseWriter, request *http.Request) {
	//TODO this will be called often
	video, err := videoFromRequest(request)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	//TODO set content type depeding on video
	response.Header().Set("Content-Type", "video/mp4")
	log.Printf("Streaming: %v", video.encodedPath())
	http.ServeFile(response, request, video.encodedPath())

}

func videoFromRequest(request *http.Request) (Video, error) {
	var video Video

	err := request.ParseForm()
	if err != nil {
		return video, err
	}
	idString := request.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return video, err
	}

	result := Db.Find(&video, id)
	if result.Error != nil {
		return video, result.Error
	}
	return video, nil
}
