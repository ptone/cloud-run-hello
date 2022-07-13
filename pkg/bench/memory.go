package bench

import (
	"log"

	"github.com/teh-cmc/mmm"
)

func (b *Bench) AllocMem(kb uint) {

	var err error
	b.memoryChunk, err = mmm.NewMemChunk([1024]byte{}, kb)
	if err != nil {
		log.Fatal(err)
	}

	dd := [1024]byte{}
	for i := 0; i < int(kb); i++ {
		b.memoryChunk.Write(i, dd)
	}
	return
}

func (b *Bench) FreeMem() {
	if err := b.memoryChunk.Delete(); err != nil {
		log.Fatal(err)
	}
	return
}
