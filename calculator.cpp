#include <iostream>

int getUserInput() // function that ask the user an integer
{
	std::cout << "Enter an integer: ";
	int numInputed{ };
	std::cin >> numInputed;

	return numInputed;
}

char getUserMath_Sign() // function that ask the user for a mathematical operator sign.
{
	std::cout << "Enter an operator sign: ";
	char op{};
	std::cin >> op;

	return op;

}

int calculator(int x, char operant, int y ) // function that calculate x and y (add, Subt, Multi, Div) with the operator
{
	switch (operant)
	{
	case '+':
		std::cout << x << " + " << y << " = " << x + y << '\n';
		return x + y;
	case '*':
		std::cout << x << " x " << y << " = " << x * y << '\n';
		return x * y;
	case '-':
		std::cout << x << " - " << y << " = " << x - y << '\n';
		return x - y;
	case '/':
		std::cout << float(x) << " / " << float(y) << " = " << float(x) / float(y) << '\n'; // convert variable into float to get an accurate divided number.
		return x / y;
	case '%':
		std::cout << x << " modulo " << y << " = " << x % y << '\n';
		return x % y;
	default:
		std::cout << "Error, invalid operator!" << '\n';
		return 404; // Error code number, try again if error.
	}
}

int main()
{
	std::cout << "Welcome to my calculator program!\n";
	tryAgain: // this is a statement label
		
		std::cout << '\n';
	
		// get the first number from the user
		int firstValue{ getUserInput() };

		// get the mathematical operation sign from the user
		char operant{ getUserMath_Sign() };

		// get the second number from the user
		int secondValue{ getUserInput() };

		std::cout << '\n';

		// print the result
		int result{ calculator(firstValue, operant, secondValue) };
		if (result == 404) // if operator sign invalid, and generate code 404, try again.
			goto tryAgain;

		/* This next three will cause your program to wait
	  * for the user to press enter before continuing
	  * which will give you time to examine your program’s
	  * output before your IDE closes the console window.
	  */

		std::cin.clear(); // reset any error flags
		std::cin.ignore(std::numeric_limits<std::streamsize>::max(), '\n'); // ignore any characters in the input buffer until we find a newline
		std::cin.get(); // get one more char from the user (waits for user to press enter)

		return 0;
}