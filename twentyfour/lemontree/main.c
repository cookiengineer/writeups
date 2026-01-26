
/* WARNING: Function: __x86.get_pc_thunk.bx replaced with injection: get_pc_thunk_bx */
/* WARNING: Globals starting with '_' overlap smaller symbols at the same address */

undefined4 main(int param_1)
{
  undefined4 uVar1;
  undefined auStack_90 [12];
  char local_84;
  char local_83;
  char local_82;
  char local_81;
  undefined4 *local_14;
  
  local_14 = &param_1;
  fflush(_stdout);
  if (param_1 < 0x29) {
    puts("I\'m sittin\' here in the boring...");
    fflush(_stdout);
    __isoc99_scanf(&DAT_0804a06b,&local_84);
    if ((((local_84 == 'r') && (local_83 == 'o')) && (local_82 == 'o')) && (local_81 == 'm')) {
      puts("It\'s just another rainy Sunday...");
      fflush(_stdout);
      __isoc99_scanf(&DAT_0804a06b,&local_84);
      if (((local_84 == 'a') && (local_83 == 'f')) && ((local_82 == 't' && (local_81 == 'e')))) {
        printf("I\'m wastin\' my time, I got nothin\' to do");
        printf("I\'m hangin\' around, I\'m waitin\' for you");
        check();
      }
    }
    return 0;
  }
                    /* WARNING: Could not recover jumptable at 0x0804924e. Too many branches */
                    /* WARNING: Treating indirect jump as call */
  uVar1 = (*(code *)auStack_90)();
  return uVar1;
}

