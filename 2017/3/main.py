import math

def distance(cell):
    root = int(math.sqrt(cell))
    if root % 2 == 0:
        root -= 1

    remainder = cell % (root**2)
    ring = (root - 1) // 2

    maxDir = root + 1
    if remainder <= maxDir:
        up = abs(remainder - maxDir / 2)
        right = ring + 1
        return up + right
    elif remainder <= 2 * maxDir:
        up = maxDir / 2
        left = abs(remainder - 3 / 2 * maxDir)
        return up + left
    # rest is not relevant for this input

    return remainder

if __name__ == "__main__":
    print(distance(289326))