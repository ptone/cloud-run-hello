package bench

import (
	"runtime"
	"time"
)

func (b *Bench) UseCPU() {
	RunCPULoad(runtime.NumCPU(), b.RequestDurationMilli, b.CPUUtilizationPercent)
	return
}

// RunCPULoad run CPU load in specify cores count and percentage
func RunCPULoad(coresCount int, timeMilliSecond int, percentage int) {
	runtime.GOMAXPROCS(coresCount)

	// every loop : run + sleep = 1 unit

	// 1 unit = 100 ms may be the best
	unitHundresOfMicrosecond := 1000
	runMicrosecond := unitHundresOfMicrosecond * percentage
	sleepMicrosecond := unitHundresOfMicrosecond*100 - runMicrosecond
	for i := 0; i < coresCount; i++ {
		go func() {
			runtime.LockOSThread()
			// endless loop
			for {
				begin := time.Now()
				for {
					// run 100%
					if time.Now().Sub(begin) > time.Duration(runMicrosecond)*time.Microsecond {
						break
					}
				}
				// sleep
				time.Sleep(time.Duration(sleepMicrosecond) * time.Microsecond)
			}
		}()
	}
	// how long
	time.Sleep(time.Duration(timeMilliSecond) * time.Millisecond)
}
