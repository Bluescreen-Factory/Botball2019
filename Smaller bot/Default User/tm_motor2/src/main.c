// best run in console

#include <kipr/botball.h>
#include <sys/time.h>
#include <inttypes.h>

int power = 400; // ticks/s

struct Data {
    int ts[1000];
    int tick[1000];
    int mtick[1000];
} data;

struct MotorD {
    int id;
    int port;
    int ticks;
    int low;
    int diff;
} mLeft, mRight;

int eval_sensor(int lowVal, int val){
    
	if (val > (lowVal + 40)){
        //printf("1:%i\n", val);
        return 1;  // high
    }
    //printf("0:%i\n", val);
    return 0;      // low
}

/*
static void trace_leftM(){
    int lastStatus = eval_sensor(mLeft.low, analog(mLeft.port));
    int status = 0;
    int ticks = 0;
    int deltaticks = 0;
    struct timespec start, stop;
    
    printf("ms, status, ticks, deltaticks\n"); 
    clock_gettime(CLOCK_MONOTONIC_RAW, &start);
    
    while (1) {
        //msleep(1);
        status = eval_sensor(mLeft.low, analog( mLeft.port));
        //printf("val: %i, status: %i\n", analog(port), status);
        
        if (status != lastStatus){
            clock_gettime(CLOCK_MONOTONIC_RAW, &stop);
        	lastStatus = status;
           
            ticks = gmpc(mLeft.id);
            //cmpc(motId);

            int64_t delta_us = (stop.tv_sec - start.tv_sec) * 1000000 + (stop.tv_nsec - start.tv_nsec) / 1000;

            //printf("%i, %i, %i, %i\n", (int)(delta_us/1000), eval_sensor(analog(mLeft.port)), ticks, ticks - deltaticks);
            deltaticks = ticks;
        }
    }
}

static void trace_rightM(){
}
*/

int eichung(struct MotorD *mot){
    int minVal=4096, maxVal=0, i;
    
    printf("eichung id: %i port: %i\n", mot->id, mot->port);
    mav(mot->id, 200);
    
    for(i=0; i<5000; i++){
        int val = analog(mot->port);
        if (val > 100 && val < minVal){
            minVal = val;
        }
        if (val < 4096 && val > maxVal){
            maxVal = val;
        }        
        //msleep(1);
    }
	printf("min: %i\nmax: %i\n", minVal, maxVal);

    /*
    while (analog(mot->port) < (minVal + 20)){msleep(2);}
    printf("s1: %i\n", analog(mot->port));
    while (analog(mot->port) > (maxVal - 20)){msleep(2);}
    printf("s2: %i\n", analog(mot->port));
    
    while (analog(mot->port) < (minVal + 20)){;}
	*/
    while (eval_sensor(minVal, analog(mot->port)) != 0) {msleep(2);}
    msleep(10);
    while (eval_sensor(minVal, analog(mot->port)) != 1) {msleep(2);}
    msleep(10);
    while (eval_sensor(minVal, analog(mot->port)) != 0) {;}
           
    freeze(mot->id);
    printf("s3: %i\n", analog(mot->port));
    msleep(500);
    ao();
    return minVal;
}

/* test threads */

void test(int runtime, struct MotorD *motLeft, struct MotorD *motRight){
    //thread id;
    
    printf("using Motor %i, analog(%i)", motLeft->id, motLeft->port);
    printf("using Motor %i, analog(%i)", motRight->id, motRight->port);

    //id = thread_create(trace_analog);

    printf("start thread\n");
    msleep(1000);
    //thread_start(id);
    // thread_wait(id);
    cmpc(motLeft->id);
    cmpc(motRight->id);
    
    // select mav or power
    mav(motLeft->id, motLeft->ticks);
    mav(motRight->id, motRight->ticks);
    //motor(motorId, power);
    
    msleep(runtime);
    
    //bmd(motId);
    freeze(motLeft->id);
    freeze(motRight->id);
    msleep(500);
    ao();    
    //thread_destroy(id);   

    /*
    int i;
    printf("\nresult:\n");
    for(i=0; i<100; i++){
        printf("%i, ", status[i]);
    }
    */
}

/*
run in console with args: ticks_motor_left, ticks_motor_right,
                    (optional): runtime
run prg:
    cd ~/Documents/KISS/Default\ User/tm_motor2
    ./bin/botball_user_program 400, 410, 3000
*/
int main(int argc, char *argv[] )
{
    //printf("press button\n");
    //wait_for_any_button();
    int ticksPerSec = 400;
    mLeft.id = 0;
    mLeft.port = 0;
    mLeft.ticks = (int)(ticksPerSec + ticksPerSec/100*8.25);   //  <--- Anpassung 
    mRight.id = 1;
    mRight.port = 1;
    mRight.ticks = ticksPerSec;
    if( argc >= 2 ) {
        printf("l: %s r : %s", argv[1], argv[2]);
        mLeft.ticks = atoi(argv[1]);
        mRight.ticks = atoi(argv[2]);
    }
    
    mLeft.low = eichung(&mLeft);
    printf("min sensor val: %i\n", mLeft.low);
    mRight.low = eichung(&mRight);
    printf("min sensor val: %i\n", mRight.low);    
    
    if( argc >= 3 ) {
    	test(atoi(argv[3]), &mLeft, &mRight);
    }

    return 0;
}


