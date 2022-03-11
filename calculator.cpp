#include <iostream>
#include <iomanip>

int get_User_Input() // function that ask the user an integer
{
	std::cout << "Enter an integer: ";
	int numInputed;
	std::cin >> numInputed;

	return numInputed;
}

char get_User_Math_Sign() // function that ask the user for a mathematical operator sign.
{
	std::cout << "Enter an operator sign: ";
	char op;
	std::cin >> op;

	return op;

}

int calculator( int x, char operant, int y ) // function that calculate x and y (add, Subt, Multi, Div) with the operator
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
		std::cout << x << " / " << y << " = " << double(x) / double(y) << '\n'; // convert variable into float to get an accurate divided number.
		return (int)((double)x / (double)y);
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
	std::cout << "=========================\n";
	std::cout << "Welcome to my Calculator!\n";
	std::cout << "=========================\n";

	tryAgain: // this is a statement label
		
		std::cout << '\n';
	
		// get the first number from the user
		int firstValue = get_User_Input();

		// get the mathematical operation sign from the user
		char operant = get_User_Math_Sign();

		// get the second number from the user
		int secondValue = get_User_Input();

		std::cout << '\n';

		// print the result
		int result = calculator(firstValue, operant, secondValue);

		if (result == 404) // if operator sign invalid, and generate code 404, try again.
			goto tryAgain;
		
		// time pause, press enter before to close the app
		system("pause");

		return 0;
}
