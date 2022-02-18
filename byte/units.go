package byte

const (
	// B byte (Unit)
	B BytesUnit = 1
	// KB Kilobyte (SI) - 1000 bytes
	KB BytesUnit = 1000
	// MB Megabyte (SI) - 1e6 bytes
	MB BytesUnit = 1e6
	// GB Gigabyte (SI) - 1e9 bytes
	GB BytesUnit = 1e9
	// TB Terabyte (SI) - 1e12 bytes
	TB BytesUnit = 1e12
	// PB Petabyte (SI) - 1e15 bytes
	PB BytesUnit = 1e15

	// KiB Kibibyte (NIST) - 1024 bytes
	KiB BytesUnit = 1024
	// MiB Mebibyte (NIST) - 2^20 bytes
	MiB = 1024 * KiB
	// GiB Gibibyte (NIST) - 2^30 bytes
	GiB = 1024 * MiB
	// TiB Tebibyte (NIST) - 2^40 bytes
	TiB = 1024 * GiB
	// PiB Pebibyte (NIST) - 2^50 bytes
	PiB = 1024 * TiB
)

var strToByteUnit = map[string]BytesUnit{
	"":    B,
	"B":   B,
	"KB":  KB,
	"KIB": KiB,
	"MB":  MB,
	"MIB": MiB,
	"GB":  GB,
	"GIB": GiB,
	"TB":  TB,
	"TIB": TiB,
	"PB":  PB,
	"PIB": PiB,
}

var byteUnitToStr = map[BytesUnit]string{
	B:   "B",
	KB:  "KB",
	KiB: "KiB",
	MB:  "MB",
	MiB: "MiB",
	GB:  "GB",
	GiB: "GiB",
	TB:  "TB",
	TiB: "TiB",
	PB:  "PB",
	PiB: "PiB",
}

var bytesUnitThresholds = []struct {
	unit BytesUnit
	less Bytes
}{
	{B, 0x400},
	{KiB, Bytes(uint64(MiB) >> 4)},
	{MiB, Bytes(uint64(GiB) >> 4)},
	{GiB, Bytes(uint64(TiB) >> 4)},
	{TiB, Bytes(uint64(PiB) >> 4)},
	{PiB, 0x7FFFFFFFFFFFFFFF},
}
var bytesSIUnitThresholds = []struct {
	unit BytesUnit
	less Bytes
}{
	{B, 5e2},
	{KB, 5e5},
	{MB, 5e8},
	{GB, 5e11},
	{TB, 5e14},
	{PB, 0x7FFFFFFFFFFFFFFF},
}
