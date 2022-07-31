def solution():
    n = int(input())
    v = input()
    if n == 0:
        return []
    values = list(map(int, v.split(' ')))

    is_changed = True
    length = len(values)
    is_ok = True
    while is_changed:
        is_changed = False
        for i in range(0, length - 1):
            if values[i] > values[i + 1]:
                is_ok = False
                values[i + 1], values[i] = values[i], values[i + 1]
                is_changed = True
        if is_changed:
            print(" ".join(list(map(str, values))))
    if is_ok:
        print(" ".join(list(map(str, values))))
    return values


if __name__ == '__main__':
    solution()
