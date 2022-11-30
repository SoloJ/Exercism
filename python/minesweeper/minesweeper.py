def annotate(minefield):
    """Create an array with the minefield"""
    mine_array = [[]]
    counter = []
    stringer = ""
    answer = []
    for x in range(len(minefield)):
        mine_array.append([*minefield[x]])
    mine_array.pop(0)

    for height_index, h_value in enumerate(mine_array):
        for width_index, mine in enumerate(mine_array[height_index]):
            counter = []
            if mine != " " and mine != "*":
                """check for invalide character"""
                raise ValueError("The board is invalid with current input.")
                """check is any entry is to short"""
            if len(mine_array[0]) != len(mine_array[height_index]):
                raise ValueError("The board is invalid with current input.")
            "skip when entry is a mine"
            if mine == "*":
                continue
                """up"""
            if height_index == 0:
                counter.append(0)
            elif mine_array[height_index-1][width_index] == "*":
                counter.append(1)
            else:
                counter.append(0)
                """down"""
            if height_index == (len(mine_array)-1):
                counter.append(0)
            elif mine_array[height_index+1][width_index] == "*":
                counter.append(1)
            else:
                counter.append(0)
                """right"""
            if width_index == (len(mine_array[height_index])-1):
                counter.append(0)
            elif mine_array[height_index][width_index+1] == "*":
                counter.append(1)
            else:
                counter.append(0)
                """left"""
            if width_index == 0:
                counter.append(0)
            elif mine_array[height_index][width_index-1] == "*":
                counter.append(1)
            else:
                counter.append(0)
                """upright"""
            if width_index == (len(mine_array[height_index])-1) or height_index == 0:
                counter.append(0)
            elif mine_array[height_index-1][width_index+1] == "*":
                counter.append(1)
            else:
                counter.append(0)
                """upleft"""
            if width_index == 0 or height_index == 0:
                counter.append(0)
            elif mine_array[height_index-1][width_index-1] == "*":
                counter.append(1)
            else:
                counter.append(0)
                """downright"""
            if height_index == (len(mine_array)-1) or width_index == (len(mine_array[height_index])-1):
                counter.append(0)
            elif mine_array[height_index+1][width_index+1] == "*":
                counter.append(1)
            else:
                counter.append(0)
                """downleft"""
            if height_index == (len(mine_array)-1) or width_index == 0:
                counter.append(0)
            elif mine_array[height_index+1][width_index-1] == "*":
                counter.append(1)
            else:
                counter.append(0)

            mine_array[height_index][width_index] = sum(counter)
    print("stop!")
    for index,  entry in enumerate(mine_array):
        for letter_entry in entry:
            if letter_entry == 0:
                letter_entry = " "
            stringer += str(letter_entry)
        answer.append(stringer)
        stringer = ""
    return answer
