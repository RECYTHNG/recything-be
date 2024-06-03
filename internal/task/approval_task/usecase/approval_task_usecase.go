package usecase

import (
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
)

type ApprovalTaskUsecase interface {
	GetAllApprovalTaskPagination(limit int, offset int) ([]*user_task.UserTaskChallenge, int, error)
	ApproveUserTask(userTaskId string) error
}
