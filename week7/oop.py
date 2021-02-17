class Car:
    def __init__(self, make, model, color):
        self.make = make
        self.model = model
        self.color = color

    # Method
    def listCar(self):
        print(self.make, self.model)

#Creates new instance of Class Car
ourCar = Car("Ford", "F-350", "Ruby Red")
yourCar = Car("Mazda", "3", "White")

print(ourCar.color)
ourCar.color = "Midnight Black"
print(ourCar.color)
