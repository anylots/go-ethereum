// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"

	"github.com/scroll-tech/go-ethereum/common/hexutil"
)

var _ = (*blsDataMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (b BLSData) MarshalJSON() ([]byte, error) {
	type BLSData struct {
		BLSSigners   []hexutil.Bytes `json:"bls_signers"`
		BLSSignature hexutil.Bytes   `json:"bls_signature"`
	}
	var enc BLSData
	if b.BLSSigners != nil {
		enc.BLSSigners = make([]hexutil.Bytes, len(b.BLSSigners))
		for k, v := range b.BLSSigners {
			enc.BLSSigners[k] = v
		}
	}
	enc.BLSSignature = b.BLSSignature
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (b *BLSData) UnmarshalJSON(input []byte) error {
	type BLSData struct {
		BLSSigners   []hexutil.Bytes `json:"bls_signers"`
		BLSSignature *hexutil.Bytes  `json:"bls_signature"`
	}
	var dec BLSData
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.BLSSigners != nil {
		b.BLSSigners = make([][]byte, len(dec.BLSSigners))
		for k, v := range dec.BLSSigners {
			b.BLSSigners[k] = v
		}
	}
	if dec.BLSSignature != nil {
		b.BLSSignature = *dec.BLSSignature
	}
	return nil
}
