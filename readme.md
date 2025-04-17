## Chrome extension boilerplate (with WebAssembly compiled from Go)

### To develop:
1. First, compile the Go code to WebAssembly `GOOS=js GOARCH=wasm go build -o extension/main.wasm wasm/main.go`
2. Copy the WebAssembly JavaScript support file from your Go installation: `cp "${go env GOROOT}/lib/wasm/wasm_exec.js extension/`
3. Load the extension in Chrome:
   1. Open Chrome and go to `chrome://extensions/`
   2. Enable "Developer mode"
   3. Click "Load unpacked" and select the `extension` directory

- Make sure you have Go 1.24 installed
- The extension uses Manifest V3, which is the current standard for Chrome extensions
- The WASM file needs to be served with the correct MIME type, but Chrome extensions handle this automatically 
- Remember to rebuild the WASM file whenever you make changes to the Go code
