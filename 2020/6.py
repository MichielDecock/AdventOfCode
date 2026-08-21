import os


def parse(filename):
    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        curLine = set()
        skip = False
        for line in [line.strip('\n') for line in file]:
            if not line:
                if curLine:
                    lines.append(curLine)
                curLine = set()
                skip = False
                continue

            if skip:
                continue

            if not curLine:
                curLine = set(line)
                continue

            curLine = curLine.intersection(line)
            if not curLine:
                skip = True

    return lines

if __name__ == "__main__":
    lines = parse('input')
    print(sum(len(l) for l in lines))
