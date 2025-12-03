#! /opt/homebrew/bin/python3
from itertools import islice

if __name__ == "__main__":
    sum = 0
    with open("input") as file:
        for line in (lines:= [l.strip() for l in file.readlines()]):
            res = ''
            for i in range(12):
                idx = line.index(max(line[:len(line) - (12 - i - 1)]))
                res += line[idx]
                line = line[idx + 1:]

            sum += int(res)

    print(sum)