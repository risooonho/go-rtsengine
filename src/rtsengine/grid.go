package rtsengine

import (
	"errors"
	"fmt"
	"image"
	"math"
)

// Grid maintains an acre grid and its span.
type Grid struct {
	// Actual data copy of a portion of the world grid
	Matrix [][]Acre

	// Width and Height of this Grid
	Span image.Rectangle

	// Where the upper left hand corner of this grid
	// is located in world coordinates. If it is 0,0 then
	// WorldOrigin == Grid
	WorldOrigin image.Point
}

// GenerateGrid will initialize all internal structures.
// It will set the grid widith and height and situate the
// grid onto the world at worldLocation
func (grid *Grid) GenerateGrid(worldLocation image.Point, width int, height int) {
	grid.WorldOrigin = worldLocation
	grid.Span = image.Rect(0, 0, width, height)

	// allocate 2d array row per row.
	grid.Matrix = make([][]Acre, height)
	for i := range grid.Matrix {
		grid.Matrix[i] = make([]Acre, width)
	}

}

// ToGridPoint Converts world coordinates to grid coordinates
func (grid *Grid) ToGridPoint(worldPoint *image.Point) image.Point {
	return worldPoint.Sub(grid.WorldOrigin)
}

// ToWorldPoint converts grid coordinates to world coordinates
func (grid *Grid) ToWorldPoint(gridPoint *image.Point) image.Point {
	return gridPoint.Add(grid.WorldOrigin)
}

// In returns true if worldPoint is In the grid. False otherwise.
func (grid *Grid) In(worldPoint *image.Point) bool {
	return grid.ToGridPoint(worldPoint).In(grid.Span)
}

// Overlaps returns true if the other grid overlaps with this grid
func (grid *Grid) Overlaps(other *Grid) bool {
	return grid.Span.Overlaps(other.Span)
}

// Remove will eliminate a unit from the grid where-ever it is fine.
// The algorithm presently is brute force.
func (grid *Grid) Remove(unit IUnit) {
	for i := range grid.Matrix {
		for j := range grid.Matrix[i] {
			if grid.Matrix[i][j].unit == unit {
				grid.Matrix[i][j].unit = nil
			}
		}
	}
}

// Add will place the unit in the grid at location. Error is returned
// if the location is invalid. That is outside the known grid.
func (grid *Grid) Add(unit IUnit, location *image.Point) error {
	if !grid.In(location) {
		return errors.New("Location not within the world")
	}

	grid.Matrix[location.X][location.Y].unit = unit

	return nil
}

// Set the unit at locus within this grid.
func (grid *Grid) Set(locus *image.Point, unit IUnit) {
	grid.Matrix[locus.X][locus.Y].unit = unit
}

// Collision returns true if the locus is already occupied
// by any other unit OR the terrain is inaccessible such as
// Mountains and Trees.
func (grid *Grid) Collision(locus *image.Point) bool {
	acre := grid.Matrix[locus.X][locus.Y]
	return acre.unit != nil || acre.terrain == Trees || acre.terrain == Mountains
	//return false
}

// Distance between two points.
func (grid *Grid) Distance(source *image.Point, destination *image.Point) int {
	x2 := (destination.X - source.X) * (destination.X - source.X)
	y2 := (destination.Y - source.Y) * (destination.Y - source.Y)
	d2 := x2 + y2
	distance := math.Sqrt(float64(d2))

	return int(math.Trunc(distance))
}

// Print the world as ascii text.
func (grid *Grid) Print() {
	for i := range grid.Matrix {
		for j := range grid.Matrix[i] {

			switch grid.Matrix[i][j].unit.(type) {
			case *Fence:
				fmt.Printf("X")
				continue
			}

			switch grid.Matrix[i][j].terrain {
			case Trees:
				fmt.Printf("T")
			case Mountains:
				fmt.Printf("M")
			case Grass:
				fmt.Printf(".")
			default:
				fmt.Printf(".")
			}
		} //j
		fmt.Println("")
	} //i
}
