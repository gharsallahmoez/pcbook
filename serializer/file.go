// serialize laptop object to file
package serializer

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

// write protobuf message to a file in binary format
func WriteProtobugToBinaryFile(message proto.Message, filename string) error {
	// serialize message to binary
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary : %w", err)
	}
	// save the data to specified file name
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file : %w", err)
	}
	return nil
}

// ReadProtobufFromBinaryFile reads protocol buffer message from binary file
func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file : %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message : %w", err)
	}
	return nil
}

// WriteProtobufToJSONFile writes proto messages to json file
func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	marshal := jsonpb.Marshaler{
		Indent :"	",
	}
	json, err := marshal.MarshalToString(message)
	if err != nil {
		return fmt.Errorf("unable to marshal proto message")
	}
	err = ioutil.WriteFile(filename, []byte(json), 0644)
	if err != nil {
		return fmt.Errorf("unable to write file %w", err)
	}
	return nil
}
