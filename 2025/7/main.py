#! /opt/homebrew/bin/python3

import copy

def findStart(line):
    return line.find('S')

if __name__ == "__main__":
    map = []

    with open("input") as file:
        lines = [l.strip() for l in file.readlines()]
        for line in lines:
            map.append([c for c in line])

        pos = [(findStart(lines[0]),1)]

        for r  in range(1, len(lines)):
            newPos = copy.deepcopy(pos)
            for p in pos:
                col = p[0]
                num = p[1]
                if map[r][col] != '^':
                    continue

                if col > 0:
                    newPos.append((col - 1,num))
                if col < len(lines[0]) - 1:
                    newPos.append((col + 1, num))
                newPos.remove(p)

            newPos2 = copy.deepcopy(newPos)
            for i, p in enumerate(newPos):
                count = p[1]
                for j, x in enumerate(newPos):
                    if i == j:
                        continue

                    if x[0] == p[0]:
                        count += x[1]
                newPos2[i] = (p[0], count)

            newPos2 = list(set(newPos2))
            pos = copy.deepcopy(newPos2)

    s = sum(p[1] for p in pos)

    print(s)