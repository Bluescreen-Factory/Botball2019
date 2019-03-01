import wallaby as w

class Sensor(object):
    sensor_hi = xrange(1000, 4048)
    sensor_low = xrange(0, 400)
    
    def __init__(self, num):
        self.num = num
        self.sensor = w.Analog(num)

    @property
    def value(self):
        return self.sensor.value()

    def __str__(self):
        return self.__repr__()

    def __repr__(self):
        return "Sensor {}, val={}".format(self.num, self.value)
            
class SensorCounter(Sensor):
    
    sensor_divider = 5.0
        
    def __init__(self, num):
        super(SensorCounter, self).__init__(num)
        self.count = 0
        self.tick = False
    
    def wait_hi(self):
        print("wait for 1 tick")
        state = 2
        print("wait for low")
        while state != 0:
            last_val = self.value
            if (state == 2) and (last_val in Sensor.sensor_low):
                state -= 1
                print("wait for hi..")
            if (state == 1) and (last_val in Sensor.sensor_hi):
                print("hi")
                state -= 1
        self.tick = True

    def update(self):
        act_val = self.value
        # print(act_val)
        # print("*** {} [{}]".format(status, act_val))
        # if act_val > (last_val + last_val*0.5):
        if self.tick and (act_val in Sensor.sensor_low):
            self.tick = False
            # print("{} [{}]".format(status, act_val))

        # if act_val < (last_val - last_val*0.5):
        if not self.tick and (act_val in Sensor.sensor_hi):
            self.tick = True
            # print("{} [{}]".format(status, act_val))

            self.count += 1

    def get_result(self):
        return (self.count/SensorCounter.sensor_divider)
        
    def __str__(self):
        return self.__repr__()

    def __repr__(self):
        return "Sensor {}, val={}, count={}".format(self.num, self.value, self.count)
