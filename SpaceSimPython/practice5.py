# ============================================ file handling slide

# Stop suffering each other
# There are too many cargo items, it is placed in a text file.
# The file includes cargo item name and its weight
# Load all cargo items from text file to store it in a list of dictionaries

cargo_items = []

with open('cargo.txt', 'r+') as file:
    for line in file:
        line_extract = line.split('=')
        item = {'name': line_extract[0], 'cargo_weight': float(line_extract[1])}
        cargo_items.append(item)

print(cargo_items)
