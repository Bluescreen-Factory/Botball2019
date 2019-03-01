#include <kipr/botball.h>
#include <math.h>

struct motor {
	int port;
    float speed;
    int tpr; // ticks per rev
};

float trackWidth = 12.8;
float wheelDiameter = 7;

float turnPerimeter;
float wheelPerimeter;

int mini(int a, int b) {
	if(a < b) return a;
    return b;
}

int maxi(int a, int b) {
	if(a > b) return a;
    return b;
}

float minf(float a, float b) {
	if(a < b) return a;
    return b;
}

float maxf(float a, float b) {
	if(a > b) return a;
    return b;
}

float clampf(float x, float min, float max) {
	return minf(maxf(x, min), max);   
}

int clampi(int x, int min, int max) {
	return mini(maxi(x, min), max);   
}

void motor_turnMs(struct motor *m, float speed, int ms) {
    m->speed = speed;
    motor(m->port, m->speed);
    msleep(ms);
    motor(m->port, 0);
    m->speed = 0;
}

void motor_turn(struct motor *m) {
    motor(m->port, m->speed);
}

void motor_setSpd(struct motor *m, float spd) {
    m->speed = spd;
    motor_turn(m);
}

int motor_turnExact(struct motor *m, int tps, int ticks, int doSleep) {
	cmpc(m->port);
    //19/12
    ticks -= (int) clampf((19./12.)*(ticks-60), 0., 115.);
    mrp(m->port, tps, ticks);
    int slp = (int)(ticks * 1100. / (float)(tps));
    if(doSleep < 1) {
    	return slp;
    }
    msleep(slp);
    return 0;
}

int motor_equalSpeed(struct motor *dest, struct motor *subj, int tps) {
	return (int) (((float) (subj->tpr) / dest->tpr)*tps);
}

void motor_rotateDeg(struct motor *r, struct motor *l, float deg) {
    float dist = turnPerimeter * deg / 360.f;
    float rev = dist / wheelPerimeter;
    float ratio = motor_equalSpeed(l, r, 1500) / 1500.f;
	int slp1 = motor_turnExact(l, 1500, (int) (rev * l->tpr), 0);
    int slp2 = motor_turnExact(r, -ratio * 1500, (int) -(rev * ratio * r->tpr), 0);
    int slp = maxi(slp1, slp2);
    msleep(slp);
	mav(r->port, 0);
    mav(l->port, 0);
}

void motor_move(struct motor *r, struct motor *l, float cm) {
    float rev = cm / wheelPerimeter;
    int slp1 = motor_turnExact(l, 1500, (int) (rev * l->tpr), 0);
    float ratio = motor_equalSpeed(l, r, 1500) / 1500.;
    int slp2 = motor_turnExact(r, 1500*ratio, (int) (rev * l->tpr * ratio), 0);
    int slp = maxi(slp1, slp2);
    msleep(slp);
	mav(r->port, 0);
    mav(l->port, 0);
}

/*

	Erste Umdrehung Zeit: 958ms
    Folgende Umdrehungen Zeit: 1176ms
	Rad radius: 3.5 cm;
    Umfang: 7*PI ~= 22cm
*/

int main()
{
    turnPerimeter = trackWidth * M_PI;
    wheelPerimeter = wheelDiameter * M_PI;
    
    
    struct motor mRight = {0, -100., 1852};
    struct motor mLeft = {1, -100., 2020};
  	/*
    motor_move(&mRight, &mLeft, 100.);
    msleep(100);
    motor_rotateDeg(&mRight, &mLeft, 90.);
    msleep(100);
    motor_move(&mRight, &mLeft, 100.);
    msleep(100);
    motor_rotateDeg(&mRight, &mLeft, 90.);
    msleep(100);
    motor_move(&mRight, &mLeft, 100.);
    msleep(100);
    motor_rotateDeg(&mRight, &mLeft, 90.);
    msleep(100);
    motor_move(&mRight, &mLeft, 100.);/*/
    //motor_rotateDeg(&mRight, &mLeft, 360.);
    // */
    /*motor_turn(&mRight);
    motor_turn(&mLeft);
    while(!digital(1)) {}
    motor_setSpd(&mRight, 100);
    motor_setSpd(&mLeft, 100);
    while(analog(0) < 2800) {}
    ao();*/
    /*motor(0, 100);
    motor(1, 100);
    float dist = 100.0; //cm
    float rev = dist / (7*M_PI);
    float time = 1176 * (rev-1);*/
    /*float sum0 = 0.;
    float sum1 = 0.;
    float max0 = 0.;
    float min0 = 0.;
    int i = 0;
    for(; i < 10; i++) {
    	//motor(1, 100);
        cmpc(0);
        cmpc(1);
        mrp(0, +1800, 1800-145);
        mrp(1, +1800, 1800-123);
        msleep(1500);
        int pos0 = gmpc(0);
        sum0 += pos0;
        printf("0: %d\n", pos0);
        int pos1 = gmpc(1);
        sum1 += pos1;
        printf("1: %d\n", pos1);
        //motor(1, 0);
        msleep(250);
    }
    printf("average %f %f\n", sum0/10., sum1/10.);//*/
    motor_move(&mRight, &mLeft, 10.);
    motor_move(&mRight, &mLeft, 10.);
    ao();
    return 0;
}
