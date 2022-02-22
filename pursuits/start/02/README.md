# Question

Please compile stuck.c and execute and it will reveal an answer on stdout. Oh no! It
seems got stuck. Can you make it finish **without** re-compiling and restarting the
process?

# Background

The application is stuck in a for loop. The code loads the value from memory into the
register and then compares the value in the register with a fixed number (immediate).
Can we change the variable at this address in memory? Can we change the code? Can we
change the register value before comparison?

# Pointers

* GDB and LLDB are two debuggers available on GNU/Linux.
* https://lldb.llvm.org/use/map.html
