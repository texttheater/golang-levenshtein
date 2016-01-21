package levenshtein

var DefaultLevenshtein []EditOperation = []EditOperation{
	Match{},
	Insertion{1},
	Deletion{1},
	Substitution{1},
}

var DefaultDamerau []EditOperation = []EditOperation{
	Match{},
	Insertion{1},
	Deletion{1},
	Substitution{1},
	Transposition{1},
}

var DefaultLCS []EditOperation = []EditOperation{
	Match{},
	Insertion{1},
	Deletion{1},
}
