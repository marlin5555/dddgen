package entity

import (
	"reflect"
	"testing"
)

func TestRefEntities(t *testing.T) {
	type args struct {
		group *Group
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{group: &Group{EntityMap: map[string]*Entity{
				"e1": {Attrs: map[string]*Attribute{
					"a": {Ref: Ref{EntityName: "e01"}}}},
				"e2": {Attrs: map[string]*Attribute{
					"a": {Ref: Ref{EntityName: "e01"}},
					"b": {Ref: Ref{EntityName: "e1"}}}},
			}}},
			want: []string{"e01", "e1", "e2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RefEntities(tt.args.group); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RefEntities() = %v, want %v", got, tt.want)
			}
		})
	}
}
