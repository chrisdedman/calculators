defmodule Calculator do
    @moduledoc """
    Provide mathematical calculation from user input
    """
    @operators ~w(+ - * / ^ %)
    
    @doc """
    Calculate x and y according to the user choice operant
    """
    def calculate(x, operator, y) when is_number(x) and is_number(y) and operator in @operators,
    do: execute(x, operator, y)
    # handles invalid operator
    def calculate(_, operator, _) when operator not in @operators, do: {:error, "Invalid operator"}
    def calculate(_, _, _), do: {:error, "Invalid arguments! Ex. `calculate(5, +, 5)`"}
     # executes the calculation
    defp execute(x, "+", y), do: x + y
    defp execute(x, "-", y), do: x - y
    defp execute(x, "*", y), do: x * y
    defp execute(_, "/", 0), do: {:error, "Division by zero"}
    defp execute(x, "/", y), do: x / y
    defp execute(x, "^", y), do: :math.pow(x, y)
    defp execute(x, "%", y), do: rem(x, y)
end

defmodule Main do
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

    def run do
        IO.puts("\t==============================\n")
        IO.puts("\t Welcome to ElixirCalculator! \n")
        IO.puts("\t==============================\n\n")
        x        = get_user_input()
        operator = get_operator()
        y        = get_user_input()

        case Calculator.calculate(x, operator, y) do
            {:error, message} -> IO.puts "Error: #{message}"
            result -> IO.puts "Result: #{x} #{operator} #{y} = #{result}"
        end
    end
end
