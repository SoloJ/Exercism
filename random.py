import random
import string

letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
print(random.choice(letters) +
      random.choice(letters) + str(random.randint(100, 999)))
print("does this even print anything")
