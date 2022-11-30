def find_fewest_coins(coins, target):
    coins.sort(reverse=True)
    temp = [0]*(len(coins))
    temp_lol = [[]]

    "create array for all values"
    for x in range(int(target/coins[-1])):
        temp_lol.append(temp)

    i = 1
    for coin in coins:
        temp_target = target
        for z in range(int(temp_target/coin)):
            while temp_target/coin > 1:
                answer = find_combo(coins, temp_target)
                temp_target -= coin
                temp_lol[i].append(int(temp_target/coin))
                i += 1
                coins.pop(0)

    print("hello")


def find_combo(coins, target):
    sum_coins = 0
    temp = [0]*(len(coins))
    for index,  coin_value in enumerate(coins):
        while sum_coins < target:
            if (target-sum_coins)-coin_value >= 0:
                temp[index] += 1
                sum_coins += coin_value
            else:
                break
    return temp
