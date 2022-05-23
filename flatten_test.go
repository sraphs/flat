package flat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	type args struct {
		src interface{}
		opt Option
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "nil",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "string",
			args: args{
				src: "foo",
			},
			want: map[string]interface{}{"": "foo"},
		},
		{
			name: "bool",
			args: args{
				src: true,
			},
			want: map[string]interface{}{"": true},
		},
		{
			name: "int",
			args: args{
				src: 42,
			},
			want: map[string]interface{}{"": 42},
		},
		{
			name: "map",
			args: args{
				src: map[string]interface{}{
					"foo": "bar",
				},
			},
			want: map[string]interface{}{"foo": "bar"},
		},
		{
			name: "struct",
			args: args{
				src: struct {
					Foo    string
					Nested struct {
						Bar string
					}
				}{
					Foo: "foo",
					Nested: struct {
						Bar string
					}{
						Bar: "bar",
					},
				},
			},
			want: map[string]interface{}{
				"Foo":        "foo",
				"Nested.Bar": "bar",
			},
		},
		{
			name: "struct with lower case",
			args: args{
				src: struct {
					Foo    string
					Nested struct {
						Bar string
					}
				}{
					Foo: "foo",
					Nested: struct {
						Bar string
					}{
						Bar: "bar",
					},
				},
				opt: Option{
					Case: CaseLower,
				},
			},
			want: map[string]interface{}{
				"foo":        "foo",
				"nested.bar": "bar",
			},
		},
		{
			name: "complex",
			args: args{
				src: map[string]interface{}{
					"nil":    nil,
					"single": "value",
					"FooBar": "baz",
					"user": map[string]interface{}{
						"string": "string",
						"int":    42,
						"bool":   true,
					},
					"nested": map[string]interface{}{
						"foo": "bar",
						"baz": map[string]interface{}{
							"qux": "quz",
						},
					},
					"struct": struct {
						Foo string
						Bar int
					}{
						Foo: "foo",
						Bar: 42,
					},
				},
				opt: Option{
					Case:      CaseLower,
					Separator: ".",
				},
			},
			want: map[string]interface{}{
				"nil":            nil,
				"single":         "value",
				"foobar":         "baz",
				"user.string":    "string",
				"user.int":       42,
				"user.bool":      true,
				"nested.foo":     "bar",
				"nested.baz.qux": "quz",
				"struct.foo":     "foo",
				"struct.bar":     42,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.opt.Flatten(tt.args.src)
			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkFlatten(b *testing.B) {
	complex := map[string]interface{}{
		"nil":    nil,
		"single": "value",
		"FooBar": "baz",
		"user": map[string]interface{}{
			"string": "string",
			"int":    42,
			"bool":   true,
		},
		"nested": map[string]interface{}{
			"foo": "bar",
			"baz": map[string]interface{}{
				"qux": "quz",
			},
		},
		"struct": struct {
			Foo string
			Bar int
		}{
			Foo: "foo",
			Bar: 42,
		},
	}

	for i := 0; i < b.N; i++ {
		Option{}.Flatten(complex)
	}
}
