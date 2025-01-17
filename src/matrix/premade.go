package matrix

var Default Rows = Rows{
	{
		{X: 0, Y: 0, NodeType: StartNode},
		{X: 1, Y: 0, NodeType: NormalNode},
		{X: 2, Y: 0, NodeType: NormalNode},
	},
	{
		{X: 0, Y: 1, NodeType: WallNode},
		{X: 1, Y: 1, NodeType: NormalNode},
		{X: 2, Y: 1, NodeType: EndNode},
	},
	{
		{X: 0, Y: 2, NodeType: WallNode},
		{X: 1, Y: 2, NodeType: NormalNode},
		{X: 2, Y: 2, NodeType: WallNode},
	},
	// {
	// 	{X: 0, Y: 0, NodeType: StartNode},
	// 	{X: 1, Y: 0, NodeType: NormalNode},
	// 	{X: 2, Y: 0, NodeType: EndNode},
	// },
	// {
	// 	{X: 0, Y: 1, NodeType: WallNode},
	// 	{X: 1, Y: 1, NodeType: WallNode},
	// 	{X: 2, Y: 1, NodeType: WallNode},
	// },
}

var FiveByFive Rows = Rows{
	{
		{X: 0, Y: 0, NodeType: StartNode},
		{X: 1, Y: 0, NodeType: NormalNode},
		{X: 2, Y: 0, NodeType: WallNode},
		{X: 3, Y: 0, NodeType: NormalNode},
		{X: 4, Y: 0, NodeType: EndNode},
	},
	{
		{X: 0, Y: 1, NodeType: NormalNode},
		{X: 1, Y: 1, NodeType: WallNode},
		{X: 2, Y: 1, NodeType: NormalNode},
		{X: 3, Y: 1, NodeType: NormalNode},
		{X: 4, Y: 1, NodeType: WallNode},
	},
	{
		{X: 0, Y: 2, NodeType: NormalNode},
		{X: 1, Y: 2, NodeType: WallNode},
		{X: 2, Y: 2, NodeType: NormalNode},
		{X: 3, Y: 2, NodeType: NormalNode},
		{X: 4, Y: 2, NodeType: NormalNode},
	},
	{
		{X: 0, Y: 3, NodeType: NormalNode},
		{X: 1, Y: 3, NodeType: NormalNode},
		{X: 2, Y: 3, NodeType: WallNode},
		{X: 3, Y: 3, NodeType: NormalNode},
		{X: 4, Y: 3, NodeType: NormalNode},
	},
	{
		{X: 0, Y: 4, NodeType: WallNode},
		{X: 1, Y: 4, NodeType: NormalNode},
		{X: 2, Y: 4, NodeType: NormalNode},
		{X: 3, Y: 4, NodeType: NormalNode},
		{X: 4, Y: 4, NodeType: NormalNode},
	},
}

var FourByFour Rows = Rows{
	{
		{X: 0, Y: 0, NodeType: StartNode},
		{X: 1, Y: 0, NodeType: WallNode},
		{X: 2, Y: 0, NodeType: WallNode},
		{X: 3, Y: 0, NodeType: EndNode},
	},
	{
		{X: 0, Y: 1, NodeType: NormalNode},
		{X: 1, Y: 1, NodeType: WallNode},
		{X: 2, Y: 1, NodeType: NormalNode},
		{X: 3, Y: 1, NodeType: NormalNode},
	},
	{
		{X: 0, Y: 2, NodeType: NormalNode},
		{X: 1, Y: 2, NodeType: WallNode},
		{X: 2, Y: 2, NodeType: WallNode},
		{X: 3, Y: 2, NodeType: NormalNode},
	},
	{
		{X: 0, Y: 3, NodeType: NormalNode},
		{X: 1, Y: 3, NodeType: NormalNode},
		{X: 2, Y: 3, NodeType: NormalNode},
		{X: 3, Y: 3, NodeType: NormalNode},
	},
}
