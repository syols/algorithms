import math


def solution():
    pairs = []
    a = int(input())
    m = int(input())
    s = input()

    length = len(s)
    degree_list = []
    prefix_list = []

    dlast = 1
    degree_list.append(1)
    for d in range(1, length + 1):
        dlast = dlast * a % m
        degree_list.append(dlast)

    plast = 0
    prefix_list.append(0)
    for p in range(1, length + 1):
        plast = ((plast * a % m + ord(s[p - 1])) % m)
        prefix_list.append(plast)

    for _ in range(int(input())):
        v = input()
        f = v.split()
        left = int(f[0]) - 1
        right = int(f[1]) - 1
        r = prefix_list[right + 1] - prefix_list[left] * degree_list[right - left + 1]
        print(r % m)



if __name__ == '__main__':
    solution()