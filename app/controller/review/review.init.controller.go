package reviewcontroller

func Init(
	reviewRepository ReviewRepositoryItf,
	bookRepository BookRepositoryItf,
	userRepository UserRepositoryItf,
) *ReviewController {
	return &ReviewController{
		reviewRepository: reviewRepository,
		bookRepository:   bookRepository,
		userRepository:   userRepository,
	}
}
