package mathlib

import "gopkg.in/go-playground/validator.v8"

// ModStruct srtuct
type ModStruct struct {
	Number  float64 `validate:"required"`
	Divisor float64 `validate:"required"`
}

func validateMod(number, divisor float64) (*ModStruct, error) {
	validate := validator.New(&validator.Config{TagName: "validate"})
	mod := &ModStruct{
		Number:  number,
		Divisor: divisor,
	}
	errs := validate.Struct(mod)
	if errs != nil {
		return nil, errs
	}
	return mod, nil
}

// Mod Function Returns the remainder after number is divided by divisor. The result has the same sign as divisor.
func Mod(number, divisor float64) float64 {
	v, err := validateMod(number, divisor)
	if err != nil {
		return 0.0
	}
	sign := Sign(v.Divisor)
	number, err = Abs(v.Number)
	divisor, err = Abs(v.Divisor)

	return sign * ((number) - (divisor)*float64(int((number)/(divisor))))
}
