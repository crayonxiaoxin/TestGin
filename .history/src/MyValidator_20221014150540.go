package src

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func TopicUrl(fl validator.FieldLevel) bool {
	if v, ok := fl.Field().Interface().(string); ok {
		fmt.Printf("v: %v\n", v)
		if matched, _ := regexp.MatchString("^\\w{4,10}$", v); matched {
			return true
		}
	}
	return false
}

func TopicSize(fl validator.FieldLevel) bool {
	fmt.Printf("fl.Top(): %v\n", fl.Top())
	fmt.Printf("fl.Top(): %v\n", fl.Top().Interface())
	fmt.Printf("fl.Top(): %v\n", fl.Top().Interface().(Topics))
	if v, ok := fl.Top().Interface().(*Topics); ok {
		fmt.Printf("v: %v\n", v)
		if len(v.TopicList) == v.TopicListSize {
			return true
		}
	}
	return false
}
