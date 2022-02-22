Get an idea of how the C code maps to assembly.

# Decompile

Use _objdump -d stuck_ to get a listing of all the code in the binary. It reads
like this (on macos):

```shell
0000000100003ec8 <_do_calc>:
100003ec8: fd 7b bf a9 	stp	x29, x30, [sp, #-16]!
100003ecc: fd 03 00 91 	mov	x29, sp
100003ed0: 28 00 00 b0 	adrp	x8, 0x100008000 <_do_calc+0x1c>
100003ed4: 09 21 40 b9 	ldr	w9, [x8, #32]
100003ed8: 2a 05 00 11 	add	w10, w9, #1
100003edc: 0a 21 00 b9 	str	w10, [x8, #32]
100003ee0: a8 1a 8f 52 	mov	w8, #30933
100003ee4: 28 dd a4 72 	movk	w8, #9961, lsl #16
100003ee8: 28 7d 08 1b 	mul	w8, w9, w8
100003eec: 08 0d 88 13 	ror	w8, w8, #3
100003ef0: e9 26 91 52 	mov	w9, #35127
100003ef4: 29 08 a0 72 	movk	w9, #65, lsl #16
100003ef8: 1f 01 09 6b 	cmp	w8, w9
100003efc: 69 00 00 54 	b.ls	0x100003f08 <_do_calc+0x40>
100003f00: fd 7b c1 a8 	ldp	x29, x30, [sp], #16
100003f04: c0 03 5f d6 	ret
100003f08: 60 04 00 50 	adr	x0, #142
100003f0c: 1f 20 03 d5 	nop
100003f10: 07 00 00 94 	bl	0x100003f2c <dyld_stub_binder+0x100003f2c>
100003f14: 20 00 80 52 	mov	w0, #1
100003f18: fd 7b c1 a8 	ldp	x29, x30, [sp], #16
100003f1c: 07 00 00 14 	b	0x100003f38 <dyld_stub_binder+0x100003f38>
```
