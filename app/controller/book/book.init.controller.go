package bookcontroller

func Init(
	bookRepository BookRepositoryItf,
) *BookController {
	return &BookController{
		bookRepository: bookRepository,
	}
}
