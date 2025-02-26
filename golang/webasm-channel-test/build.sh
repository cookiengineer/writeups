#!/bin/bash

GOROOT="$(go env GOROOT)";
ROOT="$(pwd)";

build_wasm() {

	env GOOS="js" GOARCH="wasm" go build -o "${ROOT}/public/main.wasm" "${ROOT}/main.go";

	if [[ "$?" == "0" ]]; then

		cp "${GOROOT}/lib/wasm/wasm_exec.js" "${ROOT}/public/wasm_exec.js";

		echo -e "- Generate WASM code: [\e[32mok\e[0m]";
		return 0

	else

		echo -e "- Generate WASM code: [\e[31mfail\e[0m]";
		return 1

	fi;

}

build_wasm;

