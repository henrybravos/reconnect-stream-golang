module server

go 1.20

require (
	github.com/bufbuild/connect-go v1.5.1
	golang.org/x/net v0.6.0
	pb/test v0.0.0
)

require (
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace pb/test v0.0.0 => ./pb
