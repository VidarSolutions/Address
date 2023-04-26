package Address
import (
	"Math/Big"

)
const (
	// HashLength is the expected length of the hash
		HashLength = 32
	// AddressLength is the expected length of the address
		AddressLength = 20
)
type Address struct{
	address 	[AddressLength]byte
	toString  	string
}

type uint256 struct {
	value *big.Int
}


var(
	 Addresses 			map[Address]address         	//Known Addresses
)


func getAddress(adds Address)(map[Address]address)	{
	return Addresses[adds]

}