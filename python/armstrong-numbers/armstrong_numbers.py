def is_armstrong_number(number):
    check = 0
    length = len(str(number))
    for x in str(number):
        check += int(x)**length
    if check == number:
        return True
    return False


is_armstrong_number(9474)
