package validate

import (
	"reflect"
	"testing"
)

func Test_assertMax(t *testing.T) {
	type args struct {
		assertion  map[string]string
		val        reflect.Value
		name       string
		violations *[]Violation
		path       string
	}
	tests := []struct {
		name     string
		args     args
		expected *[]Violation
	}{
		{
			name: "scenario1",
			args: args{
				name:       "Width",
				val:        reflect.ValueOf(struct{ Width int64 }{Width: 3}).Field(0),
				assertion:  map[string]string{"max": "3"},
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{},
		},
		{
			name: "scenario2",
			args: args{
				name:       "Width",
				val:        reflect.ValueOf(struct{ Width int64 }{Width: 2}).Field(0),
				assertion:  map[string]string{"max": "3"},
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{},
		},
		{
			name: "scenario3",
			args: args{
				name:       "Width",
				val:        reflect.ValueOf(struct{ Width int64 }{Width: 4}).Field(0),
				assertion:  map[string]string{"max": "4"},
				violations: &[]Violation{{Field: "Width", Constraint: "max"}},
				path:       "",
			},
			expected: &[]Violation{{Field: "Width", Constraint: "max"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := assertMax(tt.args.assertion, tt.args.val, tt.args.name, tt.args.violations, tt.args.path); !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("assertMax() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_assertMaxLength(t *testing.T) {
	type args struct {
		assertion  map[string]string
		val        reflect.Value
		name       string
		violations *[]Violation
		path       string
	}
	tests := []struct {
		name     string
		args     args
		expected *[]Violation
	}{
		{
			name: "scenario1",
			args: args{
				name:       "Name",
				assertion:  map[string]string{"maxlength": "3"},
				val:        reflect.ValueOf(struct{ Name string }{Name: "one"}).Field(0),
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{},
		},
		{
			name: "scenario2",
			args: args{
				name:       "Name",
				assertion:  map[string]string{"maxlength": "4"},
				val:        reflect.ValueOf(struct{ Name string }{Name: "one"}).Field(0),
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{},
		},
		{
			name: "scenario3",
			args: args{
				name:       "Name",
				assertion:  map[string]string{"maxlength": "2"},
				val:        reflect.ValueOf(struct{ Name string }{Name: "one"}).Field(0),
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{{Field: "Name", Constraint: "maxlength"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := assertMaxLength(tt.args.assertion, tt.args.val, tt.args.name, tt.args.violations, tt.args.path); !reflect.DeepEqual(*result, *tt.expected) {
				t.Errorf("assertMaxLength() = %+v, expected %+v", *result, *tt.expected)
			}
		})
	}
}

func Test_assertMin(t *testing.T) {
	type args struct {
		assertion  map[string]string
		val        reflect.Value
		name       string
		violations *[]Violation
		path       string
	}
	tests := []struct {
		name     string
		args     args
		expected *[]Violation
	}{
		{
			name: "scenario1",
			args: args{
				name:       "Width",
				val:        reflect.ValueOf(struct{ Width float64 }{Width: 1}).Field(0),
				assertion:  map[string]string{"min": "1"},
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{},
		},
		{
			name: "scenario2",
			args: args{
				name:       "Width",
				val:        reflect.ValueOf(struct{ Width float64 }{Width: 1}).Field(0),
				assertion:  map[string]string{"min": "0"},
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{},
		},
		{
			name: "scenario3",
			args: args{
				name:       "Width",
				val:        reflect.ValueOf(struct{ Width float64 }{Width: 1}).Field(0),
				assertion:  map[string]string{"min": "2"},
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{{Field: "Width", Constraint: "min"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := assertMin(tt.args.assertion, tt.args.val, tt.args.name, tt.args.violations, tt.args.path); !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("assertMin() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_assertMinLength(t *testing.T) {
	type args struct {
		assertion  map[string]string
		val        reflect.Value
		name       string
		violations *[]Violation
		path       string
	}
	tests := []struct {
		name     string
		args     args
		expected *[]Violation
	}{
		{
			name: "scenario1",
			args: args{
				name:       "Name",
				assertion:  map[string]string{"minlength": "1"},
				val:        reflect.ValueOf(struct{ Name string }{Name: "one"}).Field(0),
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{},
		},
		{
			name: "scenario2",
			args: args{
				name:       "Name",
				assertion:  map[string]string{"minlength": "3"},
				val:        reflect.ValueOf(struct{ Name string }{Name: "one"}).Field(0),
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{},
		},
		{
			name: "scenario3",
			args: args{
				name:       "Name",
				assertion:  map[string]string{"minlength": "4"},
				val:        reflect.ValueOf(struct{ Name string }{Name: "one"}).Field(0),
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{{Field: "Name", Constraint: "minlength"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := assertMinLength(tt.args.assertion, tt.args.val, tt.args.name, tt.args.violations, tt.args.path); !reflect.DeepEqual(*result, *tt.expected) {
				t.Errorf("assertMinLength() = %+v, expected %+v", *result, *tt.expected)
			}
		})
	}
}

func Test_assertPattern(t *testing.T) {
	type args struct {
		assertion  map[string]string
		val        reflect.Value
		name       string
		violations *[]Violation
		path       string
	}
	tests := []struct {
		name     string
		args     args
		expected *[]Violation
	}{
		{
			name: "scenario1",
			args: args{
				assertion:  map[string]string{"pattern": "^b.*$"},
				val:        reflect.ValueOf(struct{ Animal string }{Animal: "bird"}).Field(0),
				name:       "Animal",
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{},
		},
		{
			name: "scenario2",
			args: args{
				assertion:  map[string]string{"pattern": "^b.*$"},
				val:        reflect.ValueOf(struct{ Animal string }{Animal: "cat"}).Field(0),
				name:       "Animal",
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{{Field: "Animal", Constraint: "pattern"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := assertPattern(tt.args.assertion, tt.args.val, tt.args.name, tt.args.violations, tt.args.path); !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("assertPattern() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_assertRequired(t *testing.T) {
	type args struct {
		assertion  map[string]string
		val        reflect.Value
		name       string
		violations *[]Violation
		path       string
	}
	tests := []struct {
		name     string
		args     args
		expected *[]Violation
	}{
		{
			name: "scenario1",
			args: args{
				name:       "FirstName",
				assertion:  map[string]string{"required": "true"},
				val:        reflect.ValueOf(struct{ FirstName string }{FirstName: "Brad"}).Field(0),
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{},
		},
		{
			name: "scenario2",
			args: args{
				name:       "FirstName",
				val:        reflect.ValueOf(struct{ FirstName string }{FirstName: ""}).Field(0),
				assertion:  map[string]string{"required": "true"},
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{{Field: "FirstName", Constraint: "required"}},
		},
		{
			name: "scenario3",
			args: args{
				name:       "FirstName",
				val:        reflect.ValueOf(struct{ FirstName string }{}).Field(0),
				assertion:  map[string]string{"required": "true"},
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{{Field: "FirstName", Constraint: "required"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := assertRequired(tt.args.assertion, tt.args.val, tt.args.name, tt.args.violations, tt.args.path); !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("assertRequired() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_asChecks(t *testing.T) {
	type args struct {
		tag reflect.StructTag
	}
	tests := []struct {
		name     string
		args     args
		expected map[string]string
	}{
		{
			name: "scenario1",
			args: args{
				tag: reflect.TypeOf(struct {
					Count int `assert:"required=true,min=1,max=10"`
				}{}).Field(0).Tag,
			},
			expected: map[string]string{"required": "true", "min": "1", "max": "10"},
		},
		{
			name: "scenario1",
			args: args{
				tag: reflect.TypeOf(struct{ Count int }{}).Field(0).Tag,
			},
			expected: map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := asChecks(tt.args.tag); !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("asChecks() = %+v, expected %+v", result, tt.expected)
			}
		})
	}
}
