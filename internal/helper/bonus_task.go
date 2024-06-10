package helper

func BonusTask(badegUser string, userPoint int) int {
	switch badegUser {
	case "https://res.cloudinary.com/dymhvau8n/image/upload/v1717758679/achievement_badge/cq2n246e6twuksnia62t.png":
		return userPoint + userPoint*10/100
	case "https://res.cloudinary.com/dymhvau8n/image/upload/v1717758731/achievement_badge/b8igluyain8bwyjusfpk.png":
		return userPoint + userPoint*15/100
	case "https://res.cloudinary.com/dymhvau8n/image/upload/v1717758761/achievement_badge/lazzyh9tytvb4rophbc3.png":
		return userPoint + userPoint*20/100
	case "https://res.cloudinary.com/dymhvau8n/image/upload/v1717758798/achievement_badge/xc8msr6agowzhfq8ss8a.png":
		return userPoint + userPoint*25/100
	default:
		return userPoint
	}
}
