package filevalidator

type Repository interface {
}

type Validator struct {
}

func New() Validator {
	return Validator{}
}
