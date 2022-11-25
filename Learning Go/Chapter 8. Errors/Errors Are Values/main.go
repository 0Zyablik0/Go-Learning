package main

type Status int

type StatusErr struct {
	Status  Status
	Message string
}

const (
	Error1 = iota + 1
	Error2
	Error3
	Error4
)

func (se *StatusErr) Error() string {
	return se.Message
}

func process(i int) error {
	if i == 0 {
		return &StatusErr{
			Status:  Error1,
			Message: "value is zero",
		}
	}

	if i < 0 {
		return &StatusErr{
			Status:  Error2,
			Message: "value is negative",
		}
	}

	if i%2 == 0 {
		return &StatusErr{
			Status:  Error3,
			Message: "value is even",
		}
	}
	return nil
}
