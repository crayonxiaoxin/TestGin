package src

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func TopicUrl(fl validator.FieldLevel) bool {
	// date, ok := fl.Field().Interface().(time.Time)
	// if ok {
	//   today := time.Now()
	//   if today.After(date) {
	// 	return false
	//   }
	// }
	// return true
	// fmt.Printf("fl.Field().String(): %v\n", fl.Field().String())
	if v, ok := fl.Field().Interface().(string); ok {
		fmt.Printf("v: %v\n", v)
	}
	return false
}
