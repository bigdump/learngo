package main

type Address [32]byte

//NonceChange represents nonce changes of individual accounts.
type NonceChange struct {
	Account Address `json:"account,omitempty"`
	Prev    uint64  `json:"prev,omitempty"`
	New     uint64  `json:"new,omitempty"`
	Type    string  `json:"type,omitempty"`
}

// NonceChange represents nonce changes of individual accounts.

//type NonceChange struct {
//	Account Address `json:"account,omitempty"`
//	Prev    uint64          `json:"prev,omitempty"`
//	Type    string          `json:"type,omitempty"`
//}

func main() {
	//non := &NonceChange{
	//	Account: Address{12, 12, 3, 31},
	//	Prev:    1,
	//	Type:    "12312312",
	//}
	//b, _ := json.Marshal(non)
	//fmt.Println(common.BytesToHex(b))


}
