module github.com/sawvant/sawvant-go

go 1.22

require github.com/stretchr/testify v1.12.1

require go.yaml.in/yaml/v3 v3.0.5 // indirect

// v1.1.3 was re-tagged after publication, so sum.golang.org holds a different
// checksum and `go get` fails. v1.0.0 and v1.1.0 were tagged and then removed.
retract (
	v1.1.3
	v1.1.0
	v1.0.0
)
