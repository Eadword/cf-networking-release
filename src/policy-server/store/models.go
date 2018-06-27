package store

type Policy struct {
	Source      Source
	Destination Destination
}

type Source struct {
	ID   string
	Tag  string
	Type string
}

type Destination struct {
	ID       string
	Tag      string
	Protocol string
	Port     int
	Ports    Ports
	Type     string
	IPs      []IPRange
}

type IPRange struct {
	Start string
	End   string
}

type Ports struct {
	Start int
	End   int
}

type Tag struct {
	ID   string
	Tag  string
	Type string
}
