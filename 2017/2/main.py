def parse(file):
    with open(file) as f:
        return [sorted([int(x) for x in l.strip().split('\t')], reverse=True) for l in f.readlines()]
    
def checkSum(data):
    return sum([max(d) - min(d) for d in data])

if __name__ == "__main__":
    data = parse('input')
    result = checkSum(data)
    print(result)