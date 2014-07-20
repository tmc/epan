package epan

/*
#cgo pkg-config: glib-2.0
#cgo CFLAGS: -Iwireshark 
#cgo LDFLAGS: -lwireshark -lwsutil

#include "wireshark/config.h"
#include "epan/epan.h"
#include "epan/proto.h"
#include "wsutil/plugins.h"
#include "wsutil/privileges.h"

*/
import "C"
import "fmt"
import "unsafe"

type Protocol struct {
	p *C.struct__protocol
}

func (p Protocol) ID() int {
	return int(C.proto_get_id(p.p))
}

func (p Protocol) Name() string {
	return C.GoString(C.proto_get_protocol_long_name(p.p))
}

func (p Protocol) ShortName() string {
	return C.GoString(C.proto_get_protocol_short_name(p.p))
}

func (p Protocol) String() string {
	return fmt.Sprintf("%s - %s", p.Name(), p.FilterName())
}

func (p Protocol) Enabled() bool {
	return cbool(C.proto_is_protocol_enabled(p.p))
}

func (p Protocol) FilterName() string {
	return C.GoString(C.proto_get_protocol_filter_name(C.int(p.ID())))
}

func ListProtocols() []Protocol {
	result := []Protocol{}
	var i C.int
	var j int
	cookie := unsafe.Pointer(&j)
	i = C.proto_get_first_protocol(&cookie)
	for i != -1 {
		p := Protocol{C.find_protocol_by_id(i)}
		result = append(result, p)
		i = C.proto_get_next_protocol(&cookie)
	}
	return result

}
