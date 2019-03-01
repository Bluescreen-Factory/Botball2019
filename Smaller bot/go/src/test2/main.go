package main

/*
 #cgo LDFLAGS: -L/usr/lib/ -l wallaby
 #include <wallaby/wallaby.h>
*/
import "C"
import (
	"fmt"
)

/*
	Source:
	https://botballprogramming.org/site-maps/right-brain-site-map/bonus-supplemental-information-tips/library-of-functions/
	Regex1:
	Format:\s+(.*?)\s*$
	Repl1:
	$1\r\n
	Regex2:
	(\w+)\s+(\w+)\((?:(\w+)\s+(\w+))(,\s+)?(?:(\w+)\s+(\w+))?(,\s+)?(?:(\w+)\s+(\w+))?\)|(\w+)\s+(\w+)
	Repl2:
	func $2$12($4 $3$5$7 $6$8$10 $9) $1$11 {\r\n    return C.$2$12($4$5$7$8$10)\r\n}\r\n
	Regex3:
	(\w+)\(((\w+)\s+(\w+))?((,\s+)(\w+)\s+(\w+))?((,\s+)(\w+)\s+(\w+))?\)\s*(\w+)?\s*{?\s*\r?\n\s*return\s+C.(\w+)\([^)]*\)
	Repl3:
	$1($3 $4$6$7 $8$10$11 $12) $13 {\r\n    return $13(C.$14(C.$4($3)$6C.$8($7)$10C.$12($11)))
	Regex4:
	C\.\(\)
	Repl4:
	Regex5:
	C.float32
	Repl5: 
	C.float
	Regex6:
	 (?= |\))(?! *return)
	Repl6:
	Regex7:
	return \(
	Repl7:
*/

func main() {
	ch := make(chan string)
	go func () {
		ch <- "Hello World!"
	}()
	fmt.Println(<-ch)
}

func a_button() int {
    return int(C.a_button())
}
func alloff() {
    (C.alloff())
}
func analog(p int) int {
    return int(C.analog(C.int(p)))
}
func analog10(p int) int {
    return int(C.analog10(C.int(p)))
}
func ao() {
    (C.ao())
}
/*
func atan(angle float32) float32 {
    return float32(C.atan(C.float(angle)))
}
*/
func b_button() int {
    return int(C.b_button())
}
func beep() {
    (C.beep())
}
func bk(m int) {
    (C.bk(C.int(m)))
}
func black_button() int {
    return int(C.black_button())
}
func block_motor_done(m int) {
    (C.block_motor_done(C.int(m)))
}
func bmd(m int) {
    (C.bmd(C.int(m)))
}

/*func col() int {
    return int(C.col())
}
func row() int {
    return int(C.row())
}
func s() char {
    return char(C.s())
}*/
func clear_motor_position_counter(motor_nbr int) {
    (C.clear_motor_position_counter(C.int(motor_nbr)))
}
/*
func cos(angle float32) float32 {
    return float32(C.cos(C.float(angle)))
}
*/
func digital(p int) int {
    return int(C.digital(C.int(p)))
}
func disable_servos() {
    (C.disable_servos())
}

/*func down_button() int {
    return int(C.down_button())
}*/
func enable_servos() {
    (C.enable_servos())
}

/*func exp10(num float32) float32 {
    return float32(C.exp10(C.float(num)))
}
func exp(num float32) float32 {
    return float32(C.exp(C.float(num)))
}
*/
func fd(m int) {
    (C.fd(C.int(m)))
}
func freeze(m int) {
    (C.freeze(C.int(m)))
}
func get_motor_done(m int) int {
    return int(C.get_motor_done(C.int(m)))
}
func get_motor_position_counter(m int) int {
    return int(C.get_motor_position_counter(C.int(m)))
}
func get_servo_position(srv int) int {
    return int(C.get_servo_position(C.int(srv)))
}

/*func kill_process(pid int) {
    (C.kill_process(C.int(pid)))
}
func kissSimEnablePause() {
    (C.kissSimEnablePause())
}
func kissSimPause() {
    (C.kissSimPause())
}*/
func left_button() int {
    return int(C.left_button())
}
/*func log10(num float32) float32 {
    return float32(C.log10(C.float(num)))
}
func log(num float32) float32 {
    return float32(C.log(C.float(num)))
}*/
func mav(m int, vel int) {
    (C.mav(C.int(m), C.int(vel)))
}
func motor(m int, p int) {
    (C.motor(C.int(m), C.int(p)))
}
func move_at_velocity(m int, vel int) {
    (C.move_at_velocity(C.int(m), C.int(vel)))
}
func move_relative_position(m int, speed int, pos int) {
    (C.move_relative_position(C.int(m), C.int(speed), C.int(pos)))
}
func move_to_position(m int, speed int, pos int) {
    (C.move_to_position(C.int(m), C.int(speed), C.int(pos)))
}
func mrp(m int, vel int, pos int) {
    (C.mrp(C.int(m), C.int(vel), C.int(pos)))
}
func mtp(m int, vel int, pos int) {
    (C.mtp(C.int(m), C.int(vel), C.int(pos)))
}
func msleep(msec int) {
    (C.msleep(C.long(msec)))
}
func off(m int) {
    (C.off(C.int(m)))
}
func power_level() float32 {
    return float32(C.power_level())
}

