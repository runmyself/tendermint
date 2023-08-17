package bls12381

import (
	blst "github.com/prysmaticlabs/prysm/v4/crypto/bls/blst"
	blscommon "github.com/prysmaticlabs/prysm/v4/crypto/bls/common"
)


// SecretKeyFromBytes creates a BLS private key from a BigEndian byte slice.
func SecretKeyFromBytes(privKey []byte) (blscommon.SecretKey, error) {
	return blst.SecretKeyFromBytes(privKey)
}

// PublicKeyFromBytes creates a BLS public key from a  BigEndian byte slice.
func PublicKeyFromBytes(pubKey []byte) (blscommon.PublicKey, error) {
	return blst.PublicKeyFromBytes(pubKey)
}

// SignatureFromBytesNoValidation creates a BLS signature from a LittleEndian byte slice.
// It does not check validity of the signature, use only when the byte slice has
// already been verified
func SignatureFromBytesNoValidation(sig []byte) (blscommon.Signature, error) {
	return blst.SignatureFromBytesNoValidation(sig)
}

// SignatureFromBytes creates a BLS signature from a LittleEndian byte slice.
func SignatureFromBytes(sig []byte) (blscommon.Signature, error) {
	return blst.SignatureFromBytes(sig)
}

// MultipleSignaturesFromBytes creates a slice of BLS signatures from a LittleEndian 2d-byte slice.
func MultipleSignaturesFromBytes(sigs [][]byte) ([]blscommon.Signature, error) {
	return blst.MultipleSignaturesFromBytes(sigs)
}

// AggregatePublicKeys aggregates the provided raw public keys into a single key.
func AggregatePublicKeys(pubs [][]byte) (blscommon.PublicKey, error) {
	return blst.AggregatePublicKeys(pubs)
}

// AggregateMultiplePubkeys aggregates the provided decompressed keys into a single key.
func AggregateMultiplePubkeys(pubs []blscommon.PublicKey) blscommon.PublicKey {
	return blst.AggregateMultiplePubkeys(pubs)
}

// AggregateSignatures converts a list of signatures into a single, aggregated sig.
func AggregateSignatures(sigs []blscommon.Signature) blscommon.Signature {
	return blst.AggregateSignatures(sigs)
}

// AggregateCompressedSignatures converts a list of compressed signatures into a single, aggregated sig.
func AggregateCompressedSignatures(multiSigs [][]byte) (blscommon.Signature, error) {
	return blst.AggregateCompressedSignatures(multiSigs)
}

// VerifySignature verifies a single signature. For performance reason, always use VerifyMultipleSignatures if possible.
func VerifySignature(sig []byte, msg [32]byte, pubKey blscommon.PublicKey) (bool, error) {
	return blst.VerifySignature(sig, msg, pubKey)
}

// VerifyMultipleSignatures verifies multiple signatures for distinct messages securely.
func VerifyMultipleSignatures(sigs [][]byte, msgs [][32]byte, pubKeys []blscommon.PublicKey) (bool, error) {
	return blst.VerifyMultipleSignatures(sigs, msgs, pubKeys)
}

// NewAggregateSignature creates a blank aggregate signature.
func NewAggregateSignature() blscommon.Signature {
	return blst.NewAggregateSignature()
}

// RandKey creates a new private key using a random input.
func RandKey() (blscommon.SecretKey, error) {
	return blst.RandKey()
}
