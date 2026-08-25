import os

_MAX_DIFF = 3
_START_VOLTAGE = 0

def parse(filename):
    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        return [int(line.strip('\n')) for line in file]

    return lines

def addToDistro(value, distro):
    if value in distro:
        distro[value] += 1
    else:
        distro[value] = 1

def makeDistro(voltages):
    voltages.append(_START_VOLTAGE)
    voltages.sort()
    distro = dict()
    for i in range(len(voltages) - 1):
        diff = voltages[i + 1] - voltages[i]
        if (diff > _MAX_DIFF):
            print("error")

        addToDistro(diff, distro)

    addToDistro(3, distro)
        
    return distro


if __name__ == "__main__":
    lines = parse('input')
    distro = makeDistro(lines)
    print(distro[1] * distro[3])
