package schedule

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

func ProcessScheduler() {
	t := time.Now().UTC()
	t = t.In(time.FixedZone("KST", 9*60*60))
	t.Format("2006-01-12 15:04:05")
	s := gocron.NewScheduler(t.Location())

	s.Every(1).Day().At("07:00;20:00").Do(ProcessSync)

	s.StartAsync()
}

func ProcessSync() {
	fmt.Println("스케줄러")
}
