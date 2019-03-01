import random

class Wallaby:
    def __init__(self):
        pass

    class Analog:
        def __init__(self, x):
            port = x
        def value(self):
            v = random.randint(500, 2000)
            return v
   