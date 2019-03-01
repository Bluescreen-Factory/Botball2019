package main

import (
	"fmt"
        "mywall"
)

func main() {
	fmt.Println("press button");
	//mywall.Wait_for_a_button();
        mywall.Wait_for_any_button();
	fmt.Println(mywall.Right_button());
}

