#! /opt/homebrew/bin/python3

from functools import reduce
from operator import mul
from operator import add

if __name__ == "__main__":
    sum = 0

    input = []

    with open("input") as file:
        for line in (lines := [l for l in file.readlines()]):
            input.append([c for c in line])

        operator = ''
        operands = []
        for line in (input := list(map(list, zip(*input)))):
            s = ''.join(line).strip()

            if s == '':
                if operator == '*':
                    sum += reduce(mul, operands)
                elif operator == '+':
                    sum += reduce(add, operands)
                continue

            if s.endswith(('*','+')):
                operator = s[-1]
                operands = [int(s[:-1])]
                continue

            operands.append(int(s))

        if operator == '*':
            sum += reduce(mul, operands)
        elif operator == '+':
            sum += reduce(add, operands)

    print(sum)