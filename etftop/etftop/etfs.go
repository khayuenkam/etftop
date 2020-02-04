package etftop

import (
	"sync"

	"github.com/khayuenkam/etftop/etftop/api/tradegate/types"
)

var etfslock sync.Mutex

type Data struct {
	isin   string
	result types.RefreshType
	err    error
}

func (et *EtfTop) fetchPrice(isin string, ch chan<- Data, wg *sync.WaitGroup) {
	defer wg.Done()
	result, err := et.api.Refresh(isin)

	if err != nil {
		ch <- Data{isin: isin, err: err}
		return
	}

	ch <- Data{isin: isin, result: *result}
}

func (et *EtfTop) UpdateEtfs() error {
	et.log("updateEtfs()")
	etfslock.Lock()
	defer etfslock.Unlock()

	ch := make(chan Data)
	var wg sync.WaitGroup

	for _, isin := range et.State.allEtfs {
		wg.Add(1)
		go et.fetchPrice(isin, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		et.State.etfsData[res.isin] = res.result
	}

	return nil
}
