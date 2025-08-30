package types

const (
	// ModuleName defines the module name
	ModuleName = "voter"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_voter"

	// Key prefixes

	// key: "Poll/value/1"  → value: {title: "Best Color?", options: ["red","blue","green"], creator: "alice"}
	// key: "Poll/value/2"  → value: {title: "Favorite Food?", options: ["pizza","sushi"], creator: "bob"}
	PollKey = "Poll/value/"

	// key: "Poll/count/" → value: 2
	PollCountKey = "Poll/count/"

	// 	key: "Vote/value/1/alice" → value: "red"
	// key: "Vote/value/1/bob"   → value: "blue"
	// 	key: "Vote/value/2/alice" → value: "red"
	// key: "Vote/value/2/bob"   → value: "blue"
	VoteKey = "Vote/value/"

	// 	key: "Vote/count/1 → value: 5
	// 	key: "Vote/count/2/red"   → value: 1
	// key: "Vote/count/2/blue"  → value: 2
	// key: "Vote/count/2/green" → value: 0
	VoteCountKey = "Vote/count/"
)

var (
	ParamsKey = []byte("p_voter")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
