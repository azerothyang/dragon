package test

import (
	"dragon/core/dragon/util/validate"
	"log"
	"testing"
)

func TestValidator(t *testing.T) {
	validator := validate.New()
	data := map[string]string{
		"int64param": "32",
		"int32param": "3213",
	}
	rules := map[string]string{
		"int64param": "int64",
		"int32param": "int32",
	}
	validator.Validate(&data, rules)
	if validator.HasErr == true {
		log.Fatal("test fail")
	}

	validator = validate.New()
	data = map[string]string{
		"int64param": "32",
		"int32param": "aasd",
	}
	rules = map[string]string{
		"int64param": "int64",
		"int32param": "int32",
	}
	validator.Validate(&data, rules)
	if validator.HasErr == false {
		log.Fatal("test fail")
	}

}
