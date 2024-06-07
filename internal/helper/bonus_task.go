package helper

func BonusTask(badegUser string, userPoint int) int {
	switch badegUser {
	case "classic":
		return userPoint + userPoint*10/100
	case "silver":
		return userPoint + userPoint*15/100
	case "gold":
		return userPoint + userPoint*20/100
	case "platinum":
		return userPoint + userPoint*25/100
	default:
		return userPoint
	}
}
