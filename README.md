# Go Assert

Go Assert is a library for validating struct fields.

## Installation

Add the following import.

```go
import "github.com/tendryll/goassert/assert"
```

If using dep, execute:

```go
dep ensure -add github.com/tendryll/goassert/assert
```

## Usage

Go Assert executes assertions based on assertion types added to the assert tag. The assertion types supported are:

* required: Used to verify that the value is not nil.
* min: Used verify that the field value is equal to or greater than the min value specified.
* max: Used to verify that the field value is equal to or less than the max value specified.
* pattern: Used to verify that the field value, a string, matches the regular expression specified.
* maxlength: Used to verify that the length of the field of type string is no longer than the value specified.
* minlength: Used to verify that the length of the field of type string is no shorter than the value specified.

Let's assume that we have a struct named `Latitude` with two fields, `Degrees` of type float64 and `Direction` of
 type string. To add an assertion check for the `Degrees` field that ensures that field have been set and
  the value is within the range of 0.0 and 9.0 and for the `Direction` field to match a certain pattern, then we
   could do the following:

```go
type Latitude struct {
    Degrees   float64 `json:"degrees" assert:"required=true,min=0.0,max=90.0"`
    Direction string  `json:"direction" assert:"required=true,pattern=^(N|S)$"`
}
```

The Go Assert library is used by passing the desired struct to the `validate.Validate` function. 
Returned is a slice of type []Violation. Each violation specifies the constraint used and the field that
  failed the validation check. 
 
As an example,
 
 ```go
latitude := &models.Latitude{
    Degrees:   135.1098212,
    Direction: "N",
}

validate.Validate(*latitude)    
```

will return

```go
[{Field:Latitude.Degrees Constraint:max}]
```

## Contributing
Please open an issue to discuss changes you wish to be made. Pull requests are welcome. Please make sure to add or 
update tests as needed.

## License
[BSD 3-Clause License](https://opensource.org/licenses/BSD-3-Clause)


