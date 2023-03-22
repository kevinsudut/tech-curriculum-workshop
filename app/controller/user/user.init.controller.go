package usercontroller

func Init(
	userRepository UserRepositoryItf,
) *UserController {
	return &UserController{
		userRepository: userRepository,
	}
}
