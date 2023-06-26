defmodule Calculator do
    @moduledoc """
    Provide mathematical calculation from user input
    """

    @doc """
    Get user input for the numbers to be calculated
    """
    def get_user_input do
        number = IO.gets("Enter a number: ")
        if number not in [""] do
            elem(Float.parse(number), 0)
        else
            nil
        end
        
    end

    @doc """
    Get the user input operator for our calculation
    """
    def get_operator do
        IO.gets("Enter an operator sign [+ - / * % ^]: ")
        |> String.trim()
    end

    @doc """
    Calculate x and y according to the user choice operant
    """
    def calculate(x, "+", y), do: x + y
    def calculate(x, "-", y), do: x - y
    def calculate(x, "*", y), do: x * y
    def calculate(x, "/", y), do: x / y
    def calculate(x, "^", y), do: :math.pow(x, y)
    def calculate(x, "%", y), do: rem x, y
    def calculate(_, _, _), do: IO.puts("Invalid operator")
end

defmodule Main do
    def main do
        IO.puts("\t==============================\n")
        IO.puts("\t Welcome to ElixirCalculator! \n")
        IO.puts("\t==============================\n\n")
        x       = Calculator.get_user_input 
        operant = Calculator.get_operator
        y       = Calculator.get_user_input 

        result = Calculator.calculate(x, operant, y)
        IO.puts("#{x} #{operant} #{y} = #{result}")
    end
end
