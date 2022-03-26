#!/bin/bash

# package: binaryen
wasm2js qCCYI0ajpD.wasm > qCCYI0ajpD.js;

# package: wabt
wasm-decompile qCCYI0ajpD.wasm -o qCCYI0ajpD.dcmp;

