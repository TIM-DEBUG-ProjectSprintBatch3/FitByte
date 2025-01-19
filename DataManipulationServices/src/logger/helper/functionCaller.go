package functionCallerInfo

type FunctionCaller string

const (
	UserRepositoryCreate   FunctionCaller = "userRepository.Create"
	UserServiceRegister    FunctionCaller = "userService.Register"
	UserServiceUpdate      FunctionCaller = "userService.Update"
	UserControllerRegister FunctionCaller = "userController.Register"

	ActivityRepositoryCreate FunctionCaller = "activityRepository.Create"
	ActivityServiceCreate    FunctionCaller = "activityService.Create"
	ActivityControllerCreate FunctionCaller = "activityController.Create"
)
