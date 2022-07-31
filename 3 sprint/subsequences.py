def solution():
    a = input()
    b = input()
    if a == '' or a == b:
        return True
    if len(b) < len(a):
        return False
    index = 0
    for vb in b:
        if vb == a[index]:
            index += 1
            if index == len(a):
                return True
    return False


if __name__ == '__main__':
    print(solution())