package usecase

import (
	"github.com/sawalreverr/recything/internal/task/approval_task/repository"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
	"github.com/sawalreverr/recything/pkg"
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

func (usecase *ApprovalTaskUsecaseImpl) ApproveUserTask(userTaskId string) error {
	if _, err := usecase.ApprovalTaskRepository.FindUserTask(userTaskId); err != nil {
		return pkg.ErrUserTaskNotFound
	}
	status := "accept"

	if err := usecase.ApprovalTaskRepository.ApproveUserTask(status, userTaskId); err != nil {
		return err
	}
	return nil

}
