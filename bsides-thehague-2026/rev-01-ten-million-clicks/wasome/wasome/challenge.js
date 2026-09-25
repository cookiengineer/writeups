
function Table(ret) {
	// grow method not included; table is not growable
	ret.set = function(i, func) {
		this[i] = func;
	};
	ret.get = function(i) {
		return this[i];
	};
	return ret;
}

var bufferView;
var base64ReverseLookup = new Uint8Array(123/*'z'+1*/);
for (var i = 25; i >= 0; --i) {
	base64ReverseLookup[48+i] = 52+i; // '0-9'
	base64ReverseLookup[65+i] = i; // 'A-Z'
	base64ReverseLookup[97+i] = 26+i; // 'a-z'
}
base64ReverseLookup[43] = 62; // '+'
base64ReverseLookup[47] = 63; // '/'
/** @noinline Inlining this function would mean expanding the base64 string 4x times in the source code, which Closure seems to be happy to do. */
function base64DecodeToExistingUint8Array(uint8Array, offset, b64) {
	var b1, b2, i = 0, j = offset, bLength = b64.length, end = offset + (bLength*3>>2) - (b64[bLength-2] == '=') - (b64[bLength-1] == '=');
	for (; i < bLength; i += 4) {
		b1 = base64ReverseLookup[b64.charCodeAt(i+1)];
		b2 = base64ReverseLookup[b64.charCodeAt(i+2)];
		uint8Array[j++] = base64ReverseLookup[b64.charCodeAt(i)] << 2 | b1 >> 4;
		if (j < end) uint8Array[j++] = b1 << 4 | b2 >> 2;
		if (j < end) uint8Array[j++] = b2 << 6 | base64ReverseLookup[b64.charCodeAt(i+3)];
	}
}
function initActiveSegments(imports) {
	base64DecodeToExistingUint8Array(bufferView, 1024, "GBYdAc874vUD5gUNEeHDIYL11fWrnarXh4XLhGtckZqzvI+UneDX4wA=");
}
function asmFunc(imports) {
	var buffer = new ArrayBuffer(131072);
	var HEAP8 = new Int8Array(buffer);
	var HEAP16 = new Int16Array(buffer);
	var HEAP32 = new Int32Array(buffer);
	var HEAPU8 = new Uint8Array(buffer);
	var HEAPU16 = new Uint16Array(buffer);
	var HEAPU32 = new Uint32Array(buffer);
	var HEAPF32 = new Float32Array(buffer);
	var HEAPF64 = new Float64Array(buffer);
	var Math_imul = Math.imul;
	var Math_fround = Math.fround;
	var Math_abs = Math.abs;
	var Math_clz32 = Math.clz32;
	var Math_min = Math.min;
	var Math_max = Math.max;
	var Math_floor = Math.floor;
	var Math_ceil = Math.ceil;
	var Math_trunc = Math.trunc;
	var Math_sqrt = Math.sqrt;
	var global$1 = 1024;
	var global$2 = 1152;
	var global$3 = 1152;
	var global$4 = 66688;
	var global$5 = 1024;
	var global$6 = 66688;
	var global$7 = 131072;
	var global$8 = 0;
	var global$9 = 1;
	function __wasm_call_ctors() {

	}

	function press_button() {
		var $0 = 0;
		block : {
			$0 = HEAP32[(0 + 1072 | 0) >> 2] | 0;
			if ($0 >>> 0 > 9999999 >>> 0) {
				break block
			}
			$0 = $0 + 1 | 0;
			HEAP32[(0 + 1072 | 0) >> 2] = $0;
		}
		block1 : {
			if (($0 | 0) != (1e7 | 0)) {
				break block1
			}
			HEAP8[(0 + 1076 | 0) >> 0] = 1;
		}
		return $0 | 0;
	}

	function get_counter() {
		return HEAP32[(0 + 1072 | 0) >> 2] | 0 | 0;
	}

	function set_counter() {
		HEAP32[(0 + 1072 | 0) >> 2] = 10000000;
	}

	function decrypt_flag() {
		var $0 = 0, $1 = 0;
		$0 = 0;
		block : {
			if (!(HEAPU8[(0 + 1076 | 0) >> 0] | 0)) {
				break block
			}
			$0 = -40;
			label : while (1) {
				HEAP8[($0 + 1130 | 0) >> 0] = 0;
				$1 = Math_imul($0, 7);
				HEAP8[($0 + 1129 | 0) >> 0] = ((HEAPU8[($0 + 1065 | 0) >> 0] | 0) ^ ($1 + 352 | 0) | 0) - (($0 + 41 | 0) ^ 19 | 0) | 0;
				HEAP8[($0 + 1128 | 0) >> 0] = ((HEAPU8[($0 + 1064 | 0) >> 0] | 0) ^ ($1 + 89 | 0) | 0) - (($0 + 40 | 0) ^ 19 | 0) | 0;
				$0 = $0 + 2 | 0;
				if ($0) {
					continue label
				}
				break label;
			};
			$0 = 1;
		}
		return $0 | 0;
	}

	function get_flag_ptr() {
		return 1088 | 0;
	}

	bufferView = HEAPU8;
	initActiveSegments(imports);
	var FUNCTION_TABLE = Table([]);
	function __wasm_memory_size() {
		return buffer.byteLength / 65536 | 0;
	}

	return {
		"memory": Object.create(Object.prototype, {
			"grow": {

			},
			"buffer": {
				"get": function () {
					return buffer;
				}

			}
		}),
		"__wasm_call_ctors": __wasm_call_ctors,
		"press_button": press_button,
		"get_counter": get_counter,
		"set_counter": set_counter,
		"decrypt_flag": decrypt_flag,
		"get_flag_ptr": get_flag_ptr,
		"__indirect_function_table": FUNCTION_TABLE,
		"__dso_handle": {
			get value() {
				return global$1;
			},
			set value(_global$1) {
				global$1 = _global$1;
			}
		},
		"__data_end": {
			get value() {
				return global$2;
			},
			set value(_global$2) {
				global$2 = _global$2;
			}
		},
		"__stack_low": {
			get value() {
				return global$3;
			},
			set value(_global$3) {
				global$3 = _global$3;
			}
		},
		"__stack_high": {
			get value() {
				return global$4;
			},
			set value(_global$4) {
				global$4 = _global$4;
			}
		},
		"__global_base": {
			get value() {
				return global$5;
			},
			set value(_global$5) {
				global$5 = _global$5;
			}
		},
		"__heap_base": {
			get value() {
				return global$6;
			},
			set value(_global$6) {
				global$6 = _global$6;
			}
		},
		"__heap_end": {
			get value() {
				return global$7;
			},
			set value(_global$7) {
				global$7 = _global$7;
			}
		},
		"__memory_base": {
			get value() {
				return global$8;
			},
			set value(_global$8) {
				global$8 = _global$8;
			}
		},
		"__table_base": {
			get value() {
				return global$9;
			},
			set value(_global$9) {
				global$9 = _global$9;
			}
		}
	};
}

var retasmFunc = asmFunc({
});
// export var memory = retasmFunc.memory;
// export var __wasm_call_ctors = retasmFunc.__wasm_call_ctors;
// export var press_button = retasmFunc.press_button;
// export var get_counter = retasmFunc.get_counter;
// export var decrypt_flag = retasmFunc.decrypt_flag;
// export var get_flag_ptr = retasmFunc.get_flag_ptr;
// export var __dso_handle = retasmFunc.__dso_handle;
// export var __data_end = retasmFunc.__data_end;
// export var __stack_low = retasmFunc.__stack_low;
// export var __stack_high = retasmFunc.__stack_high;
// export var __global_base = retasmFunc.__global_base;
// export var __heap_base = retasmFunc.__heap_base;
// export var __heap_end = retasmFunc.__heap_end;
// export var __memory_base = retasmFunc.__memory_base;
// export var __table_base = retasmFunc.__table_base;

window.WASM = retasmFunc;

