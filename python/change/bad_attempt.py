def find_fewest_coins(coins, target):
    coins.sort(reverse=True)
    index_lol = [[]]
    index_it = 0
    while index_it < (len(coins)):
        temp_target = target
        temp = []
        for index, coin in enumerate(coins):
            if int(temp_target/coin) >= 1 and target % coin == 0:
                temp.append(int(temp_target/coin))
                temp_target -= coin
                break
            elif int(temp_target/coin) >= 1:
                temp.append(int(target/coin))
                temp_target -= coin
            else:
                temp.append(0)
        index_it += 1
        index_lol.append(temp)
        coins.pop(0)

    total, answer = sum(coins, temp)
    if total == target:
        return answer
    raise ValueError("cant make target with the coins given")


def sum(coins, temp):
    total = 0
    answer = []
    for index,  value in enumerate(temp):
        total += (value * coins[index])
        while value > 0:
            answer.append(coins[index])
            value -= 1
    answer.sort()
    return total, answer


def find_combo(coins, target):
    sum_coins = 0
    temp = []
    for coin_value in coins:
        while sum_coins < target:
            if (target-sum_coins)-coin_value >= 0:
                temp.append(coin_value)
                sum_coins += coin_value
            else:
                break
    return temp

    def alternate_combo(coins, target):


def find_combo(coins, target):
    sum_coins = 0
    temp = []
    for coin_value in coins:
        while sum_coins < target:
            if (target-sum_coins)-coin_value >= 0:
                temp.append(coin_value)
                sum_coins += coin_value
            else:
                break
    return temp


def find_fewest_coins(coins, target):
    all_coin_combinations = [[]]
    coins.sort(reverse=True)
    while len(coins) > 0:
        combos = int(target/coins[0])
        while combos >= 1:
            change = find_combo(coins, target, combos)
            if sum(change) == target:
                all_coin_combinations.append(change)
            combos -= 1
        coins.pop(0)

    """Find change combination with least number of coins within possible combination list"""

    answer = min(all_coin_combinations, key=len)
    answer.sort()
    return answer


"""Function walks through change and finds possible change combinations to hit target"""


def find_combo(coins, target, combos):
    sum_coins = 0
    temp = []
    for coin_value in coins:
        while combos >= 1:
            temp.append(coin_value)
            sum_coins += coin_value
            combos -= 1
    return temp
