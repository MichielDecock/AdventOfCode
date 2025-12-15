def parse(file):
    with open(file) as f:
        return [sorted([int(x) for x in l.strip().split('\t')], reverse=True) for l in f.readlines()]
    
def findMultiple(d):
    for n1 in d:
        for n2 in d:
            if n1 == n2:
                continue

            if max(n2, n1) % min(n2, n1) == 0:
                return max(n2, n1) // min(n2, n1)
    return 0
    
def checkSum(data):
    return sum(findMultiple(d) for d in data)

if __name__ == "__main__":
    data = parse('input')
    result = checkSum(data)
    print(result)