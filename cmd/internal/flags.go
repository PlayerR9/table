package internal

import (
	"errors"
	"flag"
	"fmt"

	ggen "github.com/PlayerR9/go-generator/generator"
)

var (
	OutputLocFlag *ggen.OutputLocVal

	TypeListFlag *ggen.TypeListVal

	GenericsFlag *ggen.GenericsSignVal

	TypeNameFlag *string
)

func init() {
	TypeNameFlag = flag.String("name", "", "The name of the generated type. It must be set.")

	OutputLocFlag = ggen.NewOutputFlag("<type>_table.go", false)

	TypeListFlag = ggen.NewTypeListFlag("type", true, 1, "The type of each table's cell.")
	GenericsFlag = ggen.NewGenericsSignFlag("g", false, 1)
}

func Parse() (string, error) {
	ggen.ParseFlags()

	if TypeNameFlag == nil {
		return "", errors.New("flag TypeNameFlag must be set")
	}

	type_name := *TypeNameFlag

	err := ggen.IsValidVariableName(type_name, nil, ggen.Exported)
	if err != nil {
		return "", fmt.Errorf("could not validate type name: %w", err)
	}

	return type_name, nil
}
