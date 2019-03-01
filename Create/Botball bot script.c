#include <kipr/botball.h>

void runChain(double width) //unter 4% ungenauigkeit
{
    cmpc(0);
    double motor1constant = 1830; //per round
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
    msleep(1000);
}
void rotate(double grad) //unter 1% ungenauigkeit
{
    while (grad > 180)
    {
        grad -= 360;
    }
    while (grad < -180)
    {
        grad += 360;
    }
    double val = (grad / 180) * 200 + 300;
    if (grad > 100)
        val = (grad / 180) * 225 + 275;
    double time = 0;
    if (grad > 0)
    {
        time = (grad / 360.0 * 1675.0 * 500 / val);
        create_drive_direct(-val, val);
    }
    if (grad < 0)
    {
        time = (-grad / 360.0 * 1675.0 * 500 / val);
        create_drive_direct(val, -val);
    }

    msleep(time);
    create_stop();
}

void ride(double distance) //unter 1% ungenauigkeit
{
    double val = 450;
    if (distance < 100)
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
}
void driveToNextHole()
{
    int val = analog(5) + 200;
    int matches = 0;
    create_drive_direct(100, 100);
    while (matches < 5)
    {
        if (analog(5) > val)
        {
            matches++;
        }
        else
            matches = 0;
    }
    ride(-150);
}
void getPeople()
{
    rotate(-90);
    runServo(3, 15);
    msleep(1000);
    ride(200);
    runServo(3, 100);
    ride(-200);
    rotate(90);
    printf("\n");
}
int main()
{
    printf("started\n");
    enable_servos();

    create_connect();
    create_full();
    printf("connected\n");
    driveToNextHole();
    getPeople();
    /* rotate(180);
    //pushAmbulance();
    //driveToLeftMedicalCenter();
    //int leftIsBurning= isBurning();
    //driveBackwardsOnGrayLine();
    //turnToLeft();
    //getPeople();
    //driveToNotBurningCenter(leftIsBurning);
    printf("finished\n");
    create_disconnect();
    printf("disconnected\n");*/
    return 0;
}
/*void driveTillBump()
{
    create_drive_direct(300, 300);
    while (get_create_rbump() == 0 && get_create_lbump() == 0)
    {
    }
    create_stop();
}
void turnAroundAndRide()
{
    rotate(180);
    create_drive_direct(300, 300);
    set_create_distance(0);
    while (get_create_distance() < 1000)
    {
    }
    create_stop();
}*/