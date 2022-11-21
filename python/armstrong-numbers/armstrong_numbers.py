def is_armstrong_number(number):
    check = 0
    length = len(str(number))
    for x in str(number):
        check += int(x)**length
    if check == number:
        return check == number
    return check == number


is_armstrong_number()
