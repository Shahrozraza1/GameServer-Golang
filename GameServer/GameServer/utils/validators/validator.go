package validators

func ValidateRectangle(levels [][]int) bool {
	firstLevelLength := len(levels[0])
	for _, level := range levels {
		if len(level) != firstLevelLength {
			return false
		}
	}
	return true
}

func ValidateLevelMaxLength(levels [][]int) bool {
	if len(levels) > 100 {
		return false
	}

	for _, level := range levels {
		if len(level) > 100 {
			return false
		}
	}

	return true
}

func ValidateMapSpaces(levels [][]int) bool {
	for _, level := range levels {
		for _, value := range level {
			if value > 2 || value < 0 {
				return false
			}
		}
	}
	return true
}
