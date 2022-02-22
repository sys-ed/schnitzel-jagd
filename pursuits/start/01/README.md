# Question

**What is the syscall number base 10 for flock(2) on ARM64? Write down the number**

# Background

A syscall (or system call) is how an userspace application requests a service
from the operating system. The syscall is the _public_ API of the Linux kernel
with an high emphasis on not breaking existing users.

An example of a syscall is
[get_robust_list](https://elixir.bootlin.com/linux/v5.15.12/source/kernel/futex.c#L3584). The
way a syscall is invoked is both CPU architecture and ABI specific. The classic way on x86 is
to put the syscall number into a register and then issue an interrupt. The interrupt handler
of the kernel. There are various optimizations in use. One is the usage of specific instructions
(syscall/sysret) to avoid the cost of general purpose interrupt handling and VDSO (where some
kernel memory that includes functions and data is mapped into userspace).


# Pointers

* https://en.wikibooks.org/wiki/X86_Assembly/Interfacing_with_Linux
* https://wiki.osdev.org/SYSENTER




