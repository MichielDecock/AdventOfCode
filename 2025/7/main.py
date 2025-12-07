#! /opt/homebrew/bin/python3

import copy

def findStart(line):
    return line.find('S')

if __name__ == "__main__":
    sum = 0

    map = []

    with open("input") as file:
        lines = [l.strip() for l in file.readlines()]
        for line in lines:
            map.append([c for c in line])

        pos = [findStart(lines[0])]

        for r  in range(1, len(lines)):
            newPos = copy.deepcopy(pos)
            for p in pos:
                if map[r][p] != '^':
                    continue

                sum += 1

                newPos.remove(p)
                if p > 0:
                    newPos.append(p - 1)
                if p < len(lines[0]) - 1:
                    newPos.append(p + 1)

            pos = list(set(copy.deepcopy(newPos)))

    print(sum)