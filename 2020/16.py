import os
from itertools import groupby


def parseFile(filename):
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        lines = [line.strip('\n') for line in file]
        return [list(g) for k, g in groupby(lines, key=bool) if k]
    return []

def parseFields(fields):
    out = {}
    for f in fields:
        name = f.split(':')[0]
        ranges = [tuple([int(x) for x in  r.split('-')]) for r in f.split(':')[1].strip().split(' or ')]
        out[name] = ranges
    return out

def parseNearby(nearby):
    out = []
    for n in nearby[1:]:
        out.append([int(x) for x in n.split(',')])
    return out

def parseGroups(groups):
    fields = parseFields(groups[0])
    nearby = parseNearby(groups[2])

    return (fields, None, nearby)

def invalid(fields, nearby):
    out = []

    for n in nearby:
        d = {}
        for i in n:
            keys = []
            for key, value in fields.items():
                for v in value:
                    if i in range(v[0], v[1] + 1):
                        keys.append(key)
            d[i] = keys

        out.extend([key for key, value in d.items() if value == []])

    return out

if __name__ == "__main__":
    groups = parseFile('input')
    fields, _, nearby = parseGroups(groups)
    print(sum(invalid(fields, nearby)))
