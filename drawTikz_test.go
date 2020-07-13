package drawTikz

import "testing"

func TestStringerPoint(t *testing.T) {
	want := "(2.0000,2.0000)"
	if got := (&PointXY{2, 2}).String(); want != got {
		t.Errorf("want %s and got = %s\n", want, got)
	}
}
