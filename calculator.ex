defmodule Calculator do
    @moduledoc """
    Provide mathematical calclation from user input
    """

    @doc """
    Get user input for the numbers to be calculated
    """
    def get_user_input do
        IO.gets("Enter a number: ") 
        |> String.trim() 
        |> String.to_integer()
    end

    @doc """
    Get the user input operator for our calculation
    """
    def get_operator do
        IO.gets("Enter an operator sign [+ - / * % ^]: ") |> String.trim()
    end

    @doc """
    Calculate x and y according to the user choice operant
    """
    def calculator(x, operant, y) do
        case {x, operant, y} do
            {x, "+", y} -> x + y
            {x, "-", y} -> x - y
            {x, "*", y} -> x * y
            {x, "/", y} -> x / y
            {x, "^", y} -> :math.pow(x, y)
            {x, "%", y} -> rem x, y
            _ -> IO.puts("Invalid operator")
        end
    end
end

defmodule Main do
    def main do
        IO.puts("\t==============================\n")
        IO.puts("\t Welcome to ElixirCalculator! \n")
        IO.puts("\t==============================\n\n")

        x       = Calculator.get_user_input 
        operant = Calculator.get_operator
        y       = Calculator.get_user_input 

        result = Calculator.calculator(x, operant, y)
        IO.puts("#{x} #{operant} #{y} = #{result}")
    end
end
