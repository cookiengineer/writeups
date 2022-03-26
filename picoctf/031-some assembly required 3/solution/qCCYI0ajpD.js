
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
  base64DecodeToExistingUint8Array(bufferView, 1024, "nW6TyLK5QYvCkItkx57JiGKVkZDaY8WVldgyxMWSjmWSlpeMYcSTkpAAAA==");
  base64DecodeToExistingUint8Array(bufferView, 1067, "8afwB+0=");
}
function asmFunc(env) {
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
 var abort = env.abort;
 var nan = NaN;
 var infinity = Infinity;
 var global$0 = 66864;
 var global$1 = 1072;
 var global$2 = 1067;
 var global$3 = 1024;
 var global$4 = 1328;
 var global$5 = 1024;
 var global$6 = 66864;
 var global$7 = 0;
 var global$8 = 1;
 function $0() {
  
 }
 
 function $1($0_1, $1_1) {
  $0_1 = $0_1 | 0;
  $1_1 = $1_1 | 0;
  var $4 = 0, $7 = 0, $11 = 0;
  $4 = global$0 - 32 | 0;
  HEAP32[($4 + 24 | 0) >> 2] = $0_1;
  HEAP32[($4 + 20 | 0) >> 2] = $1_1;
  HEAP32[($4 + 16 | 0) >> 2] = HEAP32[($4 + 24 | 0) >> 2] | 0;
  HEAP32[($4 + 12 | 0) >> 2] = HEAP32[($4 + 20 | 0) >> 2] | 0;
  label$1 : {
   label$2 : while (1) {
    $7 = HEAP32[($4 + 16 | 0) >> 2] | 0;
    HEAP32[($4 + 16 | 0) >> 2] = $7 + 1 | 0;
    HEAP8[($4 + 11 | 0) >> 0] = HEAPU8[$7 >> 0] | 0;
    $11 = HEAP32[($4 + 12 | 0) >> 2] | 0;
    HEAP32[($4 + 12 | 0) >> 2] = $11 + 1 | 0;
    HEAP8[($4 + 10 | 0) >> 0] = HEAPU8[$11 >> 0] | 0;
    label$3 : {
     if ((HEAPU8[($4 + 11 | 0) >> 0] | 0) & 255 | 0) {
      break label$3
     }
     HEAP32[($4 + 28 | 0) >> 2] = ((HEAPU8[($4 + 11 | 0) >> 0] | 0) & 255 | 0) - ((HEAPU8[($4 + 10 | 0) >> 0] | 0) & 255 | 0) | 0;
     break label$1;
    }
    if (((HEAPU8[($4 + 11 | 0) >> 0] | 0) & 255 | 0 | 0) == ((HEAPU8[($4 + 10 | 0) >> 0] | 0) & 255 | 0 | 0) & 1 | 0) {
     continue label$2
    }
    break label$2;
   };
   HEAP32[($4 + 28 | 0) >> 2] = ((HEAPU8[($4 + 11 | 0) >> 0] | 0) & 255 | 0) - ((HEAPU8[($4 + 10 | 0) >> 0] | 0) & 255 | 0) | 0;
  }
  return HEAP32[($4 + 28 | 0) >> 2] | 0 | 0;
 }
 
 function $2() {
  return (($1(1024 | 0, 1072 | 0) | 0 | 0) != (0 | 0) ^ -1 | 0) & 1 | 0 | 0;
 }
 
 function $3($0_1, $1_1) {
  $0_1 = $0_1 | 0;
  $1_1 = $1_1 | 0;
  var $4 = 0, $12 = 0;
  $4 = global$0 - 16 | 0;
  HEAP32[($4 + 12 | 0) >> 2] = $0_1;
  HEAP32[($4 + 8 | 0) >> 2] = $1_1;
  label$1 : {
   if (!(HEAP32[($4 + 12 | 0) >> 2] | 0)) {
    break label$1
   }
   $12 = 24;
   HEAP32[($4 + 12 | 0) >> 2] = (HEAP32[($4 + 12 | 0) >> 2] | 0) ^ (((HEAPU8[((4 - ((HEAP32[($4 + 8 | 0) >> 2] | 0 | 0) % (5 | 0) | 0) | 0) + 1067 | 0) >> 0] | 0) << $12 | 0) >> $12 | 0) | 0;
  }
  HEAP8[((HEAP32[($4 + 8 | 0) >> 2] | 0) + 1072 | 0) >> 0] = HEAP32[($4 + 12 | 0) >> 2] | 0;
  return;
 }
 
 bufferView = HEAPU8;
 initActiveSegments(env);
 function __wasm_memory_size() {
  return buffer.byteLength / 65536 | 0;
 }
 
 function __wasm_memory_grow(pagesToAdd) {
  pagesToAdd = pagesToAdd | 0;
  var oldPages = __wasm_memory_size() | 0;
  var newPages = oldPages + pagesToAdd | 0;
  if ((oldPages < newPages) && (newPages < 65536)) {
   var newBuffer = new ArrayBuffer(Math_imul(newPages, 65536));
   var newHEAP8 = new Int8Array(newBuffer);
   newHEAP8.set(HEAP8);
   HEAP8 = new Int8Array(newBuffer);
   HEAP16 = new Int16Array(newBuffer);
   HEAP32 = new Int32Array(newBuffer);
   HEAPU8 = new Uint8Array(newBuffer);
   HEAPU16 = new Uint16Array(newBuffer);
   HEAPU32 = new Uint32Array(newBuffer);
   HEAPF32 = new Float32Array(newBuffer);
   HEAPF64 = new Float64Array(newBuffer);
   buffer = newBuffer;
   bufferView = HEAPU8;
  }
  return oldPages;
 }
 
 return {
  "memory": Object.create(Object.prototype, {
   "grow": {
    "value": __wasm_memory_grow
   }, 
   "buffer": {
    "get": function () {
     return buffer;
    }
    
   }
  }), 
  "__wasm_call_ctors": $0, 
  "strcmp": $1, 
  "check_flag": $2, 
  "input": global$1, 
  "copy_char": $3, 
  "key": global$2, 
  "__dso_handle": global$3, 
  "__data_end": global$4, 
  "__global_base": global$5, 
  "__heap_base": global$6, 
  "__memory_base": global$7, 
  "__table_base": global$8
 };
}

var retasmFunc = asmFunc(  { abort: function() { throw new Error('abort'); }
  });
export var memory = retasmFunc.memory;
export var __wasm_call_ctors = retasmFunc.__wasm_call_ctors;
export var strcmp = retasmFunc.strcmp;
export var check_flag = retasmFunc.check_flag;
export var copy_char = retasmFunc.copy_char;
