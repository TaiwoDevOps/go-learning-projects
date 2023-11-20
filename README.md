## go_learning_projects

    - This directory contains all GoLang learning projects
    - My summary of programming: we have data, that we need to specify as a
    particular Type and hold in a variable, then we can perform different
    OPERATION on that VARIABLE based on its TYPE of DATA.
                            (D.T.V.O)

## Go Lang Programming concepts :

    -- packages (multiple Go files) and an application can have multiple packages as well.
    To use a function from a different package, it first needs to be exported (naming the function with a Capital letter) and imported into the package file that needs it.

    -- fmt package - this handles printing (fmt.Print, fmt.Println) and formatted printing (fmt.Printf) using place holders as string interpolation.. eg. %v for values %T for runtime data type.
    Note: with the fmt package, you can use the Scan method to handle I/O (Input and Output).

    -- types - Go is a statically, dynamic type with standard types in Go are: string, int(negative and positive whole numbers), uint(for positive numbers), float (for fractional numbers), booleans, structs {key, value} pair, map[keyType]valueType, [30]array(fixed size of elements), []slice (not fixed in size)
                                        ^ ^
         Variables depends              | |
          on types                      | |
                                        ^ ^
    -- variables (global and local) - this a container for holding values of the same type ( which are stored in memory) and that value can change during the lifecycle of that application
    Note; Variable scope is based on how a variable is declared and used.
                                        ^ ^
    Note: Based on the use case         | |
     (*) can either be an operator      | |
     to get the value of a pointer      | |
        or just specifying the type     | |
        of that variable in an argument | |
                                        ^ ^
    -- pointers (aka special &variables) - this is a variable that holds the memory address of that variable in other to reference (*pointerToVariable) the value


    -- func (encapsulation using functions)- with functions, we encapsulate logic and modularize the application. functions can be called within another function

### Conditional flows

    -- loops - this is used to iterate through slices and arrays using range (making it definite), or can be used without range and therefore runs indefinitely until exited manually.


    -- if, if else, else if-, switch

### Logical operator

    &&, ||,

### Comparison operator

    >,<, <=, >=, ==, !=

## Go architecture

    -- Concurrency (go asyncFunc)- this is an idea in Go for handling async tasks using multi-threaded approach. So instead of running all tasks (sync and async tasks on  a single thread, we use concurrency (go) to spin up another thread when needed and avoid blocking )

#

-> ## <-
Note: To execute Go files with the same package name, use go run . from the base directory of that application, this is an alternative to specifying each file for that package as go run main.go anotherFile.go
