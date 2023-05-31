module github.com/celestiaorg/go-cnc

go 1.19

require (
	github.com/go-resty/resty/v2 v2.7.0
	github.com/gogo/protobuf v1.3.3
	github.com/google/uuid v1.3.0
	github.com/stretchr/testify v1.8.3
	github.com/testcontainers/testcontainers-go v0.15.0
)

require (
	github.com/Microsoft/hcsshim v0.9.4 // indirect
	github.com/celestiaorg/merkletree v0.0.0-20210714075610-a84dc3ddbbe4 // indirect
	github.com/celestiaorg/rsmt2d v0.9.0 // indirect
	github.com/containerd/cgroups v1.0.4 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/klauspost/cpuid/v2 v2.1.1 // indirect
	github.com/klauspost/reedsolomon v1.11.1 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/moby/sys/mount v0.3.3 // indirect
	github.com/moby/sys/mountinfo v0.6.2 // indirect
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5 // indirect
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/tendermint/tendermint v0.34.24 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/mod v0.9.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/celestiaorg/celestia-app v1.0.0-rc0
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/containerd/containerd v1.6.8 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/docker v23.0.6+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/moby/term v0.0.0-20220808134915-39b0c02b01ae // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0-rc2 // indirect
	github.com/opencontainers/runc v1.1.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
	google.golang.org/grpc v1.52.0 // indirect
	google.golang.org/protobuf v1.28.2-0.20220831092852-f930b1dc76e8 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/docker/docker => github.com/docker/docker v20.10.17+incompatible
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/tendermint => github.com/celestiaorg/celestia-core v1.21.1-tm-v0.34.27
)
