#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>


volatile int i;

static void __attribute__((noinline)) do_calc()  {
	static unsigned int called = 0;
	if (called++ % 1000 == 0) {
		printf("I am stuck. Please help.\n");
		sleep(1);
	}
}


static unsigned char calculate_answer(void) {
	for (i = 0; i < 10; ) {
		do_calc();
	}

	int res = (0xA << 4) | (0xF0 >> 4);
	return ~res & 0xFF;
}

int main(int argc, char **argv) {

	unsigned char answer = calculate_answer();
	printf("The answer is: %c\n", answer);
}
