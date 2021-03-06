package network

import (
	"github.com/hyperledger/fabric-gateway/pkg/hash"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
)

// signingIdentity is a wrapper around an identity and signing function that implements to Fabric protoutil.Signer
// interface
type signingIdentity struct {
	serializedID []byte
	sign         identity.Sign
	hash         hash.Hash
}

// newSigningIdentity creates an implementation ot the Fabric protoutil.Signer interface from an identity and signing
// function
func newSigningIdentity(id identity.Identity, sign identity.Sign, hash hash.Hash) (*signingIdentity, error) {
	serializedID, err := identity.Serialize(id)
	if err != nil {
		return nil, err
	}

	result := &signingIdentity{
		serializedID: serializedID,
		sign:         sign,
		hash:         hash,
	}
	return result, nil
}

func (signingID *signingIdentity) Sign(message []byte) ([]byte, error) {
	digest, err := signingID.hash(message)
	if err != nil {
		return nil, err
	}

	return signingID.sign(digest)
}

func (signingID *signingIdentity) Serialize() ([]byte, error) {
	return signingID.serializedID, nil
}
