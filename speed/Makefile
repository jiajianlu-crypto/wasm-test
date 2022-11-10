all: main.wasm ./emsdk/indexc.wasm ./emsdk/indexc.wasm-2
	live-server
main.wasm: wasm_exec.js index.go
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -gcflags='-m' -o main.wasm
wasm_exec.js:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ./
./emsdk/indexc.wasm: ./emsdk/index.c
	cd ./emsdk && source ./emsdk_env.sh && emcc index.c -O3 -s SIDE_MODULE=1 -o indexc.wasm
./emsdk/indexc.wasm-2: ./emsdk/index.c
	cd ./emsdk && source ./emsdk_env.sh && emcc index.c -O3 -s WASM=1 -sEXPORTED_FUNCTIONS=_fib,_cFib30Times -sEXPORTED_RUNTIME_METHODS=ccall,cwrap -sASSERTIONS -sALLOW_MEMORY_GROWTH