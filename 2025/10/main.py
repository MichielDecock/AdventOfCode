#! /opt/homebrew/bin/python3

from sys import exit

def readInput(filename):
    contents = []
    with open(filename) as file:
        for line in [l.strip() for l in file.readlines()]:
            segments = line.split(' ')

            lights = sum([1 << i if s == '#' else 0 for i, s in enumerate(reversed(segments[0][1:-1]))])

            si = len(segments[0][1:-1]) - 1

            buttons = []
            for s in segments[1:-1]:
                t = 0
                for j in s[1:-1].split(','):
                    t += pow(2, (si - int(j)))
                buttons.append(t)

            contents.append([lights, buttons])

    return contents

def minCombinations(lights, buttons, index):
    queue = [(0, [])]
    visited = []

    while queue:
        cur, subset = queue.pop(0)
        if cur == lights:
            return len(subset)
        for s in buttons:
            nxt = cur ^ s
            if nxt not in visited:
                visited.append(nxt)
                queue.append((nxt, subset+[s]))

    exit(f"Didn't find a minimal set of combinations for light {index}")
    return 0
    


if __name__ == '__main__':
    contents = readInput('input')

    s = 0
    for index, c in enumerate(contents):
        s += minCombinations(c[0], c[1], index)

    print(s)
