p := public/files/wasm/index.wasm

dev:
	go run example/serve/main.go
wasm_prod:
	GOARCH=wasm GOOS=js go build -ldflags '-w -s' -o $(p) example/client/main.go
	wasm-opt -Oz -o ${p}.opt ${p}
	du -sh $(p).opt
	gzip -5 -f -k $(p).opt && du -sh $(p).gz
wasm:
	GOARCH=wasm GOOS=js go build -ldflags '-w -s' -o $(p) example/client/main.go
ser:
	echo "serve-run"
tailwind:
	tailwindcss -i ./tailwind.css -o ./public/files/index.css --watch
build-client:
	echo "build-client"
build-serve:
	echo "build-client"
test:
	go test ./...go test ./...
