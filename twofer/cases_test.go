package twofer

type testCase struct {
	description, input, expected string
}

var testCases = []testCase{
	{

		description: "empty name",

		input: "",

		expected: "One for you, one for me.",
	},
	{

		description: "name is Alice",

		input: "Alice",

		expected: "One for Alice, one for me.",
	},
	{

		description: "name is Bob",

		input: "Bob",

		expected: "One for Bob, one for me.",
	},
}
