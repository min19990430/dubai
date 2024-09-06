package validator

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func userValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		registerErr := v.RegisterValidation(
			"username",
			func(fl validator.FieldLevel) bool {
				return usernameValidator(5, 30, fl.Field().String()) == nil
			})
		if registerErr != nil {
			log.Fatalf("register username validator error: %s", registerErr.Error())
		}
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		registerErr := v.RegisterValidation(
			"password",
			func(fl validator.FieldLevel) bool {
				return passwordValidator(12, 30, fl.Field().String()) == nil
			})
		if registerErr != nil {
			log.Fatalf("register password validator error: %s", registerErr.Error())
		}
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		registerErr := v.RegisterValidation(
			"phone",
			func(fl validator.FieldLevel) bool {
				regex, compileErr := regexp.Compile(`^09[0-9]{8}$`)
				if compileErr != nil {
					return false
				}
				return regex.MatchString(fl.Field().String())
			})
		if registerErr != nil {
			log.Fatalf("register phone validator error: %s", registerErr.Error())
		}
	}
}

const (
	LevelD = iota
	LevelC
	LevelB
	LevelA
	LevelS
)

func usernameValidator(minLength, maxLength int, username string) error {
	if len(username) < minLength {
		return fmt.Errorf("the username is shorter than %d characters", minLength)
	}

	if len(username) > maxLength && maxLength > 1 {
		return fmt.Errorf("the username is logner than %d characters", maxLength)
	}

	minLengthStr := strconv.Itoa(minLength)
	maxLengthStr := strconv.Itoa(maxLength)
	if maxLength <= 0 {
		maxLengthStr = ""
	}

	pattern := "^([a-zA-Z0-9~!@#$%^&*?_-]){" + minLengthStr + "," + maxLengthStr + "}$"

	_, err := regexp.MatchString(pattern, username)
	if err != nil {
		return fmt.Errorf("the username does not legal")
	}

	return nil
}

func passwordValidator(minLength, maxLength int, pwd string) error {
	if len(pwd) < minLength {
		return fmt.Errorf("the password is shorter than %d characters", minLength)
	}

	if len(pwd) > maxLength {
		return fmt.Errorf("the password is logner than %d characters", maxLength)
	}

	level := LevelD

	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}

	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)

		if match {
			level++
		}
	}

	if level < LevelA {
		return fmt.Errorf("the password does not satisfy")
	}

	return nil
}
