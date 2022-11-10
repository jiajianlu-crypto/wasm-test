all: main.wasm ./rust-wasm/pkg/rust_wasm.js ./emsdk/indexc.wasm ./emsdk/a.out.wasm
	live-server
main.wasm: wasm_exec.js index.go
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -gcflags='-m' -o main.wasm
wasm_exec.js:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ./
./rust-wasm/pkg/rust_wasm.js: ./rust-wasm/src/lib.rs
	cd ./rust-wasm && wasm-pack build --target web --features wee_alloc
./emsdk/indexc.wasm: ./emsdk/index.c
	cd ./emsdk && source ./emsdk_env.sh && emcc index.c -O3 -s SIDE_MODULE=1 -o indexc.wasm
./emsdk/a.out.wasm: ./emsdk/index-2.c
	cd ./emsdk && source ./emsdk_env.sh && emcc index-2.c -O3 -s WASM=1 -sEXPORTED_FUNCTIONS=_genGarbage -sEXPORTED_RUNTIME_METHODS=ccall,cwrap -sASSERTIONS -sALLOW_MEMORY_GROWTH



	