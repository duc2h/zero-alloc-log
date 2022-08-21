define bench
	go test -cpu=1,2,5 -benchmem -benchtime=5s -bench . $(1)
endef

define check-espace
	go build -gcflags '-m' $(1)
endef

bench-poollog:
	$(call bench,./poollog/...)
bench-normallog:
	$(call bench,./normallog/...)
check-espace-poollog:
	$(call check-espace,./poollog/...)
check-espace-normallog:
	$(call check-espace,./normallog/...)

all: bench-poollog check-espace-poollog bench-normallog check-espace-normallog