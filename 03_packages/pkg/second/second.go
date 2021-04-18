package second

//CreateInt returned pointer on int variable
func CreateInt() *int {
	var x = 0
	return &x
}
