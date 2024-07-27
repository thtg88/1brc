package rawfilereadsolver

func (rfrs *RawFileReadSolver) StopProgressReport() {
	rfrs.ProgressReporter.Stop()
}
