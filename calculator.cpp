#include <iostream>
#include <iomanip>

// function that ask the user an integer
int getUserInput()
{
	std::cout << "Enter an integer: ";
	int numInputed;
	std::cin >> numInputed;

	return numInputed;
}

// function that ask the user for a mathematical operator sign.
char getUserMathSign()
{
	std::cout << "Enter an operator sign: ";
	char op;
	std::cin >> op;

	return op;

}

// function that calculate x and y (add, Subt, Multi, Div) with the operator
int calculator( int x, char operant, int y )
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
		return 404;
	}
}

int main()
{
	std::cout << "=========================\n";
	std::cout << "Welcome to my Calculator!\n";
	std::cout << "=========================\n";

	tryAgain:
		
		std::cout << '\n';
	
		int firstValue = getUserInput();
		char operant = getUserMathSign();
		int secondValue = getUserInput();

		std::cout << '\n';

		int result = calculator(firstValue, operant, secondValue);

		if (result == 404)
			goto tryAgain;
		
		system("pause");

		return 0;
}
