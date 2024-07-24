package bufferedsequentialsolver

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

func (bss *BufferedSequentialSolver) Start() {
	file, err := os.Open(bss.Config.SourceFilePath)
	if err != nil {
		bss.Logger.Fatalf("could not open temperatures.csv: %v", err)
	}
	defer file.Close()

	bss.Logger.Println("starting calculations...")

	start := time.Now()

	if bss.Config.Profile.Enabled {
		exeFile, err := os.Create(bss.Config.Profile.ExecutionFilePath)
		if err != nil {
			log.Fatal("could not create trace execution profile: ", err)
		}
		defer exeFile.Close()
		trace.Start(exeFile)
		defer trace.Stop()

		cpuFile, err := os.Create(bss.Config.Profile.CPUFilePath)
		if err != nil {
			bss.Logger.Fatalf("could not create CPU profile: %v", err)
		}
		defer cpuFile.Close()
		if err := pprof.StartCPUProfile(cpuFile); err != nil {
			bss.Logger.Fatalf("could not start CPU profile: %v", err)
		}
		defer pprof.StopCPUProfile()

		memFile, err := os.Create(bss.Config.Profile.MemoryFilePath)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer memFile.Close()
		runtime.GC()
		if err := pprof.WriteHeapProfile(memFile); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}

	sortedStats := bss.ProcessTemperatures(file)

	end := time.Now()

	duration := end.Sub(start)

	bss.Logger.Printf("calculations completed in %s!", duration.String())

	bss.Logger.Println("writing results...")

	err = bss.CSVResultsWriter.Write(sortedStats)
	if err != nil {
		bss.Logger.Fatalf("%v", err)
	}
}
