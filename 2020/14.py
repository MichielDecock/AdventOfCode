import os
from itertools import product

def getFloatingMasks(bits):
        masks = []

        for comb in product([0,1], repeat=len(bits)):
            mask = 0
            for bit, pos in zip(comb, bits):
                if bit == 1:
                    mask += 2**pos
            masks.append(mask)

        return masks

def parse(filename):
    out = []
    masks = []
    mem = []

    def flush():
        nonlocal masks, mem
        if len(masks) > 0:
            out.append((masks, mem))
            masks = []
            mem = []

    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for line in [l.split('=') for l in [line.strip('\n') for line in file]]:
            t, v = [el.strip() for el in line]

            if t == 'mask':
                flush()
                vStrip = v.strip()
                idx = sorted([len(vStrip) - (i + 1) for i, c in enumerate(vStrip) if c == 'X'])
                masks.append(int(vStrip.replace('X', '0'), 2))
                masks.extend(idx)
                continue

            address = int(t.split('[')[1].strip().split(']')[0].strip())
            value = int(v.strip())
            mem.append((address, value))

        flush()

    return out

def copyBits(dest, src, bits):
    for pos in bits:
        mask = 1 << pos
        dest &= ~mask
        bit = (src & mask)
        dest |= bit
    return dest

if __name__ == "__main__":
    mem = {}
    instructions = parse('input')

    for m, inst in instructions:
        mask = m[0]
        bits = m[1:]

        fMasks = getFloatingMasks(bits)

        for address, value in inst:
            address |= mask
            for fMask in fMasks:
                newAddress = copyBits(address, fMask, bits)
                mem[newAddress] = value

    print(sum(mem.values()))
