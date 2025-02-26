
(async () => {

	const go = new Go();
	const wasm_buffer = await fetch("main.wasm").then((response) => response.arrayBuffer());
	const wasm_module = await WebAssembly.instantiate(wasm_buffer, go.importObject);

	go.run(wasm_module.instance);

})();

