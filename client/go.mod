module client

go 1.20

require (
	golang.org/x/net v0.7.0
	pb/test v0.0.0
)

require (
	github.com/bufbuild/connect-go v1.5.1 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace pb/test v0.0.0 => ./pb
