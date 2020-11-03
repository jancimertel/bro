/*
Example of file multiline comment that is stored at the top of mainfile1.go file in rootpkg package.
Really just basic functionality for now

*/
package rootpkg

// TryNoParam - basic line comment
func TryNoParam() bool {
	// comment in function
	return false
}

/*
TryOneParam - multiline comment
This function has only one input parameter and returns hardcoded false
*/
func TryOneParam(b bool) bool {
	return false
}
