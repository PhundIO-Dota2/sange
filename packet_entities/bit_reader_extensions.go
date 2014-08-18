package packet_entities

import (
	"github.com/dotabuff/sange/utils"
)

func ReadUpdateType(br *utils.BitReader) UpdateType {
	result := Preserve
	if !br.ReadBoolean() {
		if br.ReadBoolean() {
			result = Create
		}
	} else {
		result = Leave
		if br.ReadBoolean() {
			result = Delete
		}
	}
	return result
}
