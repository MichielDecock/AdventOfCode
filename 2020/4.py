import os

_KEYS = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']
_OPTIONAL = 'cid'

def parse(filename: str):
    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        keys = []
        for line in [line.strip('\n') for line in file]:
            if not line:
                lines.append(keys)
                keys = []
                continue

            keys += [l for l in [l.split(':')[0] for l in line.split(' ')] if l != _OPTIONAL]

        if not keys:
            lines.append(keys)

    return lines

def isValid(passport):
    return len(passport) == len(_KEYS)

if __name__ == "__main__":
    passports = parse('input')
    print(sum(1 for p in passports if isValid(p)))
