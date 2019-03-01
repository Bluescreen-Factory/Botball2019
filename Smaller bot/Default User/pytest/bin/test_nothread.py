import os, sys
import wallaby as w
from time import sleep
   
result = "no set"

def task_rotate(turns, mot, velo, sens):
    global result
    
    w.cmpc(mot)
    w.motor(mot, velo)   
        
    print("wait {} turns ...".format(turns))

    sens.update()
    last_val = sens.get_result()
    while last_val < turns:    # wait till expected turns
        sens.update()
        val = sens.get_result()
        if val > last_val:
           print(last_val),
           last_val = val

    w.freeze(mot)
    print("\nstopping ...")

    result = "motor {}\nturns: {}\nmotor ticks: {}".format(mot, sens.get_result(), w.gmpc(mot) )

def run_turns(motor_id, pwr, turns, sens):
    task_rotate(turns, motor_id, pwr, sens)

    return result
