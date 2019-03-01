1#!/usr/bin/python

# motor re (0):
# 20s 20% : 6358 / 3.4  (ticks / turns)
# 20s 20% : 6354 / 3.4
# 20s 20% : 6356 / 3.4
# 60s 30% : 29235 / 15.4
# with threading *****
# 10turns 20% : [18520, 18513]   (ticks)
# 10turns 30% : [18555, 18565]
#  5turns 40% : [9555, 9583, 9570, 9565]
#  5turns 20% : [9533, 9547, 9558, 9524]
# without threading *****
# 10turns 20% : []
#  5turns 40% : [9569, 9560, 9571, 9547]

# motor li (1):
# 20s 20% : 6433 / 3.2  (ticks / turns)
# 20s 20% : 6440 / 3.2
# 20s 20% : 6428 / 3.2
# 60s 30% : 29047 / 14.2
# with threading *****
# 10turns 20% : [20137, 20126, 20132]  (ticks)
# 10turns 30% : [20442, 20418, 20420]
#  5turns 20% : [10162,]
# without threading *****
# 10turns 20% : [20386, 20387]

from sensor import *
#import wallaby as w
   
def main():
    # w = Wallaby()
    sens = SensorCounter(1)
    print(sens)



if __name__== "__main__":
    sys.stdout = os.fdopen(sys.stdout.fileno(),"w",0)
    main();
