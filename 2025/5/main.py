#! /opt/homebrew/bin/python3

if __name__ == "__main__":
    sum = 0

    fresh = []

    idMode = False

    with open("input") as file:
        for line in (lines:= [l.strip() for l in file.readlines()]):

            if line == '':
                idMode = True
                continue

            if not idMode:
                fresh.append([int(l) for l in  line.split('-')])
                continue

            id = int(line)
            
            for r in fresh:
                if r[0] <= id and id <= r[1]:
                    sum +=1
                    break
            

    print(sum)