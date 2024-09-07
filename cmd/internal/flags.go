package internal

import (
	"errors"
	"flag"
	"fmt"

	gcgen "github.com/PlayerR9/go-commons/generator"
)

var (
	OutputLocFlag *gcgen.OutputLocVal

	TypeListFlag *gcgen.TypeListVal

	GenericsFlag *gcgen.GenericsSignVal

	TypeNameFlag *string
)

func init() {
	TypeNameFlag = flag.String("name", "", "The name of the generated type. It must be set.")

	OutputLocFlag = gcgen.NewOutputFlag("<type>_table.go", false)

	TypeListFlag = gcgen.NewTypeListFlag("type", true, 1, "The type of each table's cell.")
	GenericsFlag = gcgen.NewGenericsSignFlag("g", false, 1)
}

func Parse() (string, error) {
	gcgen.ParseFlags()

	if TypeNameFlag == nil {
		return "", errors.New("flag TypeNameFlag must be set")
	}

	type_name := *TypeNameFlag

	err := gcgen.IsValidVariableName(type_name, nil, gcgen.Exported)
	if err != nil {
		return "", fmt.Errorf("could not validate type name: %w", err)
	}

	return type_name, nil
}
