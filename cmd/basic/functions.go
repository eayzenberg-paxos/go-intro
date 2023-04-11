package basic

func MyExportedFunction(arg1 int, arg2 string) (bool, string, error) {

	/* In Go, export functions are functions that are accessible from outside of the
	package in which they are defined. By convention, an exported function's name
	starts with a capital letter.
	Exported functions can be used by other packages or programs, and are an
	important way of encapsulating functionality and providing a clean API to other
	developers. */

	return true, "Alex", nil
}

// Example from Jing. Named return variables

/*
func ReadFull(r Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}
*/
