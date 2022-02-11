package schemacan

import (
	"github.com/AutonoMooseTech/canbus/golang/canbus"
	"gopkg.in/yaml.v2"
	"bytes"
	"log"
	"fmt"
)

func splitStream(in []byte) (out [][]byte) {
	return bytes.Split(in, []byte("---\n"));
}

func Unmarshal(in []byte) (err error, output []interface{}) {
	for _, slice := range splitStream(in) {
		log.Printf("Input slice: %q", slice)

		var object canbus.Object

		yaml.Unmarshal(slice, &object);

		switch object.Kind {
		case "slot":
			var slot canbus.SLOT
			yaml.Unmarshal(slice, &slot)
			slot.Validate()
			output = append(output, &slot)
		case "message":
			var message canbus.Message
			yaml.Unmarshal(slice, &message)
			message.Validate()
			output = append(output, &message)
		default:
			return fmt.Errorf("Object kind not recognised: %s", object.Kind), nil
		}
	}

	return err, output
}
