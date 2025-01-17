package matrix

import "testing"

var testMatrix Rows = Rows{
	{
		{X: 0, Y: 0, NodeType: StartNode},
		{X: 1, Y: 0, NodeType: NormalNode},
		{X: 2, Y: 0, NodeType: EndNode},
	},
	{
		{X: 0, Y: 1, NodeType: WallNode},
		{X: 1, Y: 1, NodeType: WallNode},
		{X: 2, Y: 1, NodeType: WallNode},
	},
}

func TestMatrixDisplay(t *testing.T) {
	t.Run("format properly each row of the matrix", func(t *testing.T) {
		got := testMatrix.String()
		want := "\nS . E \nW W W \n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
