package service

import (
	"GameServer/consts"
	"GameServer/db"
	"GameServer/model"
	"GameServer/utils/helper"
	"encoding/json"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func AddALevel(request *http.Request) (*model.Level, consts.ERRORMESSAGE) {
	var level model.Level
	json.NewDecoder(request.Body).Decode(&level)

	validationResult := helper.ValidateLevels(level.Levels)
	if validationResult == "" {
		level.ID = uuid.NewV4().String()
		db.GetManagerInstance().AddLevelToDatabase(&level)
		return &level, ""
	}

	return nil, validationResult
}

func RetrieveLevelByID(request *http.Request) (*model.Level, consts.ERRORMESSAGE) {

	params := mux.Vars(request)
	response, err := db.GetManagerInstance().RetrieveLevelByID(params["ID"])
	if err == "" {
		return response, ""
	}
	return nil, consts.LEVELNOTFOUND
}

//
func UpdateLevel(request *http.Request) consts.ERRORMESSAGE {
	var level model.Level
	json.NewDecoder(request.Body).Decode(&level)

	validationResult := helper.ValidateLevels(level.Levels)
	if validationResult == "" {
		return db.GetManagerInstance().UpdateALevel(&level)
	}

	return validationResult
}

func DeleteALevel(request *http.Request) consts.ERRORMESSAGE {
	params := mux.Vars(request)
	response := db.GetManagerInstance().DeleteALevel(params["ID"])
	return response
}
