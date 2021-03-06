// +build freebsd
// +build arm64
// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs disk/types_freebsd.go

package disk

const (
	sizeofPtr        = 0x8
	sizeofShort      = 0x2
	sizeofInt        = 0x4
	sizeofLong       = 0x8
	sizeofLongLong   = 0x8
	sizeofLongDouble = 0x8

	DEVSTAT_NO_DATA = 0x00
	DEVSTAT_READ    = 0x01
	DEVSTAT_WRITE   = 0x02
	DEVSTAT_FREE    = 0x03
)

const (
	sizeOfDevstat = 0x120
)

type (
	_C_short       int16
	_C_int         int32
	_C_long        int64
	_C_long_long   int64
	_C_long_double int64
)

type Devstat struct {
	Sequence0     uint32
	Allocated     int32
	Start_count   uint32
	End_count     uint32
	Busy_from     Bintime
	Dev_links     _Ctype_struct___0
	Device_number uint32
	Device_name   [16]int8
	Unit_number   int32
	Bytes         [4]uint64
	Operations    [4]uint64
	Duration      [4]Bintime
	Busy_time     Bintime
	Creation_time Bintime
	Block_size    uint32
	Tag_types     [3]uint64
	Flags         uint32
	Device_type   uint32
	Priority      uint32
	Id            *byte
	Sequence1     uint32
	Pad_cgo_0     [4]byte
}
type Bintime struct {
	Sec  int64
	Frac uint64
}

type _Ctype_struct___0 struct {
	Empty uint64
}
