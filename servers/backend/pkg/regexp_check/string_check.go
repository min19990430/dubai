package regexp_check

import (
	"regexp"
)

func IsNumeric(s string) bool {
	numeric := regexp.MustCompile(`^[0-9]+$`)
	return numeric.MatchString(s)
}

func IsInterval(s string) bool {
	interval := regexp.MustCompile(`^[0-9]+[smhd]$`)
	return interval.MatchString(s)
}

func IsDeciamlType(s string) bool {
	typeString := regexp.MustCompile("^Decimal[1-6]$")
	return typeString.MatchString(s)
}

func IsTimeZone(s string) bool {
	timezone := regexp.MustCompile(`^[\+\-](0[0-9]|1[0-3]):[0-5][0-9]$`)
	return timezone.MatchString(s)
}
