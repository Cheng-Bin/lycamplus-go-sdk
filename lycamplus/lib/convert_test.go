package lib

import (
	"testing"
)

func TestStruct2Map(t *testing.T) {
	StreamRequestModel := StreamRequestModel{}
	StreamRequestModel.Title = "zhangsan"
	StreamRequestModel.City = "ChengDu"

	Struct2Map(StreamRequestModel)
}
