package helper

import (
	achievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
	"github.com/sawalreverr/recything/internal/database"
)

type GetDB struct {
	DB database.Database
}

func (db *GetDB) GiveAchievement(point int) string {
	var achievement []*achievement.Achievement
	if err := db.DB.GetDB().Find(&achievement).Error; err != nil {
		return ""
	}

	for _, ach := range achievement {
		if point >= ach.TargetPoint {
			return ach.Level
		}
	}
	return ""
}
