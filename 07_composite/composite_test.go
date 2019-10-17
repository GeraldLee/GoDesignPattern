package composite

import "testing"

func TestCompositePattern(t *testing.T) {
	triangle := &Triangle{}
	circle := &Circle{}
	drawer := Drawing{}
	drawer.Add(triangle)
	drawer.Add(circle)
	drawer.Draw("yellow")
}
