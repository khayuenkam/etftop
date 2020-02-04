package etftop

func (et *EtfTop) refreshAll() error {
	et.log("refreshAll()")
	go func() {
		et.UpdateEtfs()
		et.RefreshTable()
		et.app.Draw()
	}()
	return nil
}

func (et *EtfTop) intervalFetchData() {
	et.log("intervalFetchData()")

	go func() {
		for {
			select {
			case <-et.forceRefresh:
				et.refreshAll()
			case <-et.refreshTicker.C:
				et.refreshAll()
			}
		}
	}()
}
