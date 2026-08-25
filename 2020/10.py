import os

_MAX_DIFF = 3
_START_VOLTAGE = 0

def parse(filename):
    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        return [int(line.strip('\n')) for line in file]

    return lines

def pathCounts(voltages):
    voltages.append(_START_VOLTAGE)
    voltages.append(max(voltages) + _MAX_DIFF)
    voltages.sort()

    n = len(voltages)
    counts = [0] * n
    counts[0] = 1
    for j in range(1, n):
        for i in range(j):
            if voltages[j] - voltages[i] <= _MAX_DIFF:
                counts[j] += counts[i]
    return counts[-1]


if __name__ == "__main__":
    lines = parse('input')
    counts = pathCounts(lines)
    print(counts)
