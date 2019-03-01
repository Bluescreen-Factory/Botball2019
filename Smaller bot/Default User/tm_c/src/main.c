#include <kipr/botball.h>
#include <sys/time.h>
#include <inttypes.h>

void test1(int mot){
    printf("speed/ticks: real ticks\n");
    int i;
    
    for (i=400; i<1500; i+=200){
        cmpc(mot);
        mrp(mot,i,2020);
        //msleep(8000);
        //freeze(mot);
        bmd(mot);
        
        printf("%i: %i\n", i, gmpc(mot));
        msleep(5000);
    }
}

void test2(int mot){
    printf("speed: real ticks\n");
    int i;
    for (i=200; i<1500; i+=200){
        cmpc(mot);
        mav(mot,i);
        msleep(5000);
        freeze(mot);
        
        printf("%i: %i\n", i, gmpc(mot));
		msleep(5000);
    }
}

int status[100];


static void trace_button(){
//static void trace_button(int (*fun)()){
    int i;
    for(i=0; i<100; i++){
    	msleep(100);
        status[i] = right_button();
        // status[i] = fun();
        printf("%i, ", status[i]);
    }
}

/* test threads */
void test3(){
    thread id;
    //void (*fun_ptr)(int) = right_button;
    id = thread_create(trace_button);
    // id = thread_create(trace_button(right_button));
    printf("start thread\n");
    msleep(1000);
    thread_start(id);
    thread_wait(id);
    thread_destroy(id);   
    
    int i;
    printf("\nresult:\n");
    for(i=0; i<100; i++){
        printf("%i, ", status[i]);
    }
}

void test_time(){
    // struct timeval stop, start;
    struct timespec start, stop;
    
    int i;
    for (i=0; i<20; i++){
        // gettimeofday(&start, NULL);
        clock_gettime(CLOCK_MONOTONIC_RAW, &start);
        
        msleep(410);

        // gettimeofday(&stop, NULL);
        // printf("took %lu\n", stop.tv_usec - start.tv_usec);
        
        clock_gettime(CLOCK_MONOTONIC_RAW, &stop);
        int64_t delta_us = (stop.tv_sec - start.tv_sec) * 1000000 + (stop.tv_nsec - start.tv_nsec) / 1000;
        printf("took %" PRId64 "us\n", delta_us);
        printf("took %ims\n", (int)(delta_us/1000));
    }
}

int main()
{
    printf("press button\n");
    //wait_for_any_button();
    //test1(0);   
    //test2(0);
    //test3();      // example using threads
    //test_time();  // does not work
    
    return 0;
}
