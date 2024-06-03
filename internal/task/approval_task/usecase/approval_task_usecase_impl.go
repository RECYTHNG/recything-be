package usecase

import (
	"github.com/sawalreverr/recything/internal/task/approval_task/repository"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
)

type ApprovalTaskUsecaseImpl struct {
	ApprovalTaskRepository repository.ApprovalTaskRepository
}

func NewApprovalTaskUsecase(approvalTaskRepository repository.ApprovalTaskRepository) *ApprovalTaskUsecaseImpl {
	return &ApprovalTaskUsecaseImpl{ApprovalTaskRepository: approvalTaskRepository}
}

func (usecase *ApprovalTaskUsecaseImpl) GetAllApprovalTaskPagination(limit int, offset int) ([]*user_task.UserTaskChallenge, int, error) {
	task, total, err := usecase.ApprovalTaskRepository.GetAllApprovalTaskPagination(limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return task, total, nil
}
