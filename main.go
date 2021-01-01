package main

import (
	"sync"
	"time"

	"github.com/my-Sakura/crawl/config"
	"github.com/my-Sakura/crawl/news"
)

const (
	worker = 1
)

var (
	wg sync.WaitGroup
)

func main() {
	// c := cron.New()

	// c.AddFunc("0 0 9,12,18 ? *", start)

	// c.Start()
	// select {}
	start()
}

func start() {
	config.Set()

	now := time.Now()
	date := now.Format("2006-01-02")
	wg.Add(worker)

	go func() {
		news.StateDepartmentCrawlStart(date)
		wg.Done()
	}()
	// go func() {
	// 	data.EconomistFiftyCrawlStart()
	// 	wg.Done()
	// }()
	// go func() {
	// 	data.AcademicianCrawlStart()
	// 	wg.Done()
	// }()

	wg.Wait()
}
