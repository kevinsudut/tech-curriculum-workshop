package errors

var (

	/////////////////
	// 5XXXX GROUP //
	/////////////////

	///////////////////////
	// 50xxx : Server error
	InternalServer Errs = Errs{
		Code:   "50001",
		Reason: "Internal server error",
	}

	/////////////////
	// 4XXXX GROUP //
	/////////////////

	//////////////////////////////
	// 40xxx : General bad request
	BadRequest Errs = Errs{
		Code:   "40001",
		Reason: "Bad request",
	}

	ForbiddenRequest Errs = Errs{
		Code:   "40002",
		Reason: "Unable to authenticate request",
	}

	////////////////////////
	// 41xxx : Missing param
	MissingParam Errs = Errs{
		Code:   "41001",
		Reason: "Missing required parameter",
	}

	InvalidParam Errs = Errs{
		Code:   "41002",
		Reason: "Invalid parameter",
	}

	///////////////////////////
	// 42020++ : Validation error
	AlreadyRegistered Errs = Errs{
		Code:   "42021",
		Reason: "User already registered",
	}

	EmailDoesNotExists Errs = Errs{
		Code:   "42022",
		Reason: "Email does not exists",
	}

	InvalidUserPassword Errs = Errs{
		Code:   "42023",
		Reason: "Invalid user password",
	}
)
