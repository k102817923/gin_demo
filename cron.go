package main

import (
	"log"
	"time"

	"github.com/robfig/cron"

	"gin_demo/models"
)

func mainTest() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	// 会创建一个新的定时器, 持续你设定的时间d后发送一个channel消息
	t1 := time.NewTimer(time.Second * 10)
	// 阻塞select等待channel
	for {
		select {
		case <-t1.C:
			// 重置定时器, 让它重新开始计时
			t1.Reset(time.Second * 10)
		}
	}
}
