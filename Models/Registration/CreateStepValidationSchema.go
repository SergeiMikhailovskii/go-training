package Registration

import "TestProject/Models/Base"

type CreateStepValidationSchemaType struct {
	FirstName []Base.Validator `json:"first_name"`
	LastName  []Base.Validator `json:"last_name"`
	Email     []Base.Validator `json:"email"`
	Password  []Base.Validator `json:"password"`
}

var CreateStepValidationSchema = CreateStepValidationSchemaType{
	FirstName: []Base.Validator{
		Base.RequiredValidator(),
		Base.OnlyLettersValidator(),
	},
	LastName: []Base.Validator{
		Base.RequiredValidator(),
		Base.OnlyLettersValidator(),
	},
	Email: []Base.Validator{
		Base.RequiredValidator(),
		Base.EmailValidator(),
	},
	Password: []Base.Validator{
		Base.RequiredValidator(),
		Base.RegExpValidator("^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*_)[a-zA-Z0-9_]+$"),
		Base.MinLengthValidator(5),
	},
}
