package regexp_check

import (
	"fmt"
	"strconv"

	"regexp"
)

const (
	LevelD = iota
	LevelC
	LevelB
	LevelA
	LevelS
)

func UsernameValidator(minLength, maxLength int, username string) error {
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

func PasswordValidator(minLength, maxLength int, pwd string) error {
	if len(pwd) < minLength {
		return fmt.Errorf("the password is shorter than %d characters", minLength)
	}

	if len(pwd) > maxLength {
		return fmt.Errorf("the password is logner than %d characters", maxLength)
	}

	var level int = LevelD

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
