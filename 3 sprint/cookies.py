def solution():
    _ = input()
    children = sorted(list(map(int, input().split(' '))))
    _ = input()
    cookies = sorted(list(map(int, input().split(' '))))
    index = 0
    ok = 0
    for cookie in cookies:
        child = children[index]
        if cookie >= child:
            ok += 1
            index += 1
            if index == len(children):
                return ok

    return ok


if __name__ == '__main__':
    print(solution())