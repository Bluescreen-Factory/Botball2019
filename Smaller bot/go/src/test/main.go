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
*/

func main() {
	fmt.Println(analog(0))
}

func a_button() int {
	return C.a_button()
}
func alloff() {
	return C.alloff()
}
func analog(p int) int {
	return C.analog(p)
}
func analog10(p int) int {
	return C.analog10(p)
}
func ao() {
	return C.ao()
}
/*
func atan(angle float32) float32 {
	return C.atan(angle)
}
*/
func b_button() int {
	return C.b_button()
}
func beep() {
	return C.beep()
}
func bk(m int) {
	return C.bk(m)
}
func black_button() int {
	return C.black_button()
}
func block_motor_done(m int) {
	return C.block_motor_done(m)
}
func bmd(m int) {
	return C.bmd(m)
}

/*func col() int {
	return C.col()
}
func row() int {
	return C.row()
}
func s() char {
	return C.s()
}*/
func clear_motor_position_counter(motor_nbr int) {
	return C.clear_motor_position_counter(motor_nbr)
}
/*
func cos(angle float32) float32 {
	return C.cos(angle)
}
*/
func digital(p int) int {
	return C.digital(p)
}
func disable_servos() {
	return C.disable_servos()
}

/*func down_button() int {
	return C.down_button()
}*/
func enable_servos() {
	return C.enable_servos()
}

/*func exp10(num float32) float32 {
	return C.exp10(num)
}
func exp(num float32) float32 {
	return C.exp(num)
}
*/
func fd(m int) {
	return C.fd(m)
}
func freeze(m int) {
	return C.freeze(m)
}
func get_motor_done(m int) int {
	return C.get_motor_done(m)
}
func get_motor_position_counter(m int) int {
	return C.get_motor_position_counter(m)
}
func get_servo_position(srv int) int {
	return C.get_servo_position(srv)
}

/*func kill_process(pid int) {
	return C.kill_process(pid)
}
func kissSimEnablePause() {
	return C.kissSimEnablePause()
}
func kissSimPause() {
	return C.kissSimPause()
}*/
func left_button() int {
	return C.left_button()
}
/*func log10(num float32) float32 {
	return C.log10(num)
}
func log(num float32) float32 {
	return C.log(num)
}*/
func mav(m int, vel int) {
	return C.mav(m, vel)
}
func motor(m int, p int) {
	return C.motor(m, p)
}
func move_at_velocity(m int, vel int) {
	return C.move_at_velocity(m, vel)
}
func move_relative_position(m int, speed int, pos int) {
	return C.move_relative_position(m, speed, pos)
}
func move_to_position(m int, speed int, pos int) {
	return C.move_to_position(m, speed, pos)
}
func mrp(m int, vel int, pos int) {
	return C.mrp(m, vel, pos)
}
func mtp(m int, vel int, pos int) {
	return C.mtp(m, vel, pos)
}
func msleep(msec int) {
	return C.msleep(msec)
}
func off(m int) {
	return C.off(m)
}
func power_level() float32 {
	return C.power_level()
}

/*func s() char {
	return C.s()
}
func r_button() int {
	return C.r_button()
}
func random(m int) int {
	return C.random(m)
}
*/
func right_button() int {
	return C.right_button()
}

/*func run_for() {
	return C.run_for()
}
func sec() float32 {
	return C.sec()
}*/
func seconds() float32 {
	return C.seconds()
}

/*func set_analog_float32s(mask int) {
	return C.set_analog_float32s(mask)
}
func set_each_analog_state() {
	return C.set_each_analog_state()
}
func set_digital_output_value(port int, value int) {
	return C.set_digital_output_value(port, value)
}*/
func set_pid_gains(motor int, p int, i int, d int, pd int, id int, dd int) int {
	return C.set_pid_gains(motor, p, i, d, pd, id, dd)
}
func motor() int {
	return C.motor()
}

/*func p() int {
	return C.p()
}

func i() int {
	return C.i()
}
func pd() int {
	return C.pd()
}
func id() int {
	return C.id()
}*/
func set_servo_position(srv int, pos int) int {
	return C.set_servo_position(srv, pos)
}
func setpwm(m int, dutycycle int) int {
	return C.setpwm(m, dutycycle)
}
/*func sin(angle float32) float32 {
	return C.sin(angle)
}
func sleep(sec float32) {
	return C.sleep(sec)
}
func sonar() int {
	return C.sonar()
}
func sqrt(num float32) float32 {
	return C.sqrt(num)
}

func start_process() int {
	return C.start_process()
}
func tan(angle float32) float32 {
	return C.tan(angle)
}
func up_button() int {
	return C.up_button()
}*/
func create_connect() int {
	return C.create_connect()
}
func create_disconnect() {
	return C.create_disconnect()
}
func create_start() {
	return C.create_start()
}
func create_passive() {
	return C.create_passive()
}
func create_safe() {
	return C.create_safe()
}
func create_full() {
	return C.create_full()
}
func create_spot() {
	return C.create_spot()
}
func create_cover() {
	return C.create_cover()
}
func create_demo(d int) {
	return C.create_demo(d)
}
func create_cover_dock() {
	return C.create_cover_dock()
}
func get_create_mode(lag float32) int {
	return C.get_create_mode(lag)
}
func get_create_lbump(lag float32) int {
	return C.get_create_lbump(lag)
}
func get_create_rbump(lag float32) int {
	return C.get_create_rbump(lag)
}
func get_create_lwdrop(lag float32) int {
	return C.get_create_lwdrop(lag)
}
func get_create_cwdrop(lag float32) int {
	return C.get_create_cwdrop(lag)
}

