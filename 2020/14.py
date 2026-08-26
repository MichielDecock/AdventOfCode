import os


def parse(filename):
    out = []
    masks: tuple[int, int] | None = None
    mem = []

    def flush():
        nonlocal masks, mem
        if masks is not None:
            out.append((masks, mem))
            masks = None
            mem = []

    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for line in [l.split('=') for l in [line.strip('\n') for line in file]]:
            t, v = [el.strip() for el in line]

            if t == 'mask':
                flush()
                masks = (int(v.strip().replace('X', '0'), 2), int(v.strip().replace('X', '1'), 2))
                continue

            address = int(t.split('[')[1].strip().split(']')[0].strip())
            value = int(v.strip())
            mem.append((address, value))

        flush()

    return out
            


if __name__ == "__main__":
    mem = {}
    instructions = parse('input')
    for (mask1, mask2), inst in instructions:
        for address, value in inst:
            mem[address] = (value |  mask1) & mask2

    print(sum(mem.values()))
