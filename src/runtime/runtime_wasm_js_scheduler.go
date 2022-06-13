//go:build wasm && !wasi && !scheduler.none && !cosmwasm
// +build wasm,!wasi,!scheduler.none,!cosmwasm

package runtime

//export resume
func resume() {
	go func() {
		handleEvent()
	}()

	if wasmNested {
		minSched()
		return
	}

	wasmNested = true
	scheduler()
	wasmNested = false
}

//export go_scheduler
func go_scheduler() {
	if wasmNested {
		minSched()
		return
	}

	wasmNested = true
	scheduler()
	wasmNested = false
}
