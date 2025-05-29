package skyline

// InternalError means that an internal logic error occurred.
// It indicates a bug in [Packer].
type InternalError string

func (err InternalError) Error() string {
	return string(err)
}

// NotInitialized means that a [Packer] instance has not been initialized with a call to Initialize.
type NotInitialized struct{}

func (err NotInitialized) Error() string {
	return string("packer is unitialised")
}

// NoSpace means that there is unsufficient space in a [Packer] to add a rectangle.
type NoSpace struct{}

func (err NoSpace) Error() string {
	return string("no space available")
}
