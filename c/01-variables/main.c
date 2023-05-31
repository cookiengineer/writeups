#include <stdio.h>

int main() {

	int number;

	int first, second = 123;
	double twice = 13.37;
	float half = 0.5f;
	char character = 'T';

	printf("first is %d\n", first);
	printf("second is %d\n", second);
	printf("twice is %.2lf\n", twice);
	printf("half is %f\n", half);
	printf("character is %c\n", character);

	printf("sizeof int is %zu bytes\n", sizeof(int));
	printf("sizeof double is %zu bytes\n", sizeof(twice));

	number = 0;

	for (int i = 0; i < 20; i++) {

		printf("number is now %d\n", number);
		number++;

	}

	return 0;

}
