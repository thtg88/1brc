package bufferedsequentialsolver

func (bss *BufferedSequentialSolver) StopProgressReport() {
	bss.ProgressReporter.Stop()
}
