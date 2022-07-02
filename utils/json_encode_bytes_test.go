package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONEncodeBytes(t *testing.T) {
	encodingTests := []struct {
		value []byte
	}{
		{
			value: []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x00, 0x1f, 0x00, 0x8b},
		},
		{
			value: []byte{0x7b, 0x0a, 0x20, 0x20, 0x22, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x3a, 0x20, 0x22, 0x22, 0x0a, 0x7d, 0x0a},
		},
	}

	for i, tt := range encodingTests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			encoded, err := JSONEncodeBytes(tt.value)
			if err != nil {
				t.Fatal(err)
			}

			_, err = json.Marshal(encoded)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}