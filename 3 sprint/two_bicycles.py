import math


def solution():
    n = int(input())
    v = input()
    if n == 0:
        return []
    k = int(input())
    days = list(map(int, v.split(' ')))
    return [get_payday(days, k), get_payday(days, 2*k)]


def get_payday(days, k):
    payday = -1
    left = 1
    right = len(days)
    while True:
        day = get_day(left, right)
        pay = days[day - 1]
        if pay >= k:
            right = day
        else:
            left = day
        if right - left <= 1:
            if days[left - 1] >= k:
                payday = left
            elif days[right - 1] >= k:
                payday = right
            break
    return payday


def get_day(left, right):
    return math.floor(((right - left) / 2) + left)


if __name__ == '__main__':
    print(' '.join(list(map(str, solution()))))