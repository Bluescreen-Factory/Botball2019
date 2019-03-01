import os, sys
import wallaby as w
from thread import start_new_thread, allocate_lock
from time import sleep
   
result = "no set"
thread_end = False
thread_started = False
threads_num = 0
lock = allocate_lock()

def task_poll(f):
    global thread_end
    global threads_num
        
    # threads_num += 1
    while not thread_end:
        f.update()

    print("poll end: {}".format(f.get_result()))
    # threads_num -= 1
            
def task_time(t, mot, velo, f):
    global thread_end
    global thread_started
    global threads_num
    global result
    
    w.cmpc(mot)
    w.motor(mot, velo)   
    
    # lock.acquire()
    # threads_num += 1
    # lock.release()
    thread_started = True
    print("wait {}s ...".format(t))
    while t > 0:    # wait till expected time
        print(t),
    	sleep(1);
        t -=1

    w.freeze(mot)
    print("\nstopping ...")

    thread_end = True
    # lock.acquire()
    # threads_num -= 1
    # lock.release()
    result = "motor {}\nruntime: {}\nmotor ticks: {}\nturns: {}".format(mot, runtime, w.gmpc(mot), f.get_result())

def task_rotate(turns, mot, velo, f):
    global thread_end
    global thread_started
    global threads_num
    global result
    
    w.cmpc(mot)
    w.motor(mot, velo)   
        
    # lock.acquire()
    # threads_num += 1
    # lock.release()
    thread_started = True
    print("wait {} turns ...".format(turns))

    #data = []
    last_val = f.get_result()
    while last_val < turns:    # wait till expected turns
        val = f.get_result()
        if val > last_val:
           print("%2.1f\t%i" % (last_val, w.gmpc(mot)))
           last_val = val

    w.freeze(mot)
    print("\nstopping ...")

    thread_end = True
    # lock.acquire()
    # threads_num -= 1
    # lock.release()
    result = "motor {}\nturns: {}\nmotor ticks: {}".format(mot, f.get_result(), w.gmpc(mot) )

def run_time(motor_id, pwr, runtime, sens):
    start_new_thread(task_poll,(sens,))
    start_new_thread(task_time,(runtime, motor_id, pwr, sens))
    while not thread_started:
        pass
    while not thread_end:
        pass
    while threads_num > 0:
        pass  
    return result

def run_turns(motor_id, pwr, turns, sens):
    start_new_thread(task_poll,(sens,))
    start_new_thread(task_rotate,(turns, motor_id, pwr, sens))
    while not thread_started:
        pass
    while not thread_end:
        pass
    while threads_num > 0:
        pass  
    return result
    