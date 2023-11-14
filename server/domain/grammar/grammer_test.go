package grammar

import (
	"reflect"
	"testing"
)

func TestRemovedValue(t *testing.T) {
	type args struct {
		grammars []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "正常系",
			args: args{
				grammars: []string{"past", "passive"},
			},
			want: []string{
				"present",
				"past participle",
				"present participle",
				"imperative",
				"infinitive",
				"gerund",
				"indicative",
				"subjunctive",
				"active",
				"perfect",
				"progressive"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemovedValue(tt.args.grammars); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemovedValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
