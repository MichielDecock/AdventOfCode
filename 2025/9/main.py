#! /opt/homebrew/bin/python3

def getTiles(filename):
    map = []
    with open(filename) as file:
        for line in [l.strip() for l in file.readlines()]:
            map.append([int(s)for s in line.split(',')])

    return map

if __name__ == "__main__":
    S = 0
    tiles = getTiles('input')
    tiles.sort(key=lambda t: (t[0], t[1]))

    for i, tile in enumerate(tiles):
        for other in tiles[i + 1:]:
            s = (abs(tile[0]-other[0]) + 1) * (abs(tile[1]-other[1]) + 1)
            S = max(s, S)

    print(S)