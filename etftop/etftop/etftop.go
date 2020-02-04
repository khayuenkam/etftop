package etftop

import (
	"sync"
	"time"

	"github.com/khayuenkam/etftop/etftop/api/tradegate"
	"github.com/khayuenkam/etftop/etftop/api/tradegate/types"
	"github.com/rivo/tview"
)

type State struct {
	allEtfs     []string
	etfsData    map[string]types.RefreshType
	refreshRate time.Duration
}

type EtfTop struct {
	debug         bool
	refreshTicker *time.Ticker
	State         *State
	table         *tview.Table
	forceRefresh  chan bool
	refreshMux    sync.Mutex
	api           *tradegate.Client
	app           *tview.Application
}

func NewEtfTop() (*EtfTop, error) {
	et := &EtfTop{
		debug: true,
		State: &State{
			allEtfs:     []string{"IE00BK5BQT80"},
			refreshRate: 60 * time.Second,
			etfsData:    make(map[string]types.RefreshType),
		},
		table:        tview.NewTable(),
		forceRefresh: make(chan bool),
		api:          tradegate.NewClient(nil),
		app:          tview.NewApplication(),
	}

	et.refreshTicker = time.NewTicker(et.State.refreshRate)

	return et, nil
}

func (et *EtfTop) Run() error {
	et.log("run()")

	et.UpdateEtfs()
	et.intervalFetchData()
	et.SetUpTable()

	if err := et.app.SetRoot(et.table, true).Run(); err != nil {
		return err
	}

	return nil
}
