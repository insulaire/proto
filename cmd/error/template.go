package main

type ErrorCode int64

func (err ErrorCode) Error() string {
	return err.String()
}

func (err ErrorCode) Code() int64 {
	return int64(err)
}

func (err ErrorCode) String() string {
	return ""
}

type ErrorInfo struct {
	Code int64
	Msg  string
}
