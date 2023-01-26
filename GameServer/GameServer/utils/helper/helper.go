package helper

import (
	"GameServer/consts"
	"GameServer/utils/validators"
	"bytes"
	"encoding/json"
	"os"
)

// function to check if database file exists

func DoesFileExist(fileName string) bool {
	_, error := os.Stat(fileName)

	// check if error is "file not exists"
	if os.IsNotExist(error) {
		return false
	} else {
		return true
	}
}

/** Since we cannot add array in database we convert it in json text first which is this function does **/

func Encode2DArray(Levels [][]int) string {
	encodedArray, _ := json.Marshal(Levels)
	return bytes.NewBuffer(encodedArray).String()
}

/** This function converts the json array into relevant data type **/

func Decode2DArray(jsonArray string) [][]int {
	var level [][]int
	json.Unmarshal([]byte(jsonArray), &level)
	return level
}

func ValidateLevels(levels [][]int) consts.ERRORMESSAGE {
	if !validators.ValidateLevelMaxLength(levels) {
		return consts.LEVELLENGTHERROR
	}

	if !validators.ValidateRectangle(levels) {
		return consts.LEVELRECTANGULARERROR
	}

	if !validators.ValidateMapSpaces(levels) {
		return consts.LEVELSPACEERROR
	}

	return consts.PASSEDVALIDATIONS
}
