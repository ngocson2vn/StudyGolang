package main

import (
	"fmt"
	"hash/crc32"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// type Logger interface {
// 	// Info logs routine messages about cron's operation.
// 	Info(msg string, keysAndValues ...interface{})
// 	// Error logs an error condition.
// 	Error(err error, msg string, keysAndValues ...interface{})
// }

type MyLogger struct{}

func (ml *MyLogger) Info(msg string, keysAndValues ...interface{}) {
	log.Printf(msg, keysAndValues...)
	fmt.Println()
}

func (ml *MyLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	log.Printf(msg, keysAndValues...)
	fmt.Println(err)
}

func main() {
	scheduler := cron.New()

	// cron.Job is a interface for submitted cron jobs
	// type Job interface {
	//     Run()
	// }

	// cron.FuncJob is a wrapper that turns a func() into a cron.Job
	// type FuncJob func()

	// cron.JobWrapper decorates the given Job with some behavior.
	// type JobWrapper func(Job) Job

	// cron.SkipIfStillRunning(cron.DefaultLogger) returns a JobWrapper

	ml := &MyLogger{}
	fmt.Printf("ml has type: %T\n", ml)

	// scheduler.AddJob("@every 1s", cron.SkipIfStillRunning(ml)(cron.FuncJob(func() {
	// 	time.Sleep(time.Second * 3)
	// 	fmt.Printf("Now: %v\n", time.Now())
	// })))

	d1 := cron.JobWrapper(func(j cron.Job) cron.Job {
		startAt := time.Now().UTC()
		formattedStartAt := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", startAt.Year(), startAt.Month(), startAt.Day(), startAt.Hour(), startAt.Minute(), startAt.Second())
		fmt.Printf("[%s] d1 is applied\n", formattedStartAt)
		return j
	})

	d2 := cron.JobWrapper(func(j cron.Job) cron.Job {
		startAt := time.Now().UTC()
		formattedStartAt := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", startAt.Year(), startAt.Month(), startAt.Day(), startAt.Hour(), startAt.Minute(), startAt.Second())
		fmt.Printf("[%s] d2 is applied\n", formattedStartAt)
		return j
	})

	d3 := cron.JobWrapper(func(j cron.Job) cron.Job {
		startAt := time.Now().UTC()
		formattedStartAt := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", startAt.Year(), startAt.Month(), startAt.Day(), startAt.Hour(), startAt.Minute(), startAt.Second())
		fmt.Printf("[%s] d3 is applied\n", formattedStartAt)
		return j
	})

	// Then(j Job) will perform the following steps:
	//   - j = d3(j); return j
	//   - j = d2(j); return j
	//   - j = d1(j); return j
	//   - Run j every 5s
	scheduler.AddJob("@every 5s", cron.NewChain(d1, d2, d3).Then(cron.FuncJob(func() {
		startAt := time.Now().UTC()
		formattedStartAt := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", startAt.Year(), startAt.Month(), startAt.Day(), startAt.Hour(), startAt.Minute(), startAt.Second())
		jobId := crc32.ChecksumIEEE([]byte(formattedStartAt))
		fmt.Printf("\n[%s][%d] Start job\n", formattedStartAt, jobId)

		time.Sleep(time.Second * 3)

		endAt := time.Now().UTC()
		formattedEndAt := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", endAt.Year(), endAt.Month(), endAt.Day(), endAt.Hour(), endAt.Minute(), endAt.Second())
		fmt.Printf("[%s][%d] End job\n", formattedEndAt, jobId)
	})))

	scheduler.Start()
	fmt.Scanln()
}
