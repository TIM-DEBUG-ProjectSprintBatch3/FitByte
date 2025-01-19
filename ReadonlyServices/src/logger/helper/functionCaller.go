package functionCallerInfo

type FunctionCaller string

const (
	UserRepositoryCreate   FunctionCaller = "userRepository.Create"
	UserServiceRegister    FunctionCaller = "userService.Register"
	UserControllerRegister FunctionCaller = "userController.Register"

	ActivityRepositoryGetAll FunctionCaller = "activityRepository.GetAll"
	ActivityServiceGetAll    FunctionCaller = "activityService.GetAll"
	ActivityControllerGetAll FunctionCaller = "activityController.GetAll"
)
