package msgp

import (
	"testing"
)

func BenchmarkReadWriteFloat32(b *testing.B) {
	var f float32 = 3.9081
	bts := AppendFloat32([]byte{}, f)
	for b.Loop() {
		bts = AppendFloat32(bts[0:0], f)
		f, bts, _ = ReadFloat32Bytes(bts)
	}
}

func BenchmarkReadWriteFloat64(b *testing.B) {
	var f = 3.9081
	bts := AppendFloat64([]byte{}, f)
	for b.Loop() {
		bts = AppendFloat64(bts[0:0], f)
		f, bts, _ = ReadFloat64Bytes(bts)
	}
}
