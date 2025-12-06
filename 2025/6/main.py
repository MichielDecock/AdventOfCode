#! /opt/homebrew/bin/python3

from functools import reduce
from operator import mul
from operator import add

if __name__ == "__main__":
    sum = 0

    input = []

    with open("test") as file:
        for line in (lines:= [l.strip() for l in file.readlines()]):
            input.append(line.split())

        input = list(map(list, zip(*input)))
        for line in input:
            if line[-1] == '*':
                sum += reduce(mul, [int(l) for l in line[:-1]])
            if line[-1] == '+':
                sum += reduce(add, [int(l) for l in line[:-1]])

    print(sum)