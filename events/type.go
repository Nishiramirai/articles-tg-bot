package events

type Fetcher interface {
	Fetch(limit int)
}

type Processor interface {
	Process()
}