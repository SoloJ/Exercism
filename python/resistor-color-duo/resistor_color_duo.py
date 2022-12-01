def value(colors):
    color_code = {
        "black": 0,
        "brown": 1,
        "red": 2,
        "orange": 3,
        "yellow": 4,
        "green": 5,
        "blue": 6,
        "violet": 7,
        "grey": 8,
        "white": 9
    }
    answer = ""
    for c in colors:
        if len(answer) == 2:
            break
        else:
            answer += str(color_code[c])
    return int(answer)
