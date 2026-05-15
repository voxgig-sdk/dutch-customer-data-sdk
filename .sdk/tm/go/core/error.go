package core

type DutchCustomerDataError struct {
	IsDutchCustomerDataError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewDutchCustomerDataError(code string, msg string, ctx *Context) *DutchCustomerDataError {
	return &DutchCustomerDataError{
		IsDutchCustomerDataError: true,
		Sdk:              "DutchCustomerData",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *DutchCustomerDataError) Error() string {
	return e.Msg
}
