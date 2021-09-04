package datablock

type Datablock struct {
	TimeStamp int64
	PreviousHash string
	Data         []string
	BlockHash string
}

type HashLessBlock struct {
	TimeStamp int64
	PreviousHash string
	Data         []string
}