package serializer

import (
	"testing"

	"br.com.rodrigodonizettio/grpc-pcbook/sample"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel() //Tests can run in parallel

	binaryFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)
}
