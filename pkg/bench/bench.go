package bench

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/teh-cmc/mmm"
)

// config items

type Bench struct {
	CPUUtilizationPercent int `env:"BENCH_CPU_UTILIZATION" envDefault:"50"`
	RequestDurationMilli  int `env:"BENCH_REQ_DURATION" envDefault:"100"`
	ColdStartMilli        int `env:"BENCH_COLDSTART" envDefault:"0"`
	MemoryPerRequestKB    int `env:"BENCH_MEMORY_KB" envDefault:"1"`
	ConcurrencyMax        int `env:"BENCH_MAX_CONCURRENCY" envDefault:"1000"`
	// numCores              int `env:"BENCH_"`    removing for now - the idea would be to use less than provided cores
	memoryChunk mmm.MemChunk
}

func NewBench() (b *Bench) {
	b = &Bench{}

	if err := env.Parse(b); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return
}
