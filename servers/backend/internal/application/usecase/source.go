package usecase

import (
	"errors"
	"fmt"
)

type Source int

const (
	Sensor Source = iota
	Device
	Debug
)

var sourceMap = map[Source]string{
	Sensor: "sensor",
	Device: "device",
	Debug:  "debug",
}

func (s Source) String() string {
	return sourceMap[s]
}

func (s Source) IsValid() bool {
	_, ok := sourceMap[s]
	return ok
}

func SourceFromString(s string) (Source, error) {
	for k, v := range sourceMap {
		if s == v {
			return k, nil
		}
	}

	return -1, errors.New(fmt.Sprintf("invalid source: %s", s))
}
