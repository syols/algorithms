import functools


def solution():
    n = int(input())
    v = input()
    if n == 0:
        return []

    res = v.split(' ')
    res = sorted(res, key=functools.cmp_to_key(compare))
    return res


def compare(a, b):
    if int(a+b) > int(b + a):
        return -1
    return 1


if __name__ == '__main__':
    print(''.join(list(map(str, solution()))))