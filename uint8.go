// Code generated by go:generate. DO NOT EDIT.
package table

import "github.com/PlayerR9/lib_units/common"

// Uint8Table represents a table of cells that can be drawn to the screen.
type Uint8Table struct {
	table         [][]uint8
	width, height int
}

// Iterator implements the common.Iterable interface.
//
// The returned iterator is a pull-model iterator that scans the table row by row 
// as it was an array of elements of type uint8.
//
// Example:
//
//	[ a b c ]
//	[ d e f ]
//
//	Iterator() -> [ a ] -> [ b ] -> [ c ] -> [ d ] -> [ e ] -> [ f ]
func (t *Uint8Table) Iterator() common.Iterater[uint8] {
	iter := common.NewDynamicIterator(
		common.NewSimpleIterator(t.table),
		func(row []uint8) common.Iterater[uint8] {
			return common.NewSimpleIterator(row)
		},
	)

	return iter
}

// Cleanup implements the Utility.Cleaner interface.
//
// It sets all cells in the table to the zero value of type uint8.
func (t *Uint8Table) Cleanup() {
	for i := 0; i < t.height; i++ {
		t.table[i] = make([]uint8, t.width)
	}
}

// NewUint8Table creates a new table of type uint8 with the given width and height.
// Negative parameters are treated as absolute values.
//
// Parameters:
//   - width: The width of the table.
//   - height: The height of the table.
//
// Returns:
//   - *Uint8Table: The new table. Never nil.
func NewUint8Table(width, height int) *Uint8Table {
	if width < 0 {
		width = -width
	}

	if height < 0 {
		height = -height
	}

	table := make([][]uint8, height)
	for i := 0; i < height; i++ {
		table[i] = make([]uint8, width)
	}

	return &Uint8Table{
		table:  table,
		width:  width,
		height: height,
	}
}

// GetWidth returns the width of the table.
//
// Returns:
//   - int: The width of the table. Never negative.
func (t *Uint8Table) GetWidth() int {
	return t.width
}

// GetHeight returns the height of the table.
//
// Returns:
//   - int: The height of the table. Never negative.
func (t *Uint8Table) GetHeight() int {
	return t.height
}

// WriteAt writes a cell to the table at the given coordinates. However, out-of-bounds
// coordinates do nothing.
//
// Parameters:
//   - x: The x-coordinate of the cell.
//   - y: The y-coordinate of the cell.
//   - cell: The cell to write to the table.
func (t *Uint8Table) WriteAt(x, y int, cell uint8) {
	if x < 0 || x >= t.width || y < 0 || y >= t.height {
		return
	}

	t.table[y][x] = cell
}

// GetAt returns the cell at the given coordinates in the table. However, out-of-bounds
// coordinates return the zero value of type uint8.
//
// Parameters:
//   - x: The x-coordinate of the cell.
//   - y: The y-coordinate of the cell.
//
// Returns:
//   - uint8: The cell at the given coordinates.
func (t *Uint8Table) GetAt(x, y int) uint8 {
	if x < 0 || x >= t.width || y < 0 || y >= t.height {
		return 0
	} else {
		return t.table[y][x]
	}
}

// WriteVerticalSequence is a function that writes the specified values to the table
// starting from the specified coordinates (top = 0, 0) and continuing down the
// table in the vertical direction until either the sequence is exhausted or
// the end of the table is reached; at which point any remaining values in the
// sequence are ignored.
//
// Due to implementation details, any value that would be written outside are ignored.
// As such, if x is out-of-bounds, the function does nothing and, if y is out-of-bounds,
// only out-of-bounds values are not written.
//
// Parameters:
//   - x: The x-coordinate of the starting cell. (Never changes)
//   - y: The y-coordinate of the starting cell.
//   - sequence: The sequence of cells to write to the table.
//
// At the end of the function, the y coordinate points to the cell right below the
// last cell in the sequence that was written.
//
// Example:
//
//	// [ a b c ]
//	// [ d e f ]
//	//
//	// seq := [ g h i ], x = 0, y = -1
//
//	WriteVerticalSequence(x, y, seq)
//
//	// [ h b c ]
//	// [ i e f ]
//	//
//	// x = 0, y = 2
//
// As you can see, the 'g' value was ignored as it would be out-of-bounds.
// Finally, if either x or y is nil, the function does nothing.
func (t *Uint8Table) WriteVerticalSequence(x, y *int, sequence []uint8) {
	if x == nil || y == nil {
		return
	}

	actualX, actualY := *x, *y

	if len(sequence) == 0 || actualX < 0 || actualX >= t.width || actualY >= t.height {
		return
	}

	if actualY < 0 {
		sequence = sequence[-actualY:]

		*y = 0
	} else if actualY+len(sequence) > t.height {
		sequence = sequence[:t.height-actualY]
	}

	for i, cell := range sequence {
		t.table[actualY+i][actualX] = cell
	}

	*y += len(sequence)
}

// WriteHorizontalSequence is the equivalent of WriteVerticalSequence but for horizontal
// sequences.
//
// See WriteVerticalSequence for more information.
//
// Parameters:
//   - x: The x-coordinate of the starting cell.
//   - y: The y-coordinate of the starting cell.
//   - sequence: The sequence of cells to write to the table.
func (t *Uint8Table) WriteHorizontalSequence(x, y *int, sequence []uint8) {
	if x == nil || y == nil {
		return
	}

	actualX, actualY := *x, *y

	if len(sequence) == 0 || actualY < 0 || actualY >= t.height || actualX >= t.width {
		return
	}

	if actualX < 0 {
		sequence = sequence[-actualX:]

		*x = 0
	} else if actualX+len(sequence) > t.width {
		sequence = sequence[:t.width-actualX]
	}

	copy(t.table[actualY][actualX:], sequence)

	*x = actualX + len(sequence)
}

// GetFullTable returns the full table as a 2D slice of elements of type uint8.
//
// Returns:
//   - [][]uint8: The full table.
func (t *Uint8Table) GetFullTable() [][]uint8 {
	return t.table
}

// IsXInBounds checks if the given x-coordinate is within the bounds of the table.
//
// Parameters:
//   - x: The x-coordinate to check.
//
// Returns:
//   - error: An error of type *common.ErrOutOfBounds if the x-coordinate is out of bounds.
func (t *Uint8Table) IsXInBounds(x int) error {
	if x < 0 || x >= t.width {
		return common.NewErrOutOfBounds(x, 0, t.width)
	} else {
		return nil
	}
}

// IsYInBounds checks if the given y-coordinate is within the bounds of the table.
//
// Parameters:
//   - y: The y-coordinate to check.
//
// Returns:
//   - error: An error of type *common.ErrOutOfBounds if the y-coordinate is out of bounds.
func (t *Uint8Table) IsYInBounds(y int) error {
	if y < 0 || y >= t.height {
		return common.NewErrOutOfBounds(y, 0, t.height)
	} else {
		return nil
	}
}

// WriteTableAt is a convenience function that copies the values from the given
// table to the table starting at the given coordinates in a more efficient way 
// than using any other methods.
//
// While it acts in the same way as both WriteVerticalSequence and WriteHorizontalSequence
// combined, it is more efficient than calling those two functions separately.
//
// See WriteVerticalSequence for more information.
//
// Parameters:
//   - table: The table to write to the table.
//   - x: The x-coordinate to write the table at.
//   - y: The y-coordinate to write the table at.
//
// If the table is nil, x or y are nil, nothing happens.
func (t *Uint8Table) WriteTableAt(table *Uint8Table, x, y *int) {
	if table == nil || x == nil || y == nil {
		return
	}

	offsetX, offsetY := 0, 0
	X, Y := *x, *y

	for offsetY < table.height && Y+offsetY < t.height {
		offsetX = 0

		for offsetX < table.width && X+offsetX < t.width {
			t.table[Y+offsetY][X+offsetX] = table.table[offsetY][offsetX]
			offsetX++
		}

		offsetY++
	}

	*x += offsetX
	*y += offsetY
}
