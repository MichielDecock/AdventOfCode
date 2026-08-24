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


if __name__ == "__main__":
    lines = parse('input')

    for pos in range(_CUTOFF, len(lines)):
        item = lines[pos]
        preamble = lines[pos -_CUTOFF: pos]

        if not findParts(preamble, item):
            print(item)
            break
