# =========================================== pass slide

# Let's suffer human by repeating the input of weigh, max_weight, cargo_weights 10 times and return the result of launch() for each case at once
# Let's suffer human more by forcing them to input 10 cargo items to each rocket
# Let's troll human by making the input infinitely :)))
# Hooman entered something like max_weight <= weight, just silently skip that case if it happens

for i in range (9):
    cargo_weights = [4000, 9000]

    # Write a program to get weight(kg), maxWeight(kg), and cargoWeight(kg) from user input
    weight = float(input('Enter weight of rocket: '))
    max_weight = float(input('Enter maximum weight of rocket: '))
    if max_weight <= weight:
        continue
    # cargo_weight = float(input('Enter cargo weight of rocket: '))
    for i in range(9):
        cargo_weight = float(input('Enter cargo weight of rocket: '))
        cargo_weights.append(cargo_weight)

    # Calculate the chance of explosion = 5 * (cargo_weight/(max_weight-weight)) (%)
    chance_of_explosion = 5 * (sum(cargo_weights)/(max_weight - weight))
    print('Chance of explosion: ' + str(chance_of_explosion) + '%\n')

    if chance_of_explosion >= 100:
        print ("Launch failed\n")
    else:
        print("Launched successfully\n")

