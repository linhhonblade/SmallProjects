# ============================================ 9. inherit slide

# Create a class Spaceship. A spaceship can always "launch" successfully (launch() always returns True :))) )
# Rocket is a child class of Spaceship with its own weight, max_weight and cargo_items
# Overwrite launch() function (use formula above)
# Implement function load_item(text_file) with input is a string of text file name

import random

class Item():
    name = ''
    cargo_weight = 0.0
    def __init__(self, name, weight):
        self.name = name
        self.cargo_weight = weight

class SpaceShip():
    def launch(self):
        return True

class Rocket(SpaceShip):
    weight = 0.0
    max_weight = 0.0
    cargo_items = []

    def load_items(self, text_file):
        with open(text_file, 'r+') as file:
            for line in file:
                line_extract = line.split('=')
                item = Item(line_extract[0], float(line_extract[1]))
                self.cargo_items.append(item)
    
    def _get_cargo_weight(self):
        sum = 0
        for item in self.cargo_items:
            sum += item.cargo_weight
        return sum

    def launch(self):
        chance = 100 * random.random() - 5 * (self._get_cargo_weight()/(self.max_weight - self.weight))
        if chance > 0:
            return True
        else:
            return False

R1 = Rocket()
R1.weight = 10000
R1.max_weight = 20000
R1.load_items('cargo1.txt')
print(R1.launch())