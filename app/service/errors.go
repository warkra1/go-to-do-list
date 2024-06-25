package service

type AuthError struct {
	message string
}

func (e *AuthError) Error() string {
	return e.message
}

func NewUserNotFoundAuthError() *AuthError {
	return &AuthError{message: "User not found!"}
}

func NewInvalidPasswordAuthError() *AuthError {
	return &AuthError{message: "Invalid password!"}
}

func NewUserAlreadyExistsAuthError() *AuthError {
	return &AuthError{message: "User Already Exists"}
}

type ToDoListError struct {
	message string
}

func (e *ToDoListError) Error() string {
	return e.message
}

func NewInvalidNumberToDoListError() *ToDoListError {
	return &ToDoListError{message: "Invalid item number"}
}