/*func s() char {
    return char(C.s())
}
func r_button() int {
    return int(C.r_button())
}
func random(m int) int {
    return int(C.random(C.int(m)))
}
*/
func right_button() int {
    return int(C.right_button())
}

/*func run_for() {
    (C.run_for())
}
func sec() float32 {
    return float32(C.sec())
}*/
func seconds() float32 {
    return float32(C.seconds())
}

/*func set_analog_float32s(mask int) {
    (C.set_analog_float32s(C.int(mask)))
}
func set_each_analog_state() {
    (C.set_each_analog_state())
}
func set_digital_output_value(port int, value int) {
    (C.set_digital_output_value(C.int(port), C.int(value)))
}*/
func set_pid_gains(motor int, p int, i int, d int, pd int, id int, dd int) {
    (C.set_pid_gains(C.int(motor), C.short(p), C.short(i), C.short(d), C.short(pd), C.short(id), C.short(dd)))
}
/*func motor() int {
    return int(C.motor())
}*/

/*func p() int {
    return int(C.p())
}

func i() int {
    return int(C.i())
}
func pd() int {
    return int(C.pd())
}
func id() int {
    return int(C.id())
}*/
func set_servo_position(srv int, pos int) {
    (C.set_servo_position(C.int(srv), C.int(pos)))
}
func setpwm(m int, dutycycle int) int {
    return int(C.setpwm(C.int(m), C.int(dutycycle)))
}
/*func sin(angle float32) float32 {
    return float32(C.sin(C.float(angle)))
}
func sleep(sec float32) {
    (C.sleep(C.float(sec)))
}
func sonar() int {
    return int(C.sonar())
}
func sqrt(num float32) float32 {
    return float32(C.sqrt(C.float(num)))
}

func start_process() int {
    return int(C.start_process())
}
func tan(angle float32) float32 {
    return float32(C.tan(C.float(angle)))
}
func up_button() int {
    return int(C.up_button())
}*/
func create_connect() int {
    return int(C.create_connect())
}
func create_disconnect() {
    (C.create_disconnect())
}
func create_start() {
    (C.create_start())
}
func create_passive() {
    (C.create_passive())
}
func create_safe() {
    (C.create_safe())
}
func create_full() {
    (C.create_full())
}
func create_spot() {
    (C.create_spot())
}
func create_cover() {
    (C.create_cover())
}
func create_demo(d int) {
    (C.create_demo(C.int(d)))
}
func create_cover_dock() {
    (C.create_cover_dock())
}
/*func get_create_mode(lag float32) int {
    return int(C.get_create_mode(C.float(lag)))
}
func get_create_lbump(lag float32) int {
    return int(C.get_create_lbump(C.float(lag)))
}
func get_create_rbump(lag float32) int {
    return int(C.get_create_rbump(C.float(lag)))
}
func get_create_lwdrop(lag float32) int {
    return int(C.get_create_lwdrop(C.float(lag)))
}
func get_create_cwdrop(lag float32) int {
    return int(C.get_create_cwdrop(C.float(lag)))
}
func get_create_rlwdrop(lag float32) int {
    return int(C.get_create_rlwdrop(C.float(lag)))
}
func get_create_wall(lag float32) int {
    return int(C.get_create_wall(C.float(lag)))
}
func get_create_lcliff(lag float32) int {
    return int(C.get_create_lcliff(C.float(lag)))
}
func get_create_lfcliff(lag float32) int {
    return int(C.get_create_lfcliff(C.float(lag)))
}
func get_create_rfcliff(lag float32) int {
    return int(C.get_create_rfcliff(C.float(lag)))
}
func get_create_rcliff(lag float32) int {
    return int(C.get_create_rcliff(C.float(lag)))
}
func get_create_vwall(lag float32) int {
    return int(C.get_create_vwall(C.float(lag)))
}
func get_create_overcurrents(lag float32) int {
    return int(C.get_create_overcurrents(C.float(lag)))
}
func get_create_infrared(lag float32) int {
    return int(C.get_create_infrared(C.float(lag)))
}
func get_create_advance_button(lag float32) int {
    return int(C.get_create_advance_button(C.float(lag)))
}
func get_create_play_button(lag float32) int {
    return int(C.get_create_play_button(C.float(lag)))
}
func get_create_distance(lag float32) int {
    return int(C.get_create_distance(C.float(lag)))
}
func set_create_distance(dist int) {
    (C.set_create_distance(C.int(dist)))
}
func get_create_normalized_angle(lag float32) int {
    return int(C.get_create_normalized_angle(C.float(lag)))
}
func get_create_total_angle(lag float32) int {
    return int(C.get_create_total_angle(C.float(lag)))
}
func set_create_normalized_angle(angle int) {
    (C.set_create_normalized_angle(C.int(angle)))
}
func set_create_total_angle(angle int) {
    (C.set_create_total_angle(C.int(angle)))
}
*/
/*func get_create_charging_state(lag float32) int {
    return int(C.get_create_charging_state(C.float(lag)))
}*/
/*func get_create_battery_voltage(lag float32) int {
    return int(C.get_create_battery_voltage(C.float(lag)))
}
func get_create_battery_current(lag float32) int {
    return int(C.get_create_battery_current(C.float(lag)))
}
func get_create_battery_temp(lag float32) int {
    return int(C.get_create_battery_temp(C.float(lag)))
}
func get_create_battery_charge(lag float32) int {
    return int(C.get_create_battery_charge(C.float(lag)))
}
func get_create_battery_capacity(lag float32) int {
    return int(C.get_create_battery_capacity(C.float(lag)))
}
func get_create_wall_amt(lag float32) int {
    return int(C.get_create_wall_amt(C.float(lag)))
}
func get_create_lcliff_amt(lag float32) int {
    return int(C.get_create_lcliff_amt(C.float(lag)))
}
func get_create_lfcliff_amt(lag float32) int {
    return int(C.get_create_lfcliff_amt(C.float(lag)))
}
func get_create_rfcliff_amt(lag float32) int {
    return int(C.get_create_rfcliff_amt(C.float(lag)))
}
func get_create_rcliff_amt(lag float32) int {
    return int(C.get_create_rcliff_amt(C.float(lag)))
}
func get_create_bay_DI(lag float32) int {
    return int(C.get_create_bay_DI(C.float(lag)))
}
func get_create_bay_AI(lag float32) int {
    return int(C.get_create_bay_AI(C.float(lag)))
}
func get_create_song_number(lag float32) int {
    return int(C.get_create_song_number(C.float(lag)))
}
func get_create_song_playing(lag float32) int {
    return int(C.get_create_song_playing(C.float(lag)))
}
func get_create_number_of_stream_packets(lag float32) int {
    return int(C.get_create_number_of_stream_packets(C.float(lag)))
}
func get_create_requested_velocity(lag float32) int {
    return int(C.get_create_requested_velocity(C.float(lag)))
}
func get_create_requested_radius(lag float32) int {
    return int(C.get_create_requested_radius(C.float(lag)))
}
func get_create_requested_right_velocity(lag float32) int {
    return int(C.get_create_requested_right_velocity(C.float(lag)))
}
func get_create_requested_left_velocity(lag float32) int {
    return int(C.get_create_requested_left_velocity(C.float(lag)))
}*/
func create_stop() {
    (C.create_stop())
}
func create_drive(speed int, radius int) {
    (C.create_drive(C.int(speed), C.int(radius)))
}
func create_drive_straight(speed int) {
    (C.create_drive_straight(C.int(speed)))
}
func create_spin_CW(speed int) {
    (C.create_spin_CW(C.int(speed)))
}
func create_spin_CCW(speed int) {
    (C.create_spin_CCW(C.int(speed)))
}
func create_drive_direct(r_speed int, l_speed int) {
    (C.create_drive_direct(C.int(r_speed), C.int(l_speed)))
}
func create_advance_led(on int) {
    (C.create_advance_led(C.int(on)))
}
func create_play_led(on int) {
    (C.create_play_led(C.int(on)))
}
func create_power_led(color int, brightness int) {
    (C.create_power_led(C.int(color), C.int(brightness)))
}
func create_load_song(num int) {
    (C.create_load_song(C.int(num)))
}
func create_play_song(num int) {
    (C.create_play_song(C.int(num)))
}

