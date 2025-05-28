package skyline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Some of these tests are ported from these tests.
// https://git.sr.ht/~jvernay/JV/tree/main/item/src/jv_pack2d/jv_pack2d_test.c

func TestPackerInternalLogic(t *testing.T) {
	p := &Packer{}
	p.Initialize(4, 4)

	testAdd(t, p, 1, 1, true, 0, 0, 2, []point{{0, 1}, {1, 0}})
	testAdd(t, p, 2, 2, true, 1, 0, 3, []point{{0, 1}, {1, 2}, {3, 0}})
	testAdd(t, p, 4, 1, true, 0, 2, 1, []point{{0, 3}})
	testAdd(t, p, 1, 2, false, 0, 0, 0, []point{})
	testAdd(t, p, 3, 1, true, 0, 3, 2, []point{{0, 4}, {3, 3}})
	testAdd(t, p, 2, 1, false, 0, 0, 0, []point{})
	testAdd(t, p, 1, 1, true, 3, 3, 2, []point{{0, 4}, {3, 4}})
	testAdd(t, p, 1, 1, false, 0, 0, 0, []point{})
	assertPackerFull(t, p)
}

func testAdd(t *testing.T, p *Packer,
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

func assertPackerFull(t *testing.T, p *Packer) {
	for i := range p.skylineCount {
		assert.Equal(t, p.height, p.skyline[i].y)
	}
}
