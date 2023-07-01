package user

type IUserService interface {
	GetYoungOrOld(age int) string
}

func NewController(srv IUserService) *Controller {
	return &Controller{srv: srv}
}

type Controller struct {
	srv IUserService
}
