package validate

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Violation struct {
	Field      string
	Constraint string
}

// a map with validation functions as values, each associated with the validation name as the key
var assertFns = map[string]func(assertions map[string]string, v reflect.Value, n string, vs *[]Violation, path string) *[]Violation{
	"required":  assertRequired,
	"min":       assertMin,
	"max":       assertMax,
	"pattern":   assertPattern,
	"maxlength": assertMaxLength,
	"minlength": assertMinLength,
}

// This exported function is interface for validating a struct. It returns a slice of @Violation structs for each
// violation encountered.
func Validate(ifc interface{}) []Violation {
	var violations = make([]Violation, 0)
	assertAll(ifc, &violations, "")
	return violations
}

// Validate iterates through each Field in the provided interface.
func assertAll(ifc interface{}, violations *[]Violation, path string) {
	v := reflect.ValueOf(ifc)
	t := reflect.TypeOf(ifc)

	// set the path to the current struct
	path = asPath(path, t)

	for i := 0; i < t.NumField(); i++ {
		// validate the struct's fields
		violations = validate(v, i, violations, path)

		// walk the rest of the object graph
		switch v.Field(i).Kind() {
		case reflect.Struct:
			fieldValue := v.Field(i).Interface()

			// set parent path
			assertAll(fieldValue, violations, path)
		case reflect.Slice, reflect.Array:
			slice := v.Field(i)

			for idx := 0; idx < slice.Len(); idx++ {
				assertAll(slice.Index(idx).Elem().Interface(), violations, path)
			}
		default:
			// todo
		}
	}
}

// Executes validation checks for each of the checks required for a Field.
func validate(v reflect.Value, fieldIndex int, violations *[]Violation, path string) *[]Violation {
	tag := v.Type().Field(fieldIndex).Tag
	val := v.Field(fieldIndex)
	name := v.Type().Field(fieldIndex).Name

	// get a map of assertions to validate
	assertions := asChecks(tag)

	for assertion, _ := range assertions {
		if fnValidation, ok := assertFns[assertion]; ok {
			violations = fnValidation(assertions, val, name, violations, path)
		}
	}

	return violations
}

// Returns a slice of validations to check.
func asChecks(tag reflect.StructTag) map[string]string {
	checks := make(map[string]string)

	if t, ok := tag.Lookup("assert"); ok {
		assertTags := strings.Split(t, ":")
		tagPairs := strings.Split(assertTags[0], ",")

		for _, pair := range tagPairs {
			key, value, err := asKeyValue(pair, "=")

			if err != nil {
				fmt.Println(err)
			}

			checks[key] = value
		}
	}

	return checks
}

// Checks that the Field value exists.
func assertRequired(assertions map[string]string, val reflect.Value, name string, violations *[]Violation, path string) *[]Violation {
	if required, ok := assertions["required"]; ok {
		if required == "true" && IsNilOrEmpty(val) {
			violation := Violation{Field: asQualifiedPath(path, name), Constraint: "required"}
			*violations = append(*violations, violation)
		}
	}

	return violations
}

// Checks that the value is not less than the minimum value.
func assertMin(assertions map[string]string, val reflect.Value, name string, violations *[]Violation, path string) *[]Violation {
	if _, ok := assertions["min"]; ok {
		switch val.Type().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

			min, err := strconv.ParseInt(assertions["min"], 10, 64)

			if err != nil {
				log.Printf("%s:%+v", "unable to parse min tag value", err)
			}

			fieldValue := val.Int()

			if !IsNilOrEmpty(val) && fieldValue < min {
				violation := Violation{Field: asQualifiedPath(path, name), Constraint: "min"}
				*violations = append(*violations, violation)
			}
		case reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:

			min, err := strconv.ParseFloat(assertions["min"], 64)

			if err != nil {
				log.Printf("%s:%+v", "unable to parse min tag value", err)
			}

			fieldValue := val.Float()

			if !IsNilOrEmpty(val) && fieldValue < min {
				violation := Violation{Field: asQualifiedPath(path, name), Constraint: "min"}
				*violations = append(*violations, violation)
			}
		default:
			log.Printf("invalid field type used with min validation")
		}
	}

	return violations
}

