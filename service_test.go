package architecture

import (
	"testing"

	"go.uber.org/mock/gomock"
)

func TestPut(t *testing.T) {
	ctl := gomock.NewController(t)
	acc := NewMockAccessor(ctl)
	p := Person{
		First: "Paul",
	}

	acc.EXPECT().Save(1, p).MinTimes(1).MaxTimes(1)

	Put(acc, 1, p)

	ctl.Finish()

}

//func ExamplePut() {
//	mdb := Db{}
//	p := Person{
//		First: "Paul",
//	}
//
//	Put(mdb, 1, p)
//	got := mdb.Retrieve(1)
//	fmt.Println(got)
//	// Output: {Paul}
//}
