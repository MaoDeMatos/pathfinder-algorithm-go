package main

var (
	// Fixed/manual matrices
	simpleInput = Input{
		{
			Node{X: 0, Y: 0, Type: StartNode},
			Node{X: 1, Y: 0, Type: NormalNode},
			Node{X: 2, Y: 0, Type: EndNode},
		},
	}

	fourByFourInput = Input{
		{
			Node{X: 0, Y: 0, Type: StartNode},
			Node{X: 1, Y: 0, Type: NormalNode},
			Node{X: 2, Y: 0, Type: WallNode},
			Node{X: 3, Y: 0, Type: NormalNode},
		},
		{
			Node{X: 0, Y: 1, Type: NormalNode},
			Node{X: 1, Y: 1, Type: WallNode},
			Node{X: 2, Y: 1, Type: NormalNode},
			Node{X: 3, Y: 1, Type: NormalNode},
		},
		{
			Node{X: 0, Y: 2, Type: NormalNode},
			Node{X: 1, Y: 2, Type: NormalNode},
			Node{X: 2, Y: 2, Type: WallNode},
			Node{X: 3, Y: 2, Type: WallNode},
		},
		{
			Node{X: 0, Y: 3, Type: WallNode},
			Node{X: 1, Y: 3, Type: NormalNode},
			Node{X: 2, Y: 3, Type: NormalNode},
			Node{X: 3, Y: 3, Type: EndNode},
		},
	}

	fiveByFiveInput = Input{
		{
			Node{X: 0, Y: 0, Type: StartNode},
			Node{X: 1, Y: 0, Type: NormalNode},
			Node{X: 2, Y: 0, Type: WallNode},
			Node{X: 3, Y: 0, Type: NormalNode},
			Node{X: 4, Y: 0, Type: NormalNode},
		},
		{
			Node{X: 0, Y: 1, Type: NormalNode},
			Node{X: 1, Y: 1, Type: WallNode},
			Node{X: 2, Y: 1, Type: NormalNode},
			Node{X: 3, Y: 1, Type: NormalNode},
			Node{X: 4, Y: 1, Type: NormalNode},
		},
		{
			Node{X: 0, Y: 2, Type: NormalNode},
			Node{X: 1, Y: 2, Type: NormalNode},
			Node{X: 2, Y: 2, Type: NormalNode},
			Node{X: 3, Y: 2, Type: WallNode},
			Node{X: 4, Y: 2, Type: NormalNode},
		},
		{
			Node{X: 0, Y: 3, Type: WallNode},
			Node{X: 1, Y: 3, Type: NormalNode},
			Node{X: 2, Y: 3, Type: NormalNode},
			Node{X: 3, Y: 3, Type: WallNode},
			Node{X: 4, Y: 3, Type: EndNode},
		},
	}

	// Errorring
	errorNoEndInput = Input{
		{
			Node{X: 0, Y: 0, Type: StartNode},
			Node{X: 1, Y: 0, Type: NormalNode},
			Node{X: 2, Y: 0, Type: NormalNode},
		},
	}

	errorNoPathInput = Input{
		{
			Node{X: 0, Y: 0, Type: WallNode},
			Node{X: 1, Y: 0, Type: StartNode},
		},
		{
			Node{X: 0, Y: 1, Type: EndNode},
			Node{X: 1, Y: 1, Type: WallNode},
		},
	}
)
