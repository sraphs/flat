package flat

import (
	"reflect"
	"testing"
)

func TestUnflatten(t *testing.T) {
	type args struct {
		src map[string]interface{}
		opt Option
	}
	tests := []struct {
		name    string
		args    args
		wantDst map[string]interface{}
	}{
		{
			name: "nil",
			args: args{
				src: nil,
			},
			wantDst: nil,
		},
		{
			name: "empty map key",
			args: args{
				src: map[string]interface{}{
					"": "foo",
				},
			},
			wantDst: map[string]interface{}{
				"": "foo",
			},
		},
		{
			name: "single key",
			args: args{
				src: map[string]interface{}{
					"nil":    nil,
					"string": "foo",
					"bool":   true,
					"int":    42,
				},
			},
			wantDst: map[string]interface{}{
				"nil":    nil,
				"string": "foo",
				"bool":   true,
				"int":    42,
			},
		},
		{
			name: "multiple keys",
			args: args{
				src: map[string]interface{}{
					"nil.nil":       nil,
					"string.string": "bar",
					"bool.bool":     true,
					"int.int":       42,
				},
			},
			wantDst: map[string]interface{}{
				"nil": map[string]interface{}{
					"nil": nil,
				},
				"string": map[string]interface{}{
					"string": "bar",
				},
				"bool": map[string]interface{}{
					"bool": true,
				},
				"int": map[string]interface{}{
					"int": 42,
				},
			},
		},
		{
			name: "multiple keys with option",
			args: args{
				src: map[string]interface{}{
					"nil_nilFoo":       nil,
					"string_stringFoo": "bar",
					"bool_boolFoo":     true,
					"int_intFoo":       42,
				},
				opt: Option{
					Case:      CaseSnake,
					Separator: "_",
				},
			},
			wantDst: map[string]interface{}{
				"nil": map[string]interface{}{
					"nil_foo": nil,
				},
				"string": map[string]interface{}{
					"string_foo": "bar",
				},
				"bool": map[string]interface{}{
					"bool_foo": true,
				},
				"int": map[string]interface{}{
					"int_foo": 42,
				},
			},
		},
		{
			name: "nested struct",
			args: args{
				src: map[string]interface{}{
					"nested.foo": struct {
						Bar string
					}{
						Bar: "baz",
					},
				},
			},
			wantDst: map[string]interface{}{
				"nested": map[string]interface{}{
					"foo": struct {
						Bar string
					}{
						Bar: "baz",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDst := tt.args.opt.Unflatten(tt.args.src)

			if !reflect.DeepEqual(gotDst, tt.wantDst) {
				t.Errorf("Unflatten() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}
