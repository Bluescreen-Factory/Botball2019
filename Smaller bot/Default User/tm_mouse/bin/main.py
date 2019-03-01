#!/usr/bin/python

# best run in console
# for c impl. see https://stackoverflow.com/questions/11451618/how-do-you-read-the-mouse-button-state-from-dev-input-mice

import os, sys
from wallaby import *
import struct, os

file = None

class Point:
    x = 0.0
    y = 0.0
    btn = 0
    bLeft = 0
    bRight = 0

def getMouseEvent():
    global file
    buf = file.read(3) 
    x,y = struct.unpack( "bb", buf[1:] )
    btn = struct.unpack( "b", buf[0] )
    dis = Point()
    dis.x = x
    dis.y = y
    #dis.btn = btn[0]
    bLeft = btn[0] & 0x1
    bRight = (btn[0] & 0x2 ) > 0
    dis.bRight = bRight
    dis.bLeft = bLeft
    return dis

def main():
    global file
    file = open( "/dev/input/mice", "rb" )

    point_x = 0
    point_y = 0    

    print("abort with left button click")
    
    btnL = getMouseEvent().bLeft
    while btnL == 1:
        dis = getMouseEvent()
        btnL = dis.bLeft
        
    print("btnL: %d" % (btnL))
        
    while btnL == 0:
        dis = getMouseEvent()
        point_x = point_x + dis.x
        point_y = point_y + dis.y
        print ("%d  %d Bl: %d Br: %d" % (point_x, point_y, dis.bLeft, dis.bRight))
        btnL = dis.bLeft

    # wait until btn released    
    while getMouseEvent().bLeft == 1:
        pass

    file.close()

if __name__== "__main__":
    sys.stdout = os.fdopen(sys.stdout.fileno(),"w",0)
    main()
