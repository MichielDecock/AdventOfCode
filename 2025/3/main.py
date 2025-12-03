#! /opt/homebrew/bin/python3
from itertools import islice

if __name__ == "__main__":
    sum = 0
    with open("input") as file:
        for line in (lines:= [l.strip() for l in file.readlines()]):
            n1 = line.index(max(line[:len(line) - 1]))
            sub = line[n1 + 1:]
            n2 = sub.index(max(sub[:len(sub)]))
            sum += int(line[n1]+sub[n2])

    print(sum)