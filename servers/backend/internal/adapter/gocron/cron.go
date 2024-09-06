package gocron

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"auto-monitoring/internal/adapter/gocron/job"

	"github.com/go-co-op/gocron/v2"
)

type App struct {
	jobs []job.IJob
}

func NewApp(jobs []job.IJob) *App {
	return &App{
		jobs: jobs,
	}
}

func Run(a *App) {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, j := range a.jobs {
		j.Setup(scheduler)
	}

	scheduler.Start()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// block until you are ready to shut down
	<-quit

	// when you're done, shut it down
	err = scheduler.Shutdown()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
