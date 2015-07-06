## Error handling in Go

- Errors are values
- Used to indicate abnormal state

- Error type is an interface type
- 'error' type with all built in types is predeclared in the universe block
- 'errors' package unexported errorString type 
- fmt package formats the error type by calling Error method
- fmt.Errorf formats string and returns it as error created by errors.New
- Custom error type implementation
- Handling of errors if err != nil - Idiomatic and avoiding repeatation
- panic and recover
