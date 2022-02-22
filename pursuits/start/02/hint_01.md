GDB and LLDB allow to attach to a running process.

```shell
ps xua | grep stuck
lldb -p <some_pid>
```

Search for _info registers_ in https://lldb.llvm.org/use/map.html
