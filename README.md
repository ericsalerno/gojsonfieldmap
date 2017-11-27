# gojsonfieldmap

This library builds a JSON Path field mapping string from an object structure. It relies on the reflect library to perform its duties. If this is a problem, I'd suggest skipping this library.

The purpose is to take an arbitrary data structure and turn it into a json representation that toggles fields based on having a 1 value. It will ignore "untagged" json fields and it will create subchild mappings of substructures using recursion (I know, sorry).

## Install

    go get -u github.com/ericsalerno/gojsonfieldmap

## Example

Create an object struct you'd like to use:

    type simpleObject struct {
	    Name  string `json:"Name"`
	    Value string `json:"Value"`
    }

Run the one function this package provides on it:

    object := simpleObject{}
    output := gojsonfieldmap.GetJSONObjectFieldMap(object)
    
    if output != `{"Name":1,"Value":1}` {
		fmt.Println("Something has gone wrong because this should work.")
	}

Output will become the "field mapping" JSON equivilent of the input structure. See the main_test.go file for more complex examples.