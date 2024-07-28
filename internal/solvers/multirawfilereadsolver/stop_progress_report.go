package multirawfilereadsolver

func (mrfrs *MultiRawFileReadSolver) StopProgressReport() {
	mrfrs.ProgressReporter.Stop()
}
