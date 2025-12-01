#! /opt/homebrew/bin/python3

import re

if __name__ == "__main__":
    hits = 0
    with open("input.txt") as file:
        pos = 50
        for line in file.readlines():
            distance = int(re.findall(r"\d+", line)[0])
            direction = line[0]
            if direction == "L":
                pos = (pos - distance) % 100
            elif direction == "R":
                pos = (pos + distance) % 100

            if pos == 0:
                hits += 1

    print(hits)
            