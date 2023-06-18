
- Linked (.so) compilation works fine, but requires a shit ton of libraries.
- Static compilation is a major headache, because the following libraries are not available on Arch as a `.a` file:

```bash
/usr/lib/go/pkg/tool/linux_amd64/link: running g++ failed: exit status 1
/sbin/ld: cannot find -lwebkit2gtk-4.0: No such file or directory
/sbin/ld: cannot find -lgtk-3: No such file or directory
/sbin/ld: cannot find -lgdk-3: No such file or directory
/sbin/ld: cannot find -lpangocairo-1.0: No such file or directory
/sbin/ld: cannot find -lpango-1.0: No such file or directory
/sbin/ld: cannot find -lharfbuzz: No such file or directory
/sbin/ld: cannot find -latk-1.0: No such file or directory
/sbin/ld: cannot find -lcairo-gobject: No such file or directory
/sbin/ld: cannot find -lcairo: No such file or directory
/sbin/ld: cannot find -lgdk_pixbuf-2.0: No such file or directory
/sbin/ld: cannot find -lsoup-2.4: No such file or directory
/sbin/ld: cannot find -ljavascriptcoregtk-4.0: No such file or directory
collect2: error: ld returned 1 exit status
```

