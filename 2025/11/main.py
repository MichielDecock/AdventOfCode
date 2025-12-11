#!/opt/homebrew/bin/python3

def parse(file):
    with open(file) as f:
        m = {}
        for l in [l.strip() for l in f.readlines()]:
            segments = l.split(':')
            v = []
            for s in [segments[0],  *segments[1].split(' ')]:
                if s == '':
                    continue
                v.append(s)
            m[v[0]] = v[1:]
        return m
    
def iterate(m):
    posIn = 'you'
    posOut = 'out'

    queue = [posIn]
    paths = 0

    while len(queue) > 0:
        p = queue.pop(0)

        if p == posOut:
            paths += 1
            continue

        for el in m[p]:
            queue.append(el)

    return paths

if __name__ == '__main__':
    m = parse('input')

    print(iterate(m))