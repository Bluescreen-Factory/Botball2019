#include <kipr/botball.h>
#include <sys/time.h>
#include <inttypes.h>

int port = 1;
int motId = 1;
int power = 200; // mav
//int power = 20; // power

volatile int turns = -1;
int portLow = 0;
struct Data {
    int ts[1000];
    int tick[1000];
    int mtick[1000];
} data;

int eval_sensor(int val){
    
	if (val > (portLow + 20)){
        return 0;
    }
    return 1;   
}

static void trace_analog(){
    int lastStatus = eval_sensor(analog(port));
    int status = 0;
    int ticks = 0;
    int deltaticks = 0;
    struct timespec start, stop;
    
    printf("ms, status, ticks, deltaticks\n"); 
    clock_gettime(CLOCK_MONOTONIC_RAW, &start);
    
    while (1) {
        //msleep(1);
        status = eval_sensor(analog(port));
        //printf("val: %i, status: %i\n", analog(port), status);
        
        if (status != lastStatus){
            clock_gettime(CLOCK_MONOTONIC_RAW, &stop);
        	lastStatus = status;
           
            ticks = gmpc(motId);
            //cmpc(motId);

            int64_t delta_us = (stop.tv_sec - start.tv_sec) * 1000000 + (stop.tv_nsec - start.tv_nsec) / 1000;

            printf("%i, %i, %i, %i\n", (int)(delta_us/1000), eval_sensor(analog(port)), ticks, ticks - deltaticks);
            deltaticks = ticks;
        }
    }
}

int eichung(int mot, int port){
    int minVal, i;
    
    printf("eichung\n");
    //if (mot == 0){ mav(mot, -200);}
    //else { mav(mot, 200);}
    mav(mot, 200);
    
    for(i=0; i<2000; i++){
        int val = analog(port);
        if (val < minVal){
            minVal = val;
        }
        //msleep(1);
    }
    
    while (analog(port) > (minVal+20)){
        ;
    }
    freeze(mot);
    msleep(500);
    ao();
    return minVal;
}

/* test threads */
void test3(int runtime, int motorId, int port, int power){
    thread id;
    
    printf("using Motor %i, time %i, analog(%i)", motorId, runtime, port);

    id = thread_create(trace_analog);

    printf("start thread\n");
    msleep(1000);
    thread_start(id);
    // thread_wait(id);
    cmpc(motorId);
    
    // select mav or power
    mav(motorId, power);
    //motor(motorId, power);
    
    msleep(runtime);
    
    //bmd(motId);
    freeze(motorId);
    msleep(500);
    ao();    
    thread_destroy(id);   

    /*
    int i;
    printf("\nresult:\n");
    for(i=0; i<100; i++){
        printf("%i, ", status[i]);
    }
    */
}

int main()
{
    //printf("press button\n");
    //wait_for_any_button();
    portLow = eichung(motId, port);
    printf("min sensor val: %i\n", portLow);
    
    /*
    if (motId == 0){
        power = power * -1;
    }*/
    
    test3(30000, motId, port, power);

    return 0;
}

