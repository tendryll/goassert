package assert

import (
	"errors"
	"reflect"
	"testing"
)

func TestAssertMax(t *testing.T) {
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
				assertion:  map[string]string{"max": "3"},
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{{Field: "Width", Constraint: "max"}},
		},
		{
			name: "scenario4",
			args: args{
				name:       "Width",
				val:        reflect.ValueOf(struct{ Width float64 }{Width: 3.0}).Field(0),
				assertion:  map[string]string{"max": "3"},
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{},
		},
		{
			name: "scenario5",
			args: args{
				name:       "Width",
				val:        reflect.ValueOf(struct{ Width float64 }{Width: 2.0}).Field(0),
				assertion:  map[string]string{"max": "3"},
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{},
		},
		{
			name: "scenario6",
			args: args{
				name:       "Width",
				val:        reflect.ValueOf(struct{ Width float64 }{Width: 4.0}).Field(0),
				assertion:  map[string]string{"max": "3"},
				violations: &[]Violation{},
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

func TestAssertMaxLength(t *testing.T) {
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

func TestAssertMin(t *testing.T) {
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
		}, {
			name: "scenario4",
			args: args{
				name:       "Width",
				val:        reflect.ValueOf(struct{ Width int64 }{Width: 1}).Field(0),
				assertion:  map[string]string{"min": "1"},
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{},
		},
		{
			name: "scenario5",
			args: args{
				name:       "Width",
				val:        reflect.ValueOf(struct{ Width int64 }{Width: 1}).Field(0),
				assertion:  map[string]string{"min": "0"},
				violations: &[]Violation{},
				path:       "",
			},
			expected: &[]Violation{},
		},
		{
			name: "scenario6",
			args: args{
				name:       "Width",
				val:        reflect.ValueOf(struct{ Width int64 }{Width: 1}).Field(0),
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

func TestAssertMinLength(t *testing.T) {
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

func TestAssertPattern(t *testing.T) {
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

func TestAssertRequired(t *testing.T) {
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

func TestAsAssertions(t *testing.T) {
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
			if result := asAssertions(tt.args.tag); !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("asAssertions() = %+v, expected %+v", result, tt.expected)
			}
		})
	}
}

func TestAsKeyValue(t *testing.T) {
	tests := []struct {
		name          string
		args          string
		expectedKey   string
		expectedValue string
		err           error
	}{
		{
			name:          "scenario1",
			args:          "required=true",
			expectedKey:   "required",
			expectedValue: "true",
			err:           nil,
		},
		{
			name:          "scenario2",
			args:          "required=true,blah",
			expectedKey:   "required",
			expectedValue: "true",
			err:           nil,
		},
		{
			name:          "scenario3",
			args:          "",
			expectedKey:   "",
			expectedValue: "",
			err:           errors.New("pair doesn't contain both a key and value"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actualKey, actualValue, err := asKeyValue(tt.args); actualKey != tt.expectedKey && actualValue != tt.
				expectedValue && err != tt.err {
				t.Errorf("asKeyValue() = %s, %s, %+v, expected %s, %s, %+v", actualKey, actualValue, err,
					tt.expectedKey, tt.expectedValue, tt.err)
			}
		})
	}
}

func TestAsPath(t *testing.T) {
	type args struct {
		path string
		t    reflect.Type
	}

	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "scenario1",
			args: args{
				path: "",
				t:    reflect.TypeOf(Person{}),
			},
			expected: "Person",
		},
		{
			name: "scenario1",
			args: args{
				path: "Person",
				t:    reflect.TypeOf(Address{}),
			},
			expected: "Person.Address",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := asPath(tt.args.path, tt.args.t); actual != tt.expected {
				t.Errorf("asPath() = %s, expected = %s", actual, tt.expected)
			}
		})
	}
}

func TestAssertAll(t *testing.T) {
	type args struct {
		person Person
		path   string
	}

	tests := []struct {
		name     string
		args     args
		expected *[]Violation
	}{
		{
			name: "scenario1",
			args: args{
				person: Person{
					FirstName:  "James",
					MiddleInit: "T",
					LastName:   "Kirk",
					Address: []*Address{
						{
							Address1: "755 Crossover Lane",
							City:     "Memphis",
							State:    "TN",
							Country:  "USA",
							ZipCode:  "38107",
							Location: &Location{
								Latitude: &Latitude{
									Degrees:   35.1098212,
									Direction: "N",
								},
								Longitude: -89.9077976,
							},
						},
					},
				},
				path: "",
			},
			expected: &[]Violation{},
		},
		{
			name: "scenario2",
			args: args{
				person: Person{
					FirstName: "James",
					Address: []*Address{
						{
							Address1: "755 Crossover Lane",
							City:     "Memphis",
							State:    "TN",
							ZipCode:  "38107",
							Location: &Location{
								Latitude: &Latitude{
									Degrees:   135.1098212,
									Direction: "N",
								},
								Longitude: -89.9077976,
							},
						},
					},
				},
				path: "",
			},
			expected: &[]Violation{
				{
					Field:      "Person.LastName",
					Constraint: "required",
				},
				{
					Field:      "Person.Address.Country",
					Constraint: "required",
				},
				{
					Field:      "Person.Address.Location.Latitude.Degrees",
					Constraint: "max",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violations := make([]Violation, 0)
			assertAll(tt.args.person, &violations, tt.args.path)

			if !reflect.DeepEqual(&violations, tt.expected) {
				t.Errorf("assertAll() violations = %+v, expected %+v", violations, tt.expected)
			}
		})
	}
}

type Latitude struct {
	Degrees   float64 `json:"degrees" assert:"required=true,min=0.0,max=90.0"`
	Direction string  `json:"direction" assert:"required=true,pattern=^(N|S)$"`
}

type Location struct {
	Latitude  *Latitude `json:"latitude" assert:"required=true"`
	Longitude float64   `json:"longitude" assert:"required=true,min=-180.0,max=180.0"`
}

type Address struct {
	Address1 string    `json:"address1" assert:"required=true"`
	Address2 string    `json:"address2"`
	Address3 string    `json:"address3"`
	City     string    `json:"city"`
	State    string    `json:"state" assert:"required=true"`
	Country  string    `json:"country" assert:"required=true,maxlength=3"`
	ZipCode  string    `json:"zipcode" assert:"required=true"`
	Location *Location `json:"location"`
}

type Person struct {
	FirstName  string     `json:"firstName" assert:"required=true"`
	MiddleInit string     `json:"middleInitial" assert:"minlength=1,maxlength=1"`
	LastName   string     `json:"lastName" assert:"required=true"`
	Address    []*Address `json:"address" assert:"required=true"`
}
