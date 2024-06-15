package repository

import (
	art "github.com/sawalreverr/recything/internal/article"
	"github.com/sawalreverr/recything/internal/database"
	rep "github.com/sawalreverr/recything/internal/report"
	ch "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	usr "github.com/sawalreverr/recything/internal/user"
)

type DashboardRepositoryImpl struct {
	DB database.Database
}

// GetTotalArticle implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetTotalArticle() (int, int, error) {
	var totalArticle int64
	var additionContentToday int64

	if err := d.DB.GetDB().Model(&art.Article{}).Count(&totalArticle).Error; err != nil {
		return 0, 0, err
	}
	if err := d.DB.GetDB().Model(&art.Article{}).Where("created_at >= ?", "now() + interval '1 day'").Count(&additionContentToday).Error; err != nil {
		return 0, 0, err
	}
	return int(totalArticle), int(additionContentToday), nil
}

// GetTotalVideo implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetTotalVideo() (int, int, error) {
	panic("unimplemented")
}

func NewDashboardRepository(db database.Database) DashboardRepository {
	return &DashboardRepositoryImpl{DB: db}
}

// GetReportLittering implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetReportLittering() (int, error) {
	panic("unimplemented")
}

// GetReportRubbish implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetReportRubbish() (int, error) {
	panic("unimplemented")
}

// GetTotalChange implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetTotalChallenge() (int, int, error) {
	var totalChange int64
	var additionChallengeSinceLastWeek int64

	if err := d.DB.GetDB().Model(&ch.TaskChallenge{}).Count(&totalChange).Error; err != nil {
		return 0, 0, err
	}

	if err := d.DB.GetDB().Model(&ch.TaskChallenge{}).Where("created_at > now() - interval '1 week'").Count(&additionChallengeSinceLastWeek).Error; err != nil {
		return 0, 0, err
	}

	return int(totalChange), int(additionChallengeSinceLastWeek), nil
}

// GetTotalReport implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetTotalReport() (int, int, error) {
	var totalReport int64
	var additionReportSinceYesterday int64

	if err := d.DB.GetDB().Model(&rep.Report{}).Count(&totalReport).Error; err != nil {
		return 0, 0, err
	}

	if err := d.DB.GetDB().Model(&rep.Report{}).Where("created_at > now() - interval '1 day'").Count(&additionReportSinceYesterday).Error; err != nil {
		return 0, 0, err
	}

	return int(totalReport), int(additionReportSinceYesterday), nil
}

// GetTotalUser implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetTotalUser() (int, int, error) {
	var totalUser int64
	var addtionUserSinceYesterday int64

	if err := d.DB.GetDB().Model(&usr.User{}).Count(&totalUser).Error; err != nil {
		return 0, 0, err
	}

	if err := d.DB.GetDB().Model(&usr.User{}).Where("created_at > now() - interval '1 day'").Count(&addtionUserSinceYesterday).Error; err != nil {
		return 0, 0, err
	}
	return int(totalUser), int(addtionUserSinceYesterday), nil
}

// GetUserClassic implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetUserClassic() (int, error) {
	panic("unimplemented")
}

// GetUserGold implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetUserGold() (int, error) {
	panic("unimplemented")
}

// GetUserPlatinum implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetUserPlatinum() (int, error) {
	panic("unimplemented")
}

// GetUserSilver implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetUserSilver() (int, error) {
	panic("unimplemented")
}
