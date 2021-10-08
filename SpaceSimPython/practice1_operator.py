# ========================================== 1. input/output slide

# Calculate chance of explosion when launching rocket
# A rocket has its own weight (without cargo)
# total weight of cargo is cargo_weight
# A rocket has it own max_weight (it cannot carry more cargo if the total weight exceeds this value)

# Write a program to get weight(kg), maxWeight(kg), and cargoWeight(kg) from user input
# Calculate the chance of explosion = 5 * (cargo_weight/(max_weight-weight)) (%)
# Return the value calculated
# Print the data in human-like format: "A rocket with xxx kg weight carrying yyy kg cargo has a chance of explosion equal to zzz % if its maximum weight is xzy kg."

# ============================================ 2. end of section 2: variables and collections

# The rocket already pre-filled with 2 cargo items (4000kg and 9000kg), The cargo_weight is actualy sum weight of all cargo items
# cargo_weights is a list of dictionary to store weight of cargo items
# the cargo_weight (input by user) is additional cargo
# Re-calculate the chance of explosion
# Create a dictionary to store all values (weight, max_weight, cargo_weights, chance_of_explosion)
# Show the dictionary in console ouput

# ============================================ 3. if statement slide

# If chance of explosion greater or equal 100, print "Launch failed"
# Esle print "Launched successfully"

# =========================================== 4. pass slide

# Let's suffer human by repeating the input of weigh, max_weight, cargo_weights 10 times and return the result of launch() for each case at once
# Let's suffer human more by forcing them to input 10 cargo items to each rocket
# Let's troll human by making the input infinitely :)))
# Hooman entered something like max_weight <= weight, just silently skip that case if it happens

# ============================================ 5. file handling slide

# Stop suffering each other
# There are too many cargo items, it is placed in a text file.
# The file includes cargo item name and its weight
# Load all cargo items from text file to store it in a list of dictionaries

# ============================================ 6. function slide

# Create a function launch() to check if rocket can launch successfully
# The chance to launch successfully = 100 - 5 * (cargo_weight/(max_weight-weight)) (%)
# if the result > 0 return True (launched successfully)
# else return False

# ============================================ 7. import slide

# Actually chance_of_explosion doesn't depend on weights only
# More accurately, we should consider some randomness factor
# The chance to launch successfully = (random number from 0 to 100) - 5 * (cargo_weight/(max_weight-weight)) (%)
# Re implement launch() (hint: use random module)


# ============================================ 8. class slide

# Create a class Item to represent a cargo item
# A cargo item includes name and its weight
# Store the data loaded from text file to a list of Item object

# ============================================ 9. inherit slide

# Create a class Spaceship. A spaceship can always "launch" successfully (launch() always returns True :))) )
# Rocket is a child class of Spaceship with its own weight, max_weight and cargo_items
# Overwrite launch() function (use formula above)
# Implement function load_item(text_file) with input is a string of text file name

# ============================================ homework project
# Create a simulation to check whether we can launch a rocket successfully with two use cases:
# Use-case 1:
    # rocket: weight = 18000, max_weight_29000
    # cargo_items: cargo1.txt
# Use-case 2:
    # rocket: weight = 10000, max_weight = 18000
    # cargo_items: cargo2.txt
# Hint:
    # Create a Simulation class which contains
    # The Rocket
    # a prepare() function to load all parameter for the use case
    # run() function to run simulation (actually running the launch() function of rocket)
# More: Reconstruct the project
    # pacakage space: SpaceShip, Rocket, Item
    # package simulator: Simulation
    # main.py to run simulation for each use case

