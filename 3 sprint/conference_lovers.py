import functools


def solution():
    n = int(input())
    v = input()
    if n == 0:
        return []
    k = int(input())

    res = v.split(' ')
    ids = {}
    for e in res:
        if e not in ids:
            ids[e] = 1
        else:
            ids[e] += 1
    res = [k for k, _ in sorted(ids.items(), key=lambda item: item[1], reverse=True)]
    return res[0:k]


if __name__ == '__main__':
    print(' '.join(list(map(str, solution()))))