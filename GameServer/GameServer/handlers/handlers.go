package handlers

import (
	"GameServer/consts"
	"GameServer/service"
	"encoding/json"
	"net/http"
)

func GetLevelByIDHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	response, err := service.RetrieveLevelByID(request)
	if err != consts.PASSEDVALIDATIONS {
		json.NewEncoder(writer).Encode(err)
		return
	}
	json.NewEncoder(writer).Encode(response)
}

func DeleteLevelHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	response := service.DeleteALevel(request)
	json.NewEncoder(writer).Encode(response)
}

func PostAddALevelHandler(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	response, err := service.AddALevel(request)
	if err == consts.PASSEDVALIDATIONS {
		json.NewEncoder(write).Encode(response)
		return
	}
	json.NewEncoder(write).Encode(err)
}

func PutUpdateALevel(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	response := service.UpdateLevel(request)
	json.NewEncoder(writer).Encode(response)
}
