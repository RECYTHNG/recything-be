package repository

import "github.com/sawalreverr/recything/internal/achievement/entity"

type AchievementRepository interface {
	Create(achievement *entity.Achievement) (*entity.Achievement, error)
	FindAll(limit int, offset int) ([]entity.Achievement, int, error)
	FindById(id int) (*entity.Achievement, error)
	Update(achievement *entity.Achievement) (*entity.Achievement, error)
	Delete(id int) error
}
