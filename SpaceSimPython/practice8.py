# ============================================ 8. class slide

# Create a class Item to represent a cargo item
# A cargo item includes name and its weight
# Store the data loaded from text file to a list of Item object


class Item():
    name = ''
    cargo_weight = 0.0
    def __init__(self, name, weight):
        self.name = name
        self.cargo_weight = weight

cargo_items = []

with open('cargo1.txt', 'r+') as file:
    for line in file:
        line_extract = line.split('=')
        item = Item(line_extract[0], float(line_extract[1]))
        cargo_items.append(item)

print(cargo_items)