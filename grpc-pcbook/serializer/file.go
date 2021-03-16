package serializer

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

//Writes Protocol Buffer message to binary file (ProtoMessage -> Binary)
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("Can't marshal proto message to binary: %w", err) //%w ~ %v (value in default format)
	}
	err = ioutil.WriteFile(filename, data, 0644) //ioutil.WriteFile(filename string, data []byte, perm os.FileMode)
	if err != nil {
		return fmt.Errorf("Can't write binary data to file: %w", err)
	}
	return nil
}

//Reads Protocol Buffer message from binary file (Binary -> ProtoMessage)
func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Can't read binary data from file: %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("Can't unmarshal binary data to proto message: %w", err)
	}
	return nil
}
