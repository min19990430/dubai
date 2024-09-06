package job

import "github.com/go-co-op/gocron/v2"

type IJob interface {
	Setup(s gocron.Scheduler)
}
