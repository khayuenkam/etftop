package etftop

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func (et *EtfTop) log(msg string) {
	if !et.debug {
		return
	}

	filename := "/tmp/etftop.log"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	text := fmt.Sprintf("%v %s\n", time.Now().Unix(), msg)

	if _, err := f.WriteString(text); err != nil {
		log.Fatal(err)
	}
}
