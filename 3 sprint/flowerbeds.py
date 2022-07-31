def solution():
    count = int(input())
    data = []
    if count == 0:
        return []
    for i in range(count):
        values = sorted(list(map(int, input().split(' '))))
        data.append((values[0], values[1]))
    data = (sorted(list(set(data)), key=lambda t: t[0]))

    res = [data[0]]
    for a, b in data[1:]:
        la, lb = res[-1]
        if a in range(la, lb + 1) and b >= lb:
            res[-1] = (la, b)
        elif a > la and b > lb:
            res.append((a, b))

    return res


if __name__ == '__main__':
    for e in solution():
        print(" ".join(map(str, e)))
