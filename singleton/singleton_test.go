package singleton

import (
	"reflect"
	"testing"
)

func Test_getSingleInstance(t *testing.T) {
	s := getSingleInstance()
	tests := []struct {
		name string
		want *single
	}{
		{"should get the same instance", s},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSingleInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSingleInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
