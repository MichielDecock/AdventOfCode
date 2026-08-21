import os


def parse(filename):
    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        curLine = ''
        for line in [line.strip('\n') for line in file]:
            if not line:
                lines.append(''.join(set(curLine)))
                curLine = ''
            else:
                curLine += line

    return lines

if __name__ == "__main__":
    lines = parse('input')
    print(sum(len(l) for l in lines))
