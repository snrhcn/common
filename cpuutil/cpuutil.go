package cpuutil

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "Where to write CPU profile")

// Run starts up stuff at the beginning of a main function, and returns a
// function to defer until the function completes.  See cpuutil_test.go
func Run() func() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatalf("could not open cpu profile file %q", *cpuprofile)
		}
		pprof.StartCPUProfile(f)
		return func() {
			pprof.StopCPUProfile()
			f.Close()
		}
	}
	return func() {}
}
