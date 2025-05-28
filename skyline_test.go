package skyline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Some of these tests are ported from these tests.
// https://git.sr.ht/~jvernay/JV/tree/main/item/src/jv_pack2d/jv_pack2d_test.c

const worstcaseAtlasSize = 512

func TestPackerInternalLogic(t *testing.T) {
	p := &Packer{}
	p.Initialize(4, 4)

	add := func(
		width int, height int,
		expectedNilError bool,
		expectedX int, expectedY int,
		expectedSkylineCount int, expectedSkyline []point,
	) {
		x, y, err := p.AddRect(width, height)
		assert.Equal(t, expectedNilError, err == nil)
		if err == nil {
			assert.Equal(t, expectedX, x)
			assert.Equal(t, expectedY, y)
			assert.Equal(t, expectedSkylineCount, p.skylineCount)
			assert.Equal(t, expectedSkyline, p.skyline[:len(expectedSkyline)])
		}
	}

	add(1, 1, true, 0, 0, 2, []point{{0, 1}, {1, 0}})
	add(2, 2, true, 1, 0, 3, []point{{0, 1}, {1, 2}, {3, 0}})
	add(4, 1, true, 0, 2, 1, []point{{0, 3}})
	add(1, 2, false, 0, 0, 0, []point{})
	add(3, 1, true, 0, 3, 2, []point{{0, 4}, {3, 3}})
	add(2, 1, false, 0, 0, 0, []point{})
	add(1, 1, true, 3, 3, 2, []point{{0, 4}, {3, 4}})
	add(1, 1, false, 0, 0, 0, []point{})
	assertPackerFull(t, p)
}

func TestPackerBlogExample(t *testing.T) {
	p := &Packer{}
	p.Initialize(8, 8)

	add := func(width int, height int, expectedX int, expectedY int) {
		x, y, err := p.AddRect(width, height)
		assert.Nil(t, err)
		assert.Equal(t, expectedX, x)
		assert.Equal(t, expectedY, y)
	}

	add(3, 3, 0, 0) // A
	add(3, 1, 3, 0) // B
	add(4, 2, 3, 1) // C
	add(1, 4, 7, 0) // D
	add(2, 5, 0, 3) // E
	add(3, 2, 2, 3) // F
	add(2, 3, 5, 3) // G
	add(3, 2, 2, 5) // H
	add(4, 1, 2, 7) // I
	add(2, 2, 6, 6) // J
}

func TestPackerWorstCaseWidth(t *testing.T) {
	p := &Packer{}
	p.Initialize(worstcaseAtlasSize, 2)

	for range p.height {
		for range p.width {
			_, _, err := p.AddRect(1, 1)
			assert.Nil(t, err)
		}
	}

	assertPackerFull(t, p)
}

func TestPackerWorstCaseHeight(t *testing.T) {
	p := &Packer{}
	p.Initialize(2, worstcaseAtlasSize)

	for range p.height {
		for range p.width {
			_, _, err := p.AddRect(1, 1)
			assert.Nil(t, err)
		}
	}

	assertPackerFull(t, p)
}

func TestPackerWorstCaseDiagonalVertical(t *testing.T) {
	p := &Packer{}
	p.Initialize(worstcaseAtlasSize, worstcaseAtlasSize)

	for i := range worstcaseAtlasSize - 1 {
		x, y, err := p.AddRect(1, worstcaseAtlasSize-1-i)
		assert.Nil(t, err)
		assert.Equal(t, i, x)
		assert.Equal(t, 0, y)
	}

	for i := range worstcaseAtlasSize {
		x, y, err := p.AddRect(1, worstcaseAtlasSize-i)
		assert.Nil(t, err)
		assert.Equal(t, worstcaseAtlasSize-i-1, x)
		assert.Equal(t, i, y)
	}

	assertPackerFull(t, p)
}

func TestPackerWorstCaseDiagonalHorizontal(t *testing.T) {
	p := &Packer{}
	p.Initialize(worstcaseAtlasSize, worstcaseAtlasSize)

	for i := range worstcaseAtlasSize {
		width := worstcaseAtlasSize - 1 - i
		if width > 0 {
			x, y, err := p.AddRect(width, 1)
			assert.Nil(t, err)
			assert.Equal(t, 0, x)
			assert.Equal(t, i, y)
		}

		x, y, err := p.AddRect(1+i, 1)
		assert.Nil(t, err)
		assert.Equal(t, worstcaseAtlasSize-1-i, x)
		assert.Equal(t, i, y)
	}

	assertPackerFull(t, p)
}

func assertPackerFull(t *testing.T, p *Packer) {
	for i := range p.skylineCount {
		assert.Equal(t, p.height, p.skyline[i].y)
	}
}
