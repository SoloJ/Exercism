import random


class Robot:
    def __init__(self):
        self.name = self.new_name()

    def reset(self):
        random.seed()
        self.name = self.new_name()

    def new_name(self):
        letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        output = random.choice(
            letters) + random.choice(letters) + str(random.randint(100, 999))
        return output
