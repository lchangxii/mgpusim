package runner

import (
	"flag"
	"fmt"

	"gitlab.com/akita/akita"
	"gitlab.com/akita/gcn3/benchmarks"
	"gitlab.com/akita/gcn3/driver"
	"gitlab.com/akita/gcn3/platform"
)

var timing = flag.Bool("timing", false, "Run detailed timing simulation.")
var parallel = flag.Bool("parallel", false, "Run the simulation in parallel.")
var isaDebug = flag.Bool("debug-isa", false, "Generate the ISA debugging file.")
var visTracing = flag.Bool("trace-vis", false,
	"Generate trace for visualization purposes.")
var verify = flag.Bool("verify", false, "Verify the emulation result.")
var memTracing = flag.Bool("trace-mem", false, "Generate memory trace")

// Runner is a class that helps running the benchmarks in the official samples.
type Runner struct {
	engine            akita.Engine
	GPUDriver         *driver.Driver
	kernelTimeCounter *driver.KernelTimeCounter
	Benchmark         benchmarks.Benchmark
}

// Init initializes the platform simulate
func (r *Runner) Init() {
	if *parallel {
		platform.UseParallelEngine = true
	}

	if *isaDebug {
		platform.DebugISA = true
	}

	if *visTracing {
		platform.TraceVis = true
	}

	if *memTracing {
		platform.TraceMem = true
	}

	r.kernelTimeCounter = driver.NewKernelTimeCounter()
	if *timing {
		r.engine, r.GPUDriver = platform.BuildNR9NanoPlatform(4)
	} else {
		r.engine, _, r.GPUDriver, _ = platform.BuildEmuPlatform()
	}
	r.GPUDriver.AcceptHook(r.kernelTimeCounter)
}

// Run runs the benchmark on the simulator
func (r *Runner) Run() {
	r.Benchmark.Run()
	if *verify {
		r.Benchmark.Verify()
	}
	r.engine.Finished()
	fmt.Printf("Kernel time: %.12f\n", r.kernelTimeCounter.TotalTime)
	fmt.Printf("Total time: %.12f\n", r.engine.CurrentTime())
}