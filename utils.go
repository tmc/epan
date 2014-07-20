package epan

/*
#cgo pkg-config: glib-2.0
#include "epan/epan.h"
#include "wsutil/privileges.h"
*/
import "C"
import (
	"fmt"
	"os/user"
)

//export print_current_user
func print_current_user() {
	u, err := user.Current()
	if err != nil {
		return
	}
	fmt.Println("running as", u.Username)
	if (C.int)(C.running_with_special_privs()) == 1 {
		fmt.Println("could be dangerous.")
	}
}

func cbool(b C.gboolean) bool {
	if int(C.int(b)) == 0 {
		return false
	}
	return true
}
