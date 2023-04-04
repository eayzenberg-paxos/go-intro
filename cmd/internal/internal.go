package internal

/*
In Go, an internal package is a package that is only visible to other packages
within the same module. Internal packages are a way to structure your code and
enforce encapsulation within your project.

Internal packages are identified by their location within the module's directory
structure. An internal package is a directory that is located inside the
internal directory at any level, but not directly under src. For example, if
your module is located at /path/to/my/module, you could have an internal package
located at /path/to/my/module/internal/pkg.

The contents of an internal package are treated the same as any other package,
and can include Go source code, test files, and documentation. However, internal
packages are only visible to other packages within the same module, and cannot
be imported by packages outside of the module.
 */
