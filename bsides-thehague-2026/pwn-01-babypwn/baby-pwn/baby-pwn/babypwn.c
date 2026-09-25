
undefined8 dothing(void)

{
  long in_FS_OFFSET;
  char local_68 [48];
  char local_38 [40];
  long local_10;
  
  local_10 = *(long *)(in_FS_OFFSET + 0x28);
  builtin_strncpy(local_68,"/bin/echo \'you failed\'",0x17);
  local_68[0x17] = '\0';
  local_68[0x18] = '\0';
  local_68[0x19] = '\0';
  local_68[0x1a] = '\0';
  local_68[0x1b] = '\0';
  local_68[0x1c] = '\0';
  local_68[0x1d] = '\0';
  local_68[0x1e] = '\0';
  local_68[0x1f] = '\0';
  local_68[0x20] = '\0';
  local_68[0x21] = '\0';
  local_68[0x22] = '\0';
  local_68[0x23] = '\0';
  local_68[0x24] = '\0';
  local_68[0x25] = '\0';
  local_68[0x26] = '\0';
  local_68[0x27] = '\0';
  local_38[0] = '\0';
  local_38[1] = '\0';
  local_38[2] = '\0';
  local_38[3] = '\0';
  local_38[4] = '\0';
  local_38[5] = '\0';
  local_38[6] = '\0';
  local_38[7] = '\0';
  local_38[8] = '\0';
  local_38[9] = '\0';
  local_38[10] = '\0';
  local_38[0xb] = '\0';
  local_38[0xc] = '\0';
  local_38[0xd] = '\0';
  local_38[0xe] = '\0';
  local_38[0xf] = '\0';
  local_38[0x10] = '\0';
  local_38[0x11] = '\0';
  local_38[0x12] = '\0';
  local_38[0x13] = '\0';
  local_38[0x14] = '\0';
  local_38[0x15] = '\0';
  local_38[0x16] = '\0';
  local_38[0x17] = '\0';
  local_38[0x18] = '\0';
  local_38[0x19] = '\0';
  local_38[0x1a] = '\0';
  local_38[0x1b] = '\0';
  local_38[0x1c] = '\0';
  local_38[0x1d] = '\0';
  local_38[0x1e] = '\0';
  local_38[0x1f] = '\0';
  local_38[0x20] = '\0';
  local_38[0x21] = '\0';
  local_38[0x22] = '\0';
  local_38[0x23] = '\0';
  local_38[0x24] = '\0';
  local_38[0x25] = '\0';
  local_38[0x26] = '\0';
  local_38[0x27] = '\0';
  fgets(local_38,0x78,stdin);
  system(local_68);
  if (local_10 != *(long *)(in_FS_OFFSET + 0x28)) {
                    /* WARNING: Subroutine does not return */
    __stack_chk_fail();
  }
  return 0;
}


