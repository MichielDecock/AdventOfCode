from collections import defaultdict

def spiral(goal):
    grid = defaultdict(int)
    directions = [(1, 0), (0, 1), (-1, 0), (0, -1)]
    neighbors = [(-1, 1), (0, 1), (1, 1), (-1, 0), (1, 0), (-1, -1), (0, -1), (1, -1)]

    x = 0
    y = 0
    grid[(x, y)] = 1

    d = 0
    step = 1

    while True:
        for _ in range(2):
            dx, dy = directions[d % 4]
            for _ in range(step):
                x += dx
                y += dy

                value = sum(grid[(x + nx, y + ny)] for nx , ny in neighbors)
                grid[(x, y)] = value
                if value > goal:
                    return value

            d += 1
        step += 1
    

if __name__ == "__main__":
    print(spiral(289326))