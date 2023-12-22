package application

type TaskService interface {
	GetAllTasks() error
	CreateTask() error
	GetTask() error
	DeleteTask() error
}

type taskService struct {
}

func NewTaskService() TaskService {
	return &taskService{}
}

func (ts *taskService) GetAllTasks() error {
	return nil
}

func (ts *taskService) CreateTask() error {
	return nil
}

func (ts *taskService) GetTask() error {
	return nil
}

func (ts *taskService) DeleteTask() error {
	return nil
}
