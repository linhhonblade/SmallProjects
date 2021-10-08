# Write a program to get weight(kg), maxWeight(kg), and cargoWeight(kg) from user input
weight = float(input('Enter weight of rocket: '))
max_weight = float(input('Enter maximum weight of rocket: '))
# cargo_weight = float(input('Enter cargo weight of rocket: '))
cargo_weight = float(input('Enter cargo weight of rocket: '))

# Calculate the chance of explosion = 5 * (cargo_weight/(max_weight-weight)) (%)
chance_of_explosion = 5 * (cargo_weight/(max_weight - weight))

print('Chance of explosion: ' + str(chance_of_explosion) + '%\n')

# Print the data in human-like format: "A rocket with xxx kg weight carrying yyy kg cargo has a chance of explosion equal to zzz % if its maximum weight is xzy kg."
message = 'A rocket with ' + str(weight) + ' kg weight carrying ' + str(cargo_weight) + ' kg cargo has a chance of explosion equal to ' + str(chance_of_explosion) + ' % if its maximum weight is ' + str(max_weight) + ' kg\n'
print(message)

another_message = 'A rocket with {a} kg weight carrying {b} kg cargo has a chance of explosion equal to {c} % if its maximum weight is {d} kg.\n'.format(a = weight, b = cargo_weight, c = chance_of_explosion, d = max_weight)
print(message)

yet_another_message = f'A rocket with {weight} kg weight carrying {cargo_weight} kg cargo has a chance of explosion equal to {chance_of_explosion} % if its maximum weight is {max_weight} kg.\n'
print(yet_another_message)

# ==================================================================
# The rocket already pre-filled with 2 cargo items (4000kg and 9000kg), The cargo_weight is actualy sum weight of all cargo items
# cargo_weights is a list of dictionary to store weight of cargo items

cargo_weights = [4000, 9000]
