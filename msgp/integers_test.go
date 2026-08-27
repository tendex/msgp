package msgp

import (
	"encoding/binary"
	"testing"
)

func BenchmarkIntegers(b *testing.B) {
	bytes64 := make([]byte, 9)
	bytes32 := make([]byte, 5)
	bytes16 := make([]byte, 3)

	b.Run("Int64", func(b *testing.B) {
		b.Run("Put", func(b *testing.B) {
			for b.Loop() {
				putMint64(bytes64, -1234567890123456789)
			}
		})
		b.Run("Get", func(b *testing.B) {
			putMint64(bytes64, -1234567890123456789)
			for b.Loop() {
				getMint64(bytes64)
			}
		})
	})
	b.Run("Int32", func(b *testing.B) {
		b.Run("Put", func(b *testing.B) {
			for b.Loop() {
				putMint32(bytes32, -123456789)
			}
		})
		b.Run("Get", func(b *testing.B) {
			putMint32(bytes32, -123456789)
			for b.Loop() {
				getMint32(bytes32)
			}
		})
	})
	b.Run("Int16", func(b *testing.B) {
		b.Run("Put", func(b *testing.B) {
			for b.Loop() {
				putMint16(bytes16, -12345)
			}
		})
		b.Run("Get", func(b *testing.B) {
			putMint16(bytes16, -12345)
			for b.Loop() {
				getMint16(bytes16)
			}
		})
	})

	b.Run("Uint64", func(b *testing.B) {
		b.Run("Put", func(b *testing.B) {
			for b.Loop() {
				putMuint64(bytes64, 1234567890123456789)
			}
		})
		b.Run("Get", func(b *testing.B) {
			putMuint64(bytes64, 1234567890123456789)
			for b.Loop() {
				getMuint64(bytes64)
			}
		})
	})
	b.Run("Uint32", func(b *testing.B) {
		b.Run("Put", func(b *testing.B) {
			for b.Loop() {
				putMuint32(bytes32, 123456789)
			}
		})
		b.Run("Get", func(b *testing.B) {
			putMuint32(bytes32, 123456789)
			for b.Loop() {
				getMuint32(bytes32)
			}
		})
	})
	b.Run("Uint16", func(b *testing.B) {
		b.Run("Put", func(b *testing.B) {
			for b.Loop() {
				putMuint16(bytes16, 12345)
			}
		})
		b.Run("Get", func(b *testing.B) {
			putMuint16(bytes16, 12345)
			for b.Loop() {
				getMuint16(bytes16)
			}
		})
	})
}

func BenchmarkIntegersUnix(b *testing.B) {
	bytes := make([]byte, 12)
	var sec int64 = 1609459200
	var nsec int32 = 123456789

	b.Run("Get", func(b *testing.B) {
		binary.BigEndian.PutUint64(bytes, uint64(sec))
		binary.BigEndian.PutUint32(bytes[8:], uint32(nsec))
		for b.Loop() {
			getUnix(bytes)
		}
	})

	b.Run("Put", func(b *testing.B) {
		for b.Loop() {
			putUnix(bytes, sec, nsec)
		}
	})
}

func BenchmarkIntegersPrefix(b *testing.B) {
	bytesU16 := make([]byte, 3)
	bytesU32 := make([]byte, 5)
	bytesU64 := make([]byte, 9)

	b.Run("u16", func(b *testing.B) {
		var pre byte = 0x01
		var sz uint16 = 12345
		for b.Loop() {
			prefixu16(bytesU16, pre, sz)
		}
	})
	b.Run("u32", func(b *testing.B) {
		var pre byte = 0x02
		var sz uint32 = 123456789
		for b.Loop() {
			prefixu32(bytesU32, pre, sz)
		}
	})
	b.Run("u64", func(b *testing.B) {
		var pre byte = 0x03
		var sz uint64 = 1234567890123456789
		for b.Loop() {
			prefixu64(bytesU64, pre, sz)
		}
	})
}
