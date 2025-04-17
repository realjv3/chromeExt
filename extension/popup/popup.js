// Initialize Go WASM
const go = new Go();

WebAssembly.instantiateStreaming(fetch('main.wasm'), go.importObject)
    .then((result) => {
        go.run(result.instance);

        // Call our Go function and display the result
        document.getElementById('output').textContent = goSayHello();
    })
    .catch(console.error);