package mypackage

// CarPublic Car with public access
type CarPublic struct {
	Brand string
	Year  int
}

// Types of access in go: Public, private, protected
// Private: not accessible from other packages, it starts with lower case.
type carPrivate struct {
	brand string
	year  int
}

// PrintMessage print message
func PrintMessage() {
	println("Hello, World!")
}
