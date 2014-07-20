package epan

/*
#cgo pkg-config: glib-2.0
#cgo CFLAGS: -I/usr/local/include/wireshark -I/usr/include/wireshark
#cgo LDFLAGS: -lwireshark -lwsutil

#include "wireshark/config.h"
#include "epan/epan.h"
#include "epan/proto.h"
#include "wsutil/plugins.h"
#include "wsutil/privileges.h"

void epan_setup(void) {
  init_process_policies();
  relinquish_special_privs_perm();
  epan_register_plugin_types();
  scan_plugins();
  epan_init(register_all_protocols, register_all_protocol_handoffs, NULL, NULL);
}

*/
import "C"

func Version() string {
	return C.GoString((*C.char)(C.epan_get_version()))
}

func init() {
	C.epan_setup()
}
