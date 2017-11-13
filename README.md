# builder
code generator for a java builder class

```
NAME:
   builder - a java builder class code generator

USAGE:
   builder [global options] command [command options] [arguments...]

VERSION:
   0.0.1

DESCRIPTION:
   generate code for a java builder class

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --string name, -s name           string variable by name
   --int name, -i name              int variable by name
   --integer name, -I name          Integer variable by name
   --boolean name, -b name          boolean variable by name
   --Boolean name, -B name          Boolean variable by name
   --custom type                    define a variable with a custom type, the var name will be lowercase custom name
   --string-list name, --sl name    string list variable by name
   --integer-list name, --IL name   Integer list variable by name
   --int-array name, --ia name      int array variable by name
   --boolean-array name, --ba name  boolean array variable by name
   --Boolean-list name, --BL name   Boolean list variable by name
   --custom-list type               variable with a custom type, the var name will be lowercase custom name
   --class value, -c value          class to generate builder for (default: "CLAZZ")
   --constructor                    include the parent classes constructor (default: false)
   --help, -h                       show help (default: false)
   --version, -v                    print the version (default: false)
```

builder takes a number of flags, each flag represents a single variable for the builder class.

## Example:
This command creates a builder class for the `BookTitle` parent class, with a single string variable `title`
```
builder -s title --class BookTitle
```
Gives the output
```
public static class Builder {

    private String title;

    public Builder title(final String title) {
        this.title = title;
        return this;
    }

    public BookTitle build() {
        return new BookTitle(this);
    }
}

```
The `--constructor` flag will generate the constructor for the parent class:
```
builder -s title --class BookTitle --constructor
```
Gives the output
```

public BookTitle(final Builder builder) { 
    this.title = builder.title;
}

public static class Builder {

    private String title;

    public Builder title(final String title) {
        this.title = title;
        return this;
    }

    public BookTitle build() {
        return new BookTitle(this);
    }
}

```