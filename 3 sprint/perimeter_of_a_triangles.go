def solution():
    _ = input()
    hs = sorted(list(map(int, input().split(' '))), reverse=True)
    for i, s0 in enumerate(hs[:-2]):
        s1 = hs[i + 1]
        s2 = hs[i + 2]
        if s0 < s1 + s2:
            return s1 + s2 + s0
    return 0


if __name__ == '__main__':
    print(solution())