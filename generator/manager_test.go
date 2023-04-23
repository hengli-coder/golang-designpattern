package generator

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func Test_manage_BuildHouse(t *testing.T) {
	tests := []struct {
		name   string
		fields string
	}{
		{name: "ice", fields: "ice"},
		{name: "wood", fields: "wood"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManage(getGenerator(tt.fields))
			got := m.BuildHouse()
			spew.Dump(got)
		})
	}
}