// Checks that the value is not greater than the maximum value.
func assertMax(assertions map[string]string, val reflect.Value, name string, violations *[]Violation, path string) *[]Violation {
	if _, ok := assertions["max"]; ok {
		switch val.Type().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

			max, err := strconv.ParseInt(assertions["max"], 10, 64)

			if err != nil {
				log.Printf("%s:%+v", "unable to parse max tag value", err)
			}

			fieldValue := val.Int()

			if !IsNilOrEmpty(val) && fieldValue > max {
				violation := Violation{Field: asQualifiedPath(path, name), Constraint: "max"}
				*violations = append(*violations, violation)
			}
		case reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:

			max, err := strconv.ParseFloat(assertions["max"], 64)

			if err != nil {
				log.Printf("%s:%+v", "unable to parse max tag value", err)
			}

			fieldValue := val.Float()

			if !IsNilOrEmpty(val) && fieldValue > max {
				violation := Violation{Field: asQualifiedPath(path, name), Constraint: "max"}
				*violations = append(*violations, violation)
			}
		default:
			log.Printf("invalid field type used with max validation")
		}
	}

	return violations
}

func assertPattern(assertions map[string]string, val reflect.Value, name string, violations *[]Violation, path string) *[]Violation {
	if pattern, ok := assertions["pattern"]; ok {
		if matched, _ := regexp.MatchString(pattern, val.Interface().(string)); !matched {
			violation := Violation{Field: asQualifiedPath(path, name), Constraint: "pattern"}
			*violations = append(*violations, violation)
		}
	}

	return violations
}

func assertMaxLength(assertions map[string]string, val reflect.Value, name string, violations *[]Violation, path string) *[]Violation {
	if m, ok := assertions["maxlength"]; ok {
		maxLength, err := strconv.Atoi(m)

		if err != nil {
			log.Printf("%s:%+v", "unable to parse maxlength tag value", err)
		}

		if !IsNilOrEmpty(val) && len(val.Interface().(string)) > maxLength {
			violation := Violation{Field: asQualifiedPath(path, name), Constraint: "maxlength"}
			*violations = append(*violations, violation)
		}
	}

	return violations
}

func assertMinLength(assertions map[string]string, val reflect.Value, name string, violations *[]Violation, path string) *[]Violation {
	if m, ok := assertions["minlength"]; ok {
		minLength, err := strconv.Atoi(m)

		if err != nil {
			log.Printf("%s:%+v", "unable to parse minlength tag value", err)
		}

		if !IsNilOrEmpty(val) && len(val.Interface().(string)) < minLength {
			violation := Violation{Field: asQualifiedPath(path, name), Constraint: "minlength"}
			*violations = append(*violations, violation)
		}
	}

	return violations
}

func asQualifiedPath(path string, name string) string {
	var qualifiedPath string
	if path != "" {
		qualifiedPath = path + "." + name
	} else {
		qualifiedPath = name
	}
	return qualifiedPath
}

// Returns true if the value is nil or empty.
func IsNilOrEmpty(v reflect.Value) bool {
	switch v.Type().Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32,
		reflect.Float64, reflect.Complex64, reflect.Complex128:
		return &v == nil
	case reflect.Array:
		if &v != nil {
			for i := 0; i < v.Len(); i++ {
				if !v.Index(i).IsZero() {
					return false
				}
			}
			return true
		}
		return false
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return v.IsNil()
	case reflect.String:
		return &v == nil || v.Len() == 0
	case reflect.Struct:
		if &v != nil {
			for i := 0; i < v.NumField(); i++ {
				if !IsNilOrEmpty(v.Field(i)) {
					return false
				}
			}
			return true
		}
		return false
	default:
		panic(&reflect.ValueError{"unknown value kind", v.Kind()})
	}
}

// Returns the constructed struct path from the type's name.
func asPath(path string, t reflect.Type) string {
	if path != "" {
		path += "." + t.Name()
	} else {
		path += t.Name()
	}
	return path
}

func asKeyValue(pair, sep string) (string, string, error) {
	if pair == "" {
		return "", "", errors.New("pair is an empty string")
	}

	if sep == "" {
		return "", "", errors.New("sep is an empty string")
	}

	if !strings.Contains(pair, sep) {
		return "", "", errors.New(fmt.Sprintf("%s doesn't contain the '%s' separator", pair, sep))
	}

	keyValue := strings.Split(pair, "=")

	if len(keyValue) != 2 {
		return "", "", errors.New("pair doesn't contain both a key and value")
	}

	return keyValue[0], keyValue[1], nil
}
