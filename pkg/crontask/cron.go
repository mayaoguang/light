package crontask

import (
	"github.com/robfig/cron"
	"log"
)

// cronTask 一个定时任务
type (
	cronTask struct {
		f           func()
		Spec        string
		immediately bool // 是否立即执行
	}
	Crontab struct {
		cron  *cron.Cron
		tasks []*cronTask
	}
)

var CronManage *Crontab

func NewCrontab() *Crontab {
	if CronManage == nil {
		return &Crontab{}
	}
	return CronManage
}

// AddFunc 新增定时任务
func (s *Crontab) AddFunc(spec string, f func(), immediately bool) {
	s.tasks = append(s.tasks, &cronTask{
		f:           f,
		Spec:        spec,
		immediately: immediately,
	})
}

// Start 启动
func (s *Crontab) Start() {
	for _, item := range s.tasks {
		if err := s.cron.AddFunc(item.Spec, item.f); err != nil {
			log.Fatal(err.Error())
		}
		// 立刻执行
		if item.immediately {
			go item.f()
		}
	}
	s.cron.Start()
}
