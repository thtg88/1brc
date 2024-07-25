package sequentialsolver

func (ss *SequentialSolver) StopProgressReport() {
	ss.ProgressReporter.Stop()
}
