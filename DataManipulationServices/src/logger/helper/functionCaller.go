package functionCallerInfo

type FunctionCaller string

const (
	UserRepositoryLogin    FunctionCaller = "userRepository.Login"
	UserRepositoryCreate   FunctionCaller = "userRepository.Create"
	UserServiceLogin       FunctionCaller = "userService.Login"
	UserServiceRegister    FunctionCaller = "userService.Register"
	UserServiceUpdate      FunctionCaller = "userService.Update"
	UserControllerLogin    FunctionCaller = "userController.Login"
	UserControllerRegister FunctionCaller = "userController.Register"

	ActivityRepositoryCreate FunctionCaller = "activityRepository.Create"
	ActivityServiceCreate    FunctionCaller = "activityService.Create"
	ActivityControllerCreate FunctionCaller = "activityController.Create"

	ActivityRepositoryUpdate FunctionCaller = "activityRepository.Update"
	ActivityServiceUpdate    FunctionCaller = "activityService.Update"
	ActivityControllerUpdate FunctionCaller = "activityController.Update"
)
