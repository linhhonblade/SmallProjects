# ============================================ 6. function slide

# Create a function launch() to check if rocket can launch successfully
# The chance to launch successfully = 100 - 5 * (cargo_weight/(max_weight-weight)) (%)
# if the result > 0 return True (launched successfully)
# else return False

weight = float(input('Enter weight of rocket: '))
max_weight = float(input('Enter maximum weight of rocket: '))
# cargo_weight = float(input('Enter cargo weight of rocket: '))
cargo_weight = float(input('Enter cargo weight of rocket: '))

def launch():
    chance = 100 - 5 * (cargo_weight/(max_weight - weight))
    if chance > 0:
        return True
    else:
        return False

print(launch())