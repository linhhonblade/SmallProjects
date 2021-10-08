# ============================================ 7. import slide

# Actually chance_of_explosion doesn't depend on weights only
# More accurately, we should consider some randomness factor
# The chance to launch successfully = (random number from 0 to 100) - 5 * (cargo_weight/(max_weight-weight)) (%)
# Re implement launch() (hint: use random module)

import random

weight = float(input('Enter weight of rocket: '))
max_weight = float(input('Enter maximum weight of rocket: '))
# cargo_weight = float(input('Enter cargo weight of rocket: '))
cargo_weight = float(input('Enter cargo weight of rocket: '))

def launch():
    chance = 100 * random.random() - 5 * (cargo_weight/(max_weight - weight))
    if chance > 0:
        return True
    else:
        return False

print(launch())