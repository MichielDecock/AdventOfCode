import os

_CUTOFF = 25

def parse(filename) -> list[int]:
    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for line in [line.strip('\n') for line in file]:
            lines.append(int(line))

    return lines

def findParts(preamble, target):
    for u in preamble:
        for v in preamble:
            if u != v and u + v == target:
                return True

    return False

def findProblem(lines):
    for pos in range(_CUTOFF, len(lines)):
        item = lines[pos]
        preamble = lines[pos -_CUTOFF: pos]

        if not findParts(preamble, item):
            return (item, pos)

    return (None, None)


if __name__ == "__main__":
    lines = parse('input')
    target, index = findProblem(lines)
    assert(target)
    assert(index)

    start = 0
    end = 1
    while True:
        if end > len(lines):
            break

        s = sum(lines[start:end])
        while s < target:
            s += lines[end]
            end += 1

        if s == target:
            print(min(lines[start:end]) + max(lines[start:end]))
            break

        while s > target:
            s -= lines[start]
            start += 1
