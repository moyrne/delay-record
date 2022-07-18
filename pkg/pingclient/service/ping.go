package service

import (
	"fmt"
	"time"

	"github.com/moyrne/delay-record/pkg/pingclient/internal/biz"
	"github.com/moyrne/delay-record/pkg/pingclient/internal/data"
)

func Run(host string) {
	repo, err := data.NewPingCSV(fmt.Sprintf("delay_%s.csv", time.Now().Format("20060102")))
	if err != nil {
		panic(err)
	}

	defer repo.Close()

	client := biz.NewPingClient(host, repo)

	tick := time.Tick(time.Millisecond * 300)
	for {
		client.Ping()
		<-tick
	}
}
