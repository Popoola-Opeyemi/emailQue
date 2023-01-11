package shared

import (
	"emailQue/model"
	"emailQue/utils"
	"sync"
	"time"

	"github.com/go-pg/pg"
	"github.com/jinzhu/gorm"
)

// Worker ...
func Worker() error {
	dbc := utils.Env.DB
	log := utils.Env.Log

	records := []model.EmailtaskQueue{}
	gMon := GroupMonitor{}

	for {
		err := dbc.Where("status = ?", 0).Find(&records).Error
		if err != nil {
			if err == pg.ErrNoRows {
				// no tasks to perform sleep
				log.Debug("sleep 10")
				time.Sleep(20 * time.Second)
				continue
			}

			log.Debug(err)
			return err
		}

		if len(records) == 0 {
			// no tasks to perform sleep
			// log.Debug("sleep 10.1")
			time.Sleep(20 * time.Second)
			continue
		}

		// create workers to handle the tasks
		distribute(records, &gMon, 10)
	}

}

func distribute(tasks []model.EmailtaskQueue, gMon *GroupMonitor, maxWorkers int64) error {

	// check if the worker limit has been reached
	if gMon.Count() == maxWorkers {
		return nil
	}

	for i := 0; i < len(tasks); i++ {
		t := tasks[i]
		switch t.Type {
		case 1:
			// send email
			gMon.Add(1)
			func() {
				setTaskMode(sendEmail(t, gMon), t.ID)
			}()
		}

		// if worker limit has been reached, wait for a free worker
		if gMon.Count() >= maxWorkers {
			gMon.Available()
		}
	}

	return nil
}

func setTaskMode(failed bool, id int64) {
	que := model.EmailtaskQueue{ID: id}

	dbc := utils.Env.DB
	log := utils.Env.Log

	if failed {
		err := dbc.Model(&que).UpdateColumn("status", gorm.Expr("status - ?", 1)).Error
		if err != nil {
			log.Debug(err)
			return
		}
	} else {
		err := dbc.Model(&que).Where("id = ?", que.ID).Update("status", 1).Error
		if err != nil {
			log.Debug(err)
			return
		}
	}
}

// GroupMonitor ...
type GroupMonitor struct {
	mt    sync.Mutex
	count int64
}

// Add ...
func (s *GroupMonitor) Add(delta int64) {
	s.mt.Lock()
	defer func() {
		s.mt.Unlock()
	}()

	s.count += delta
}

// Done ...
func (s *GroupMonitor) Done() {
	s.Add(-1)
}

// Count ...
func (s *GroupMonitor) Count() int64 {
	s.mt.Lock()
	defer func() {
		s.mt.Unlock()
	}()

	return s.count
}

// Wait sleep until all workers are done
func (s *GroupMonitor) Wait() {

	for s.Count() != 0 {
		time.Sleep(1 * time.Second)
	}
}

// Available sleeps until a worker is available
func (s *GroupMonitor) Available() {
	cur := s.Count()
	for s.Count() >= cur {
		time.Sleep(1 * time.Second)
	}
}
