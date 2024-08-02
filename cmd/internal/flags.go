package internal

import (
	"errors"
	"flag"
	"fmt"

	ggen "github.com/PlayerR9/lib_units/generator"
)

var (
	TypeNameFlag *string
)

func init() {
	TypeNameFlag = flag.String("name", "", "The name of the generated type. It must be set.")

	ggen.SetOutputFlag("<type>_table.go", false)

	ggen.SetTypeListFlag("type", true, 1, "The type of each table's cell.")
	ggen.SetGenericsSignFlag("g", false, 1)
}

func Parse() (string, error) {
	err := ggen.ParseFlags()
	if err != nil {
		return "", fmt.Errorf("could not parse flags: %w", err)
	}

	if TypeNameFlag == nil {
		return "", errors.New("flag TypeNameFlag must be set")
	}

	type_name := *TypeNameFlag

	err = ggen.IsValidName(type_name, nil, ggen.Exported)
	if err != nil {
		return "", fmt.Errorf("could not validate type name: %w", err)
	}

	return type_name, nil
}
