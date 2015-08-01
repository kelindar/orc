// generated by stringer -type=RLEEncodingType; DO NOT EDIT

package orc

import "fmt"

const _RLEEncodingType_name = "RLEV2IntShortRepeatRLEV2IntDirectRLEV2IntPatchedBaseRLEV2IntDelta"

var _RLEEncodingType_index = [...]uint8{19, 33, 52, 65}

func (i RLEEncodingType) String() string {
	if i < 0 || i >= RLEEncodingType(len(_RLEEncodingType_index)) {
		return fmt.Sprintf("RLEEncodingType(%d)", i)
	}
	hi := _RLEEncodingType_index[i]
	lo := uint8(0)
	if i > 0 {
		lo = _RLEEncodingType_index[i-1]
	}
	return _RLEEncodingType_name[lo:hi]
}