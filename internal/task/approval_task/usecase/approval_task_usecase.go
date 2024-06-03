package usecase

import (
	"github.com/sawalreverr/recything/internal/task/approval_task/dto"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
)

type ApprovalTaskUsecase interface {
	GetAllApprovalTaskPagination(limit int, offset int) ([]*user_task.UserTaskChallenge, int, error)
	ApproveUserTask(userTaskId string) error
	RejectUserTask(request *dto.RejectUserTaskRequest, userTaskId string) error
}
