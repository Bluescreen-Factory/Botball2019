#include <kipr/botball.h>
int fin = 0;
int leftBurning = 1;

void runChain(double width) //unter 4% ungenauigkeit
{
    cmpc(0);
    double motor1constant = -1830; //per round
    double speed = 1000;
    mrp(0, speed, motor1constant * width - (speed / motor1constant));
    bmd(0);
    freeze(0);
    msleep(500);
}
void runServo(int port, double prozent)
{
    set_servo_position(port, (int)(2047.0 * prozent / 100.0));

    while (get_servo_position(port) < ((int)(2047.0 * prozent / 100.0)) - 2 && get_servo_position(port) > ((int)(2047.0 * prozent / 100.0)) + 2)
    {
    }
    msleep(500);
}
void rotate(double grad) //unter 1% ungenauigkeit
{
    double speed = 250.0;
    while (grad > 180)
    {
        grad -= 360;
    }
    while (grad < -180)
    {
        grad += 360;
    }

    double time;
    double ratio = 1.851867 * speed + 1592.5933;
    if (grad > 0)
    {
        time = ((ratio * grad) / speed);
        create_drive_direct(speed, -speed);
    }
    if (grad < 0)
    {
        time = ((-grad * ratio) / speed);
        create_drive_direct(-speed, speed);
    }
    msleep(time);
    create_stop();
}

void ride(double distance, double val) //unter 1% ungenauigkeit
{
    if (val == -1)
        val = 450;
    if (distance < 100 && distance > -100)
        val = 50;
    double time = distance * 2.2 * (500 / val);
    if (time < 0)
    {
        time -= time + time;
        create_drive_direct(-val, -val);
    }
    else
        create_drive_direct(val, val);
    msleep(time);
    create_stop();
    msleep(250);
}
void door(int i)
{
    if (i)
    {
        runServo(2, 50);
    }
    else
    {
        runServo(2, 94);
    }
}
void ignoreOneHole()
{

    int matches = 0;
    create_drive_direct(100, 100);
    msleep(400);
    int val = analog(5) + 400;
    while (matches < 5)
    {
        if (analog(5) > val)
        {
            matches++;
            msleep(10);
        }
        else
            matches = 0;
    }
    msleep(400);
    val = analog(5) - 450;
    while (matches < 5)
    {
        if (analog(5) < val)
        {
            matches++;
            msleep(10);
        }
        else
            matches = 0;
    }
    ride(100, 200);
}
void driveToNextHole()
{
    int val = analog(5) + 300;
    int matches = 0;
    create_drive_direct(100, 100);
    while (matches < 5)
    {
        if (analog(5) > val)
        {
            matches++;
            msleep(10);
        }
        else
            matches = 0;
    }
    ride(-185, 350);
}
void getUpperPeople()
{

    runServo(3, 55);
    msleep(500);

    runChain(2);

    runServo(3, 50);
    msleep(200);
    runServo(3, 47);
    msleep(200);
    runServo(3, 45);
    msleep(500);

    runChain(-2);
    fin = 0;
}
void getPeople()
{
    rotate(90);
    fin = 1;
    thread tid = thread_create(getUpperPeople);
    thread_start(tid);

    door(1);
    msleep(500);
    ride(160, 250);
    door(0);
    while (fin)
    {
        /* code */
    }
    ride(-160, 250);
    rotate(-90);
}
void getFirstPeople()
{
    rotate(90);
    fin = 1;
    thread tid = thread_create(getUpperPeople);
    thread_start(tid);

    door(1);
    msleep(500);
    ride(140, 250);
    door(0);
    while (fin)
    {
        /* code */
    }
    ride(-140, 250);
    rotate(-90);
}

void driveToLastHole()
{
    ride(150, 200);
}
void saveCivis()
{
    driveToNextHole();
    getFirstPeople();
    ignoreOneHole();
    driveToNextHole();
    getPeople();
    ignoreOneHole();
    driveToNextHole();
    getPeople();
    ignoreOneHole();
    driveToLastHole();
    getPeople();
}

int testCamObjects()
{
    int gCount = get_object_count(0);
    rectangle g[gCount];
    int rCount = get_object_count(1);
    rectangle r[rCount];
    int i = 0;
    while (i < gCount)
    {
        g[i] = get_object_bbox(0, i);
        int j = 0;
        while (j < rCount)
        {
            r[j] = get_object_bbox(1, j);
            if (g[i].ulx < r[j].ulx && g[i].uly < r[j].uly)
            {
                if (g[i].ulx + g[i].width > r[j].ulx + r[j].width && g[i].uly + g[i].height > r[j].uly + r[j].height)
                {
                    return 1;
                }
            }
            j++;
        }
        i++;
    }
    return 0;
}
void lookIfBurning()
{
    camera_open_black();
    int counter = 0;
    int i = 0;
    int burns = 0;
    while (i < 20)
    {
        while (!camera_update() && counter < 1000)
        {
            counter++;
        }
        if (counter == 100)
        {
            printf("noImage\n");
            beep();
            return;
        }
        if (testCamObjects())
        {
            burns++;
            if (burns == 5)
            {
                leftBurning = 1;
                return;
            }
        }
        leftBurning = 0;
        i++;
    }
}
void startDrive()
{
    ride(850, 300);
    rotate(90);
    ride(400, 200);
    door(1);
    rotate(90);
    lookIfBurning();
    ride(1300, 300);
    door(0);
    rotate(-180);
}
void toNonBurningCenter()
{
    if (leftBurning)
    {
        ride(300, 200);
        rotate(90);
        ride(300, 200);
        rotate(-90);
        ride(200, 200);
    }

    else
    {
        ride(550, 250);
    }
}
int main()
{
    //start
    printf("started\n");
    enable_servos();
    create_connect();
    create_full();
    printf("connected\n");
    //*reset
    door(0);
    runChain(-2);
    //*/
    //*fullGame
    startDrive();
    saveCivis();
    toNonBurningCenter();
    //*/
    //end
    printf("finished\n");
    create_disconnect();
    printf("disconnected\n");
    return 0;
}

/*
    1600 = x * 90 / 100
    1600 / 90 = x / 100
    1600 / 90 * 100 = x
    1777,7777777777777777777777777778 = x
    
    740 = x * 90 / 250
    740 / 90 * 250 = x
    2055,5555555555555555555555555556 = x
    
    S(100, 1777,78)
    E(250, 2055,56)
    k = (1777,78 - 2055,56) / (100 - 250) = 1,8518666666666666666666666666667
    1777,78 = 1,851867 * 100 + d
    1777,78 - 185,1867 = d
    d = 1592,5933
    r = 1,851867 * spd + 1592,5933
    */
