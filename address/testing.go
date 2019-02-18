package address

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/filecoin-project/go-filecoin/bls-signatures"
	"github.com/filecoin-project/go-filecoin/crypto"
)

// newSecp256k1ForTestGetter returns a closure that returns an address unique to that invocation.
// The address is unique wrt the closure returned, not globally.
func newSecp256k1ForTestGetter() func() Address {
	return func() Address {
		sk, err := crypto.GenerateKey()
		if err != nil {
			panic(err)
		}

		pk, ok := sk.Public().(*ecdsa.PublicKey)
		if !ok {
			panic("here")
		}

		newAddr, err := NewFromSECP256K1(Testnet, pk)
		if err != nil {
			panic(err)
		}

		return newAddr
	}
}

// newActorForTestGetter returns a closure that returns an address unique to that invocation.
// The address is unique wrt the closure returned, not globally.
func newActorForTestGetter() func() Address {
	return func() Address {
		i := 0
		s := fmt.Sprintf("address%d", i)
		i++
		newAddr, err := NewFromActor(Testnet, []byte(s))
		if err != nil {
			panic(err)
		}
		return newAddr
	}
}

// newActorIDForTestGetter returns a closure that returns an address unique to that invocation.
// The address is unique wrt the closure returned, not globally.
func newActorIDForTestGetter() func() Address {
	return func() Address {
		i := 0
		i++
		newAddr, err := NewFromActorID(Testnet, uint64(i))
		if err != nil {
			panic(err)
		}
		return newAddr
	}
}

// newBLSForTestGetter returns a closure that returns an address unique to that invocation.
// The address is unique wrt the closure returned, not globally.
func newBLSForTestGetter() func() Address {
	return func() Address {
		blsAddress, err := NewFromBLS(Testnet, bls.PrivateKeyPublicKey((bls.PrivateKeyGenerate())))
		if err != nil {
			panic(err)
		}
		return blsAddress
	}
}
