package repository

type ModelNotFoundError struct {
}

func (e *ModelNotFoundError) Error() string {
	return "Model not found"
}
