import os


_LAST_ID = 0
_ID_MAP = dict()
_TARGET = 'shiny gold'


def parse(filename):
    global _LAST_ID
    lines = dict()
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for line in [line.strip('\n') for line in file]:
            line = line.replace('bags',' ')
            line = line.replace('bag',' ')
            line = line.replace('no other',' ')
            line = line.replace('.',' ')
            line = ''.join(c for c in line if not c.isdigit())
            [key, _values] = [l.strip() for l in line.split('contain')]
            values = [v.strip() for v in _values.split(',') if v.strip()]
            if key not in _ID_MAP.keys():
                _ID_MAP[key] = _LAST_ID
                _LAST_ID += 1
            lines[key] = values

    return lines

def tagDirectContainment(lines):
    return [_ID_MAP[key] for key in lines if _TARGET in lines[key]]

def containsTarget(key, lines, holders):
    if _ID_MAP[key] in holders:
        return True

    if _TARGET in lines[key]:
        holders.append(_ID_MAP[key])
        return True

    for v in lines[key]:
        if containsTarget(v, lines, holders):
            return True

    return False


if __name__ == "__main__":
    lines = parse('input')
    holders = tagDirectContainment(lines)
    print(sum(1 for key in lines if containsTarget(key, lines, holders)))
