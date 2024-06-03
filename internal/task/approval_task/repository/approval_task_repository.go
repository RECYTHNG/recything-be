package repository

import (
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
)

type ApprovalTaskRepository interface {
	GetAllApprovalTaskPagination(limit int, offset int) ([]*user_task.UserTaskChallenge, int, error)
	FindUserTask(userTaskId string) (*user_task.UserTaskChallenge, error)
	ApproveUserTask(status string, userTaskId string) error
}
