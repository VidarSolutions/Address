package address

import (
	"math/big"
)

const (
	HashLength    = 32
	AddressLength = 20
)

type Address struct {
	Address  [AddressLength]byte

	ToString 	string
}

type uint256 struct {
	value *big.Int
}

type addressBook struct {
	addresses map[Address]Address
}

var AddressBook = addressBook{
	addresses: make(map[Address]Address),
}

func (ab *addressBook) GetAddresses() map[Address]Address {
	return ab.addresses
}

func (ab *addressBook) AddAddress(addr Address) {
	ab.addresses[addr] = addr
}
