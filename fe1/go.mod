module github.com/platinasystems/goes-platina-mk1/fe1

require (
	github.com/docker/go-units v0.3.3 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.8.0 // indirect
	github.com/platinasystems/fe1 v0.0.0-20181015210308-3790cbed6f50
	github.com/platinasystems/firmware-fe1a v0.0.0-20181017185739-2d160f42cff1
	github.com/platinasystems/go v0.0.0-20181011011039-90c571c632a1
)

replace github.com/platinasystems/fe1 => ../../fe1

replace github.com/platinasystems/firmware-fe1a => ../../firmware-fe1a

replace github.com/platinasystems/go => ../../go
