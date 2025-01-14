package types

const (
	// ModuleName defines the module name
	ModuleName = "memoir"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_memoir"
)

var (
	ParamsKey = []byte("p_memoir")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
