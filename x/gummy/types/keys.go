package types

const (
	// ModuleName defines the module name
	ModuleName = "gummy"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_gummy"
)

var (
	ParamsKey = []byte("p_gummy")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
