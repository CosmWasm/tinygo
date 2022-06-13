//go:build cosmwasm
// +build cosmwasm

// NOTE(cosmwasm): this is a copy of runtime_wasm_wasi.go with the following modifications:
//   * sleepTicks and ticks are placeholders:
//     - should not be used by a contract;
//   * os_runtime_args simplified (tinygo 0.19 version);

package runtime

import (
	"unsafe"
)

type timeUnit int64

func ticksToNanoseconds(ticks timeUnit) int64 {
	return int64(ticks)
}

func nanosecondsToTicks(ns int64) timeUnit {
	return timeUnit(ns)
}

//export runtime.sleepTicks
// This function is called by the scheduler.
// Schedule a call to runtime.scheduler, do not actually sleep.
// A placeholder since it should never be used.
func sleepTicks(d timeUnit) {}

//export runtime.ticks
// A placeholder since it should never be used.
func ticks() timeUnit {
	return 1234567890
}

//export _start
func _start() {
	// These need to be initialized early so that the heap can be initialized.
	heapStart = uintptr(unsafe.Pointer(&heapStartSymbol))
	heapEnd = uintptr(wasm_memory_size(0) * wasmPageSize)
	run()
}

// This is the default set of arguments, if nothing else has been set.
// This may be overriden by modifying this global at runtime init (for example,
// on Linux where there are real command line arguments).
var args = []string{"/proc/self/exe"}

//go:linkname os_runtime_args os.runtime_args
func os_runtime_args() []string {
	return args
}

// putchar is a placeholder.
func putchar(c byte) {}
