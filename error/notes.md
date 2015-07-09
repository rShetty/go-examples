Flow of the talk:

- Go has a built in error type
- It uses error values to indicate abnormal state
- Simple err handling example simple_error.go
- error type is an interface type error_1.go
- error type predeclared in the universe block
- errors(errorString) is the error implementation most commonly used
- errors.New and examples error_2.method
- fmt package formats the error type by calling Error() method
- To add useful information use fmt.Errorf
- Custom error types 
- Verbose code because of error handling
- Code samples


Defer/Panic/Recover

- defer pushes a functional call on to a list
- Executed after surrounding function returns
- Clean up operations
- Examples
- Behiavior and properties

Panic:
- Stops flow of control and starts returning to its caller
- Runtime panics
- Recover -regains control inside defer
- error and unmarshal method decode.go

