package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

//proto.Message -> JSON
func ProtobufToJson(message proto.Message) (string, error) {
	//Custom configurations of Marshaller inside the scope bellow
	marshaler := protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		Indent:          "  ",
		UseProtoNames:   false,
	}
	data, err := marshaler.Marshal(message)
	return string(data), err
}
