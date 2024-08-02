// This command generates a table of the given type implemented as a boundless table (i.e., out-of-bounds errors are not thrown).
//
// To use it, run the following command:
//
// //go:generate go run table/cmd -name=<type_name> -type=<type> [ -g=<generics>] [ -o=<output_file> ]
//
// **Flag: Name**
//
// The "name" flag is used to specify the name of the table. As such, it must be set and,
// not only does it have to be a valid Go identifier, but it also must start with an upper case letter.
//
// **Flag: Type**
//
// The "fields" flag is used to specify type of the table's cells. This flag must be set.
//
// For instance, running the following command:
//
//	//go:generate treenode -name=Table -type=int
//
// will generate a table with the following structure:
//
//	type Table struct {
//		table [][]int
//	}
//
// It is important to note that spaces are not allowed in any of the flags.
//
// Also, it is possible to specify generics by following the value with the generics between square brackets;
// like so: "MyType[T,C]"
//
// **Flag: Generics**
//
// This optional flag is used to specify the type(s) of the generics. However, this only applies if at least one
// generic type is specified in the type flag. If none, then this flag is ignored.
//
// As an edge case, if this flag is not specified but the type flag contains generics, then
// all generics are set to the default value of "any".
//
// Its argument is specified as a list of key-value pairs where each pair is separated
// by a comma (",") and a slash ("/") is used to separate the key and the value. The key indicates the name of
// the generic and the value indicates the type of the generic.
//
// For instance, running the following command:
//
//	//go:generate table -name=Table -type=MyType[T] -g=T/any
//
// will generate a table with the following fields:
//
//	type Table[T any] struct {
//		table [][]T
//	}
//
// **Flag: Output File**
//
// This optional flag is used to specify the output file. If not specified, the output will be written to
// standard output, that is, the file "<type_name>_table.go" in the root of the current directory.
package main

import (
	"os"
	"path/filepath"

	ggen "github.com/PlayerR9/lib_units/generator"
	pkg "github.com/PlayerR9/table/cmd/internal"
)

func main() {
	type_name, err := pkg.Parse()
	if err != nil {
		pkg.Logger.Fatal(err.Error())
	}

	g := &pkg.GenData{
		TypeName:     type_name,
		GenericsSign: ggen.GenericsSigFlag.String(),
	}

	res, err := pkg.Generator.Generate(type_name, "_table.go", g)
	if err != nil {
		pkg.Logger.Fatalf("Could not generate code: %s", err.Error())
	}

	dir := filepath.Dir(res.DestLoc)

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		pkg.Logger.Fatalf("Could not create directory: %s", err.Error())
	}

	err = os.WriteFile(res.DestLoc, res.Data, 0644)
	if err != nil {
		pkg.Logger.Fatalf("Could not write to file: %s", err.Error())
	}

	pkg.Logger.Printf("Successfully Generated %s", res.DestLoc)
}
