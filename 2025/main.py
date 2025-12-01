#! /opt/homebrew/bin/python3

import re
import math

if __name__ == "__main__":
    hits = 0
    with open("input.txt") as file:
        pos = 50
        for line in file.readlines():
            distance = int(re.findall(r"\d+", line)[0])
            direction = line[0]

            wasZero = pos == 0

            if direction == "L":
                pos -= distance
            elif direction == "R":
                pos += distance
            
            if pos == 0:
                if not wasZero:
                    hits += 1
                continue

            if pos > 99:
                while pos > 99:
                    pos -= 100
                    hits += 1
                continue

            if pos < 0:
                while pos < 0:
                    pos += 100
                    hits += 1
                if wasZero:
                    hits -= 1
            
            if pos == 0:
                hits += 1

    print(hits)
            