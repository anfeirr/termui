// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"image"

	. "github.com/gizak/termui"
)

/*Table is like:
┌ Awesome Table ───────────────────────────────────────────────┐
│  Col0          | Col1 | Col2 | Col3  | Col4  | Col5  | Col6  |
│──────────────────────────────────────────────────────────────│
│  Some Item #1  | AAA  | 123  | CCCCC | EEEEE | GGGGG | IIIII |
│──────────────────────────────────────────────────────────────│
│  Some Item #2  | BBB  | 456  | DDDDD | FFFFF | HHHHH | JJJJJ |
└──────────────────────────────────────────────────────────────┘
*/
type Table struct {
	Block
	Rows         [][]string
	ColumnWidths []int
	TextStyle    Style
	RowSeparator bool
	TextAlign    Alignment
}

func NewTable() *Table {
	return &Table{
		Block:        *NewBlock(),
		TextStyle:    Theme.Table.Text,
		RowSeparator: true,
	}
}

func (self *Table) Draw(buf *Buffer) {
	self.Block.Draw(buf)

	columnWidths := self.ColumnWidths
	if len(columnWidths) == 0 {
		columnCount := len(self.Rows[0])
		colWidth := self.Inner.Dx() / columnCount
		for i := 0; i < columnCount; i++ {
			columnWidths = append(columnWidths, colWidth)
		}
	}

	yCoordinate := self.Inner.Min.Y

	// draw rows
	for i := 0; i < len(self.Rows) && yCoordinate < self.Inner.Max.Y; i++ {
		row := self.Rows[i]
		colXCoordinate := self.Inner.Min.X
		// draw row cells
		for j := 0; j < len(row); j++ {
			col := ParseText(row[j], self.TextStyle)
			// draw row cell
			if len(col) > columnWidths[j] || self.TextAlign == AlignLeft {
				for k, cell := range col {
					if k == columnWidths[j] || colXCoordinate+k == self.Inner.Max.X {
						cell.Rune = DOTS
						buf.SetCell(cell, image.Pt(colXCoordinate+k-1, yCoordinate))
						break
					} else {
						buf.SetCell(cell, image.Pt(colXCoordinate+k, yCoordinate))
					}
				}
			} else if self.TextAlign == AlignCenter {
				xCoordinateOffset := (columnWidths[j] - len(col)) / 2
				stringXCoordinate := xCoordinateOffset + colXCoordinate
				for k, cell := range col {
					buf.SetCell(cell, image.Pt(stringXCoordinate+k, yCoordinate))
				}
			} else if self.TextAlign == AlignRight {
				stringXCoordinate := MinInt(colXCoordinate+columnWidths[j], self.Inner.Max.X) - len(col)
				for k, cell := range col {
					buf.SetCell(cell, image.Pt(stringXCoordinate+k, yCoordinate))
				}
			}
			colXCoordinate += columnWidths[j] + 1
		}

		// draw vertical separators
		separatorXCoordinate := self.Inner.Min.X
		verticalCell := NewCell(VERTICAL_LINE, NewStyle(ColorWhite))
		for _, width := range columnWidths {
			separatorXCoordinate += width
			buf.SetCell(verticalCell, image.Pt(separatorXCoordinate, yCoordinate))
			separatorXCoordinate++
		}

		yCoordinate++

		// draw horizontal separator
		horizontalCell := NewCell(HORIZONTAL_LINE, NewStyle(ColorWhite))
		if self.RowSeparator && yCoordinate < self.Inner.Max.Y && i != len(self.Rows)-1 {
			buf.Fill(horizontalCell, image.Rect(self.Inner.Min.X, yCoordinate, self.Inner.Max.X, yCoordinate+1))
			yCoordinate++
		}
	}
}
