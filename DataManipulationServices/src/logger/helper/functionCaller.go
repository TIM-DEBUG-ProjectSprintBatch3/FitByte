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

	ActivityRepositoryUpdate FunctionCaller = "activityRepository.Update"
	ActivityServiceUpdate    FunctionCaller = "activityService.Update"
	ActivityControllerUpdate FunctionCaller = "activityController.Update"

	ActivityRepositoryDelete FunctionCaller = "activityRepository.Delete"
	ActivityServiceDelete    FunctionCaller = "activityService.Delete"
	ActivityControllerDelete FunctionCaller = "activityController.Delete"
)
