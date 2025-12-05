#! /opt/homebrew/bin/python3

if __name__ == "__main__":
    sum = 0

    fresh = []

    with open("input") as file:
        for line in (lines:= [l.strip() for l in file.readlines()]):

            if line == '':
                break

            r = [int(l) for l in  line.split('-')]

            if len(fresh) == 0:
                fresh.append(r)
                continue

            for prev in fresh:
                if prev[0] == -1:
                    continue

                if r[0] >= prev[0] and r[0] <= prev[1]:
                    r[0] = prev[1] + 1
                elif r[1] >= prev[0] and r[1] <= prev[1]:
                    r[1] = prev[0] - 1

                if r[0] > r[1]:
                    break

                if r[0] < prev[0] and r[1] > prev[1]:
                    prev[0] = -1

            if r[0] <= r[1]:
                fresh.append(r)

        for item in fresh:
            if item[0] == -1:
                continue

            sum += item[1] - item[0] + 1
            

    print(sum)