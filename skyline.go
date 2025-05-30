package skyline

import (
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

/*
Packer implements the skyline algorithm for packing 2D rectangles.

The Initialize method must be called before adding rectangles with AddRect.

Packer is based on https://jvernay.fr/en/blog/skyline-2d-packer/implementation/.
The original C implementation can be found at https://git.sr.ht/~jvernay/JV/tree/main/item/src/jv_pack2d
*/
type Packer struct {
	width        int
	height       int
	skyline      []point
	skylineCount int
}

/*
Initialize prepares the Packer.

The width and height define the size of the rectangle in to which
rectangles will be packed.
*/
func (p *Packer) Initialize(width int, height int) {
	p.width = width
	p.height = height

	p.skylineCount = 1
	// Bottom-left point, indicating available space.
	p.skyline = []point{{x: 0, y: 0}}
}

/*
AddRect adds a rectangle to the Packer.

If the Packer has space for the new rectangle, its location is returned.
The returned x and y are the location of the bottom left point of the added rectangle.

If there is no available space for the rectangle an error is returned.

If the Packer has not been initialized (with a call to Initialize) an error will be returned.
*/
func (p *Packer) AddRect(width int, height int) (x int, y int, err error) {
	if p.skyline == nil || p.width == 0 || p.height == 0 {
		return 0, 0, NotInitialized{}
	}

	// Stores the best candidate so far.
	idxBest := math.MaxInt
	idxBest2 := math.MaxInt
	bestX := math.MaxInt
	bestY := math.MaxInt

	// Search loop for best location.
	for idx := 0; idx < p.skylineCount; idx++ {
		x := p.skyline[idx].x
		y := p.skyline[idx].y

		if width > p.width-x {
			break // We have reached the atlas' right boundary.
		}
		if y >= bestY {
			continue // We will not beat the current best.
		}

		// Raise 'y' such that we are above all intersecting skylines.
		xMax := x + width
		var idx2 int
		for idx2 = idx + 1; idx2 < p.skylineCount; idx2++ {
			if xMax <= p.skyline[idx2].x {
				break // We do not reach the next skylines.
			}
			if y < p.skyline[idx2].y {
				y = p.skyline[idx2].y // Raise 'y' as to not intersect.
			}
		}

		if y >= bestY {
			continue // We did not beat the current best.
		}
		if height > p.height-y {
			continue // We have reached the atlas' top boundary.
		}

		idxBest = idx
		idxBest2 = idx2
		bestX = x
		bestY = y
	}

	if idxBest == math.MaxInt {
		return 0, 0, NoSpace{}
	}
	if idxBest >= idxBest2 {
		return 0, 0, InternalError(fmt.Sprintf("internal error : idxBest >= idBest2 : %d >= %d", idxBest, idxBest2))
	}
	if idxBest2 <= 0 {
		return 0, 0, InternalError(fmt.Sprintf("internal error : idxBest2 <= 0 : %d <= 0", idxBest2))
	}

	// Replace the points overshadowed by the current rect, by new points.

	removedCount := idxBest2 - idxBest

	var newTopLeft point
	var newBottomRight point
	newTopLeft.x = bestX
	newTopLeft.y = bestY + height
	newBottomRight.x = bestX + width
	newBottomRight.y = p.skyline[idxBest2-1].y
	bottomRightPoint := false
	if idxBest2 < p.skylineCount {
		if newBottomRight.x < p.skyline[idxBest2].x {
			bottomRightPoint = true
		}
	} else {
		if newBottomRight.x < p.width {
			bottomRightPoint = true
		}
	}
	// TopLeft is always inserted
	insertedCount := 1
	if bottomRightPoint {
		insertedCount++
	}

	if p.skylineCount+insertedCount-removedCount > p.width {
		panic(fmt.Sprintf(
			"p.skylineCount + insertedCount - removedCount> p.width :"+
				"%d + %d - %d > %d",
			p.skylineCount, insertedCount, removedCount, p.width,
		))
	}

	if insertedCount > removedCount {
		// Expansion. Shift elements to the right. We need to iterate backwards.
		idx := p.skylineCount - 1
		idx2 := idx + (insertedCount - removedCount)
		for idx >= idxBest2 {
			for len(p.skyline) < idx2+1 {
				p.skyline = append(p.skyline, point{})
			}
			p.skyline[idx2] = p.skyline[idx]
			idx--
			idx2--
		}
		p.skylineCount += insertedCount - removedCount
	} else if insertedCount < removedCount {
		// Shrinking. Shift elements to the left. We need to iterate forwards.
		idx := idxBest2
		idx2 := idx - (removedCount - insertedCount)
		for idx < p.skylineCount {
			p.skyline[idx2] = p.skyline[idx]
			idx++
			idx2++
		}
		p.skylineCount -= removedCount - insertedCount
	}

	p.skyline[idxBest] = newTopLeft
	if bottomRightPoint {
		for len(p.skyline) < p.skylineCount {
			p.skyline = append(p.skyline, point{})
		}
		p.skyline[idxBest+1] = newBottomRight
	}

	return bestX, bestY, nil
}
