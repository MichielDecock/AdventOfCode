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
    
def iterate(m, posIn, posOut):
    queue = [[posIn, 1]]
    numPaths = 0

    while len(queue) > 0:
        e = queue.pop(0)
        p = e[0]
        count = e[1]

        if p == posOut:
            numPaths += count
            continue

        if p == 'out':
            continue

        for el in m[p]:
            found = False
            for index, s in enumerate(queue):
                if s[0] == el:
                    queue[index][1] += count
                    found = True
                    break

            if not found:
                queue.append([el, count])

    return numPaths

if __name__ == '__main__':
    m = parse('input')

    fft2dac = iterate(m, 'fft', 'dac')
    fft2out = iterate(m, 'fft', 'out')
    dac2out = iterate(m, 'dac', 'out')
    dac2fft = iterate(m, 'dac', 'fft')

    print(iterate(m, 'svr', 'fft') * fft2dac * dac2out + iterate(m, 'svr', 'dac') * dac2fft * fft2out)
