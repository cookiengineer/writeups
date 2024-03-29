/* Automatically generated by wasm2c */
#ifndef TEST_H_GENERATED_
#define TEST_H_GENERATED_

#include <stdint.h>

#include "wasm-rt.h"

/* TODO(binji): only use stdint.h types in header */
#ifndef WASM_RT_CORE_TYPES_DEFINED
#define WASM_RT_CORE_TYPES_DEFINED
typedef uint8_t u8;
typedef int8_t s8;
typedef uint16_t u16;
typedef int16_t s16;
typedef uint32_t u32;
typedef int32_t s32;
typedef uint64_t u64;
typedef int64_t s64;
typedef float f32;
typedef double f64;
#endif

#ifdef __cplusplus
extern "C" {
#endif

typedef struct Z_Z5AoRd23o0wd_instance_t {
  u32 w2c_g0;
  u32 w2c_input;
  u32 w2c___dso_handle;
  u32 w2c___data_end;
  u32 w2c___global_base;
  u32 w2c___heap_base;
  u32 w2c___memory_base;
  u32 w2c___table_base;
  wasm_rt_memory_t w2c_memory;
  wasm_rt_funcref_table_t w2c_T0;
} Z_Z5AoRd23o0wd_instance_t;

void Z_Z5AoRd23o0wd_init_module(void);
void Z_Z5AoRd23o0wd_instantiate(Z_Z5AoRd23o0wd_instance_t*);
void Z_Z5AoRd23o0wd_free(Z_Z5AoRd23o0wd_instance_t*);

/* export: 'memory' */
wasm_rt_memory_t* Z_Z5AoRd23o0wdZ_memory(Z_Z5AoRd23o0wd_instance_t* instance);

/* export: '__wasm_call_ctors' */
void Z_Z5AoRd23o0wdZ___wasm_call_ctors(Z_Z5AoRd23o0wd_instance_t*);

/* export: 'strcmp' */
u32 Z_Z5AoRd23o0wdZ_strcmp(Z_Z5AoRd23o0wd_instance_t*, u32, u32);

/* export: 'check_flag' */
u32 Z_Z5AoRd23o0wdZ_check_flag(Z_Z5AoRd23o0wd_instance_t*);

/* export: 'input' */
u32* Z_Z5AoRd23o0wdZ_input(Z_Z5AoRd23o0wd_instance_t* instance);

/* export: 'copy_char' */
void Z_Z5AoRd23o0wdZ_copy_char(Z_Z5AoRd23o0wd_instance_t*, u32, u32);

/* export: '__dso_handle' */
u32* Z_Z5AoRd23o0wdZ___dso_handle(Z_Z5AoRd23o0wd_instance_t* instance);

/* export: '__data_end' */
u32* Z_Z5AoRd23o0wdZ___data_end(Z_Z5AoRd23o0wd_instance_t* instance);

/* export: '__global_base' */
u32* Z_Z5AoRd23o0wdZ___global_base(Z_Z5AoRd23o0wd_instance_t* instance);

/* export: '__heap_base' */
u32* Z_Z5AoRd23o0wdZ___heap_base(Z_Z5AoRd23o0wd_instance_t* instance);

/* export: '__memory_base' */
u32* Z_Z5AoRd23o0wdZ___memory_base(Z_Z5AoRd23o0wd_instance_t* instance);

/* export: '__table_base' */
u32* Z_Z5AoRd23o0wdZ___table_base(Z_Z5AoRd23o0wd_instance_t* instance);

#ifdef __cplusplus
}
#endif

#endif  /* TEST_H_GENERATED_ */