/*func track_update() {
    (C.track_update())
}
func track_get_frame() int {
    return int(C.track_get_frame())
}
func track_count(ch int) int {
    return int(C.track_count(C.int(ch)))
}
func track_size(ch int, i int) int {
    return int(C.track_size(C.int(ch), C.int(i)))
}
func track_x(ch int, i int) int {
    return int(C.track_x(C.int(ch), C.int(i)))
}
func track_y(ch int, i int) int {
    return int(C.track_y(C.int(ch), C.int(i)))
}
func track_confidence(ch int, i int) int {
    return int(C.track_confidence(C.int(ch), C.int(i)))
}
func track_bbox_left(ch int, i int) int {
    return int(C.track_bbox_left(C.int(ch), C.int(i)))
}
func track_bbox_right(ch int, i int) int {
    return int(C.track_bbox_right(C.int(ch), C.int(i)))
}
func track_bbox_top(ch int, i int) int {
    return int(C.track_bbox_top(C.int(ch), C.int(i)))
}
func track_bbox_bottom(ch int, i int) int {
    return int(C.track_bbox_bottom(C.int(ch), C.int(i)))
}
func track_bbox_width(ch int, i int) int {
    return int(C.track_bbox_width(C.int(ch), C.int(i)))
}
func track_bbox_height(ch int, i int) int {
    return int(C.track_bbox_height(C.int(ch), C.int(i)))
}
func track_angle(ch int, i int) int {
    return int(C.track_angle(C.int(ch), C.int(i)))
}
func track_major_axis(ch int, i int) int {
    return int(C.track_major_axis(C.int(ch), C.int(i)))
}
func track_minor_axis(ch int, i int) int {
    return int(C.track_minor_axis(C.int(ch), C.int(i)))
}*/
