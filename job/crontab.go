package job

import "github.com/robfig/cron"

var Crontab *cron.Cron
func GetCron() *cron.Cron {
	c := cron.New()
	Crontab = c
	return  Crontab
}