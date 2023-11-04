
#  go mod init

### Enable dependency tracking for your code
  when your code imports packages contained in other modules, you manage those dependencies through your code's own module.
  That module is defined by a `go.mod` file that tracks the modules that provide those packages.

  That go.mod file stays with your code, including in your source code repository.

  To enable dependency tracking for your code by creating a `go.mod` file, run the go mod init command, giving it the name of the module your code will be in.
  The name is the module's module path

### Naming a module
  when you run go mod init to create a module for tracking dependencies, you specify a module path that serves as the module's name.

  The module path becomes the import path prefix for packages in the module.

  Be sure to specify a module path that won't conflict with the module path of other modules.

  At a minimum, a module path need only indicate something about its origin, such as a company or author name.

  But the path might also be more descriptive about what the module is  or does.

`<prefix>/<descriptive-text>`

1. the prefix is typically a string that partially describes the module, such as string that describes its origin.