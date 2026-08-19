import os

_TARGET = 2020

def getNumbers(filename: str) -> list[int]:
    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for line in file:
            lines.append(int(line))
    return lines

def findCombination(numbers: list[int]) -> list[int] | None:
    for idx, x in enumerate(numbers):
        for idy, y in enumerate(numbers):
            if idx == idy:
                continue

            for idz, z in enumerate(numbers):
                if idx == idz or idy == idz:
                    continue

                if x + y + z == _TARGET:
                    return [x, y, z]
    return None


if __name__ == "__main__":
    numbers = getNumbers('input')
    if combination := findCombination(numbers):
        print(combination[0] * combination[1] * combination[2])
