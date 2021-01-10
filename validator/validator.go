package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func TopicUrl(fl validator.FieldLevel) bool{
	fmt.Print(fl.Top(), fl.Parent())
	return false
}
