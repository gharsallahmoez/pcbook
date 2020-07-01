package serializer_test

import (
	"github.com/gharsallahmoez/pcbook/pb"
	"github.com/gharsallahmoez/pcbook/sample"
	"github.com/gharsallahmoez/pcbook/serializer"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"
	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobugToBinaryFile(laptop1,binaryFile)
	require.NoError(t,err)
	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile,laptop2)
	require.NoError(t,err)
	require.True(t,proto.Equal(laptop1,laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1,jsonFile)
	require.NoError(t,err)
}