/*func get_create_rlwdrop(lag float32) int {
	return C.get_create_rlwdrop(lag)
}*/
func get_create_wall(lag float32) int {
	return C.get_create_wall(lag)
}
func get_create_lcliff(lag float32) int {
	return C.get_create_lcliff(lag)
}
func get_create_lfcliff(lag float32) int {
	return C.get_create_lfcliff(lag)
}
func get_create_rfcliff(lag float32) int {
	return C.get_create_rfcliff(lag)
}
func get_create_rcliff(lag float32) int {
	return C.get_create_rcliff(lag)
}
func get_create_vwall(lag float32) int {
	return C.get_create_vwall(lag)
}
func get_create_overcurrents(lag float32) int {
	return C.get_create_overcurrents(lag)
}
func get_create_infrared(lag float32) int {
	return C.get_create_infrared(lag)
}
func get_create_advance_button(lag float32) int {
	return C.get_create_advance_button(lag)
}
func get_create_play_button(lag float32) int {
	return C.get_create_play_button(lag)
}
func get_create_distance(lag float32) int {
	return C.get_create_distance(lag)
}
func set_create_distance(dist int) {
	return C.set_create_distance(dist)
}
func get_create_normalized_angle(lag float32) int {
	return C.get_create_normalized_angle(lag)
}
func get_create_total_angle(lag float32) int {
	return C.get_create_total_angle(lag)
}
func set_create_normalized_angle(angle int) {
	return C.set_create_normalized_angle(angle)
}
func set_create_total_angle(angle int) {
	return C.set_create_total_angle(angle)
}

/*func get_create_charging_state(lag float32) int {
	return C.get_create_charging_state(lag)
}*/
func get_create_battery_voltage(lag float32) int {
	return C.get_create_battery_voltage(lag)
}
func get_create_battery_current(lag float32) int {
	return C.get_create_battery_current(lag)
}
func get_create_battery_temp(lag float32) int {
	return C.get_create_battery_temp(lag)
}
func get_create_battery_charge(lag float32) int {
	return C.get_create_battery_charge(lag)
}
func get_create_battery_capacity(lag float32) int {
	return C.get_create_battery_capacity(lag)
}
func get_create_wall_amt(lag float32) int {
	return C.get_create_wall_amt(lag)
}
func get_create_lcliff_amt(lag float32) int {
	return C.get_create_lcliff_amt(lag)
}
func get_create_lfcliff_amt(lag float32) int {
	return C.get_create_lfcliff_amt(lag)
}
func get_create_rfcliff_amt(lag float32) int {
	return C.get_create_rfcliff_amt(lag)
}
func get_create_rcliff_amt(lag float32) int {
	return C.get_create_rcliff_amt(lag)
}
func get_create_bay_DI(lag float32) int {
	return C.get_create_bay_DI(lag)
}
func get_create_bay_AI(lag float32) int {
	return C.get_create_bay_AI(lag)
}
func get_create_song_number(lag float32) int {
	return C.get_create_song_number(lag)
}
func get_create_song_playing(lag float32) int {
	return C.get_create_song_playing(lag)
}
func get_create_number_of_stream_packets(lag float32) int {
	return C.get_create_number_of_stream_packets(lag)
}
func get_create_requested_velocity(lag float32) int {
	return C.get_create_requested_velocity(lag)
}
func get_create_requested_radius(lag float32) int {
	return C.get_create_requested_radius(lag)
}
func get_create_requested_right_velocity(lag float32) int {
	return C.get_create_requested_right_velocity(lag)
}
func get_create_requested_left_velocity(lag float32) int {
	return C.get_create_requested_left_velocity(lag)
}
func create_stop() {
	return C.create_stop()
}
func create_drive(speed int, radius int) {
	return C.create_drive(speed, radius)
}
func create_drive_straight(speed int) {
	return C.create_drive_straight(speed)
}
func create_spin_CW(speed int) {
	return C.create_spin_CW(speed)
}
func create_spin_CCW(speed int) {
	return C.create_spin_CCW(speed)
}
func create_drive_direct(r_speed int, l_speed int) {
	return C.create_drive_direct(r_speed, l_speed)
}
func create_advance_led(on int) {
	return C.create_advance_led(on)
}
func create_play_led(on int) {
	return C.create_play_led(on)
}
func create_power_led(color int, brightness int) {
	return C.create_power_led(color, brightness)
}
func create_load_song(num int) {
	return C.create_load_song(num)
}
func create_play_song(num int) {
	return C.create_play_song(num)
}

/*func track_update() {
	return C.track_update()
}
func track_get_frame() int {
	return C.track_get_frame()
}
func track_count(ch int) int {
	return C.track_count(ch)
}
func track_size(ch int, i int) int {
	return C.track_size(ch, i)
}
func track_x(ch int, i int) int {
	return C.track_x(ch, i)
}
func track_y(ch int, i int) int {
	return C.track_y(ch, i)
}
func track_confidence(ch int, i int) int {
	return C.track_confidence(ch, i)
}
func track_bbox_left(ch int, i int) int {
	return C.track_bbox_left(ch, i)
}
func track_bbox_right(ch int, i int) int {
	return C.track_bbox_right(ch, i)
}
func track_bbox_top(ch int, i int) int {
	return C.track_bbox_top(ch, i)
}
func track_bbox_bottom(ch int, i int) int {
	return C.track_bbox_bottom(ch, i)
}
func track_bbox_width(ch int, i int) int {
	return C.track_bbox_width(ch, i)
}
func track_bbox_height(ch int, i int) int {
	return C.track_bbox_height(ch, i)
}
func track_angle(ch int, i int) int {
	return C.track_angle(ch, i)
}
func track_major_axis(ch int, i int) int {
	return C.track_major_axis(ch, i)
}
func track_minor_axis(ch int, i int) int {
	return C.track_minor_axis(ch, i)
}*/
