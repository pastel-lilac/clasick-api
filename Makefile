DOCKER_COMPOSE_DEV_PATH=./deployments/docker-compose.dev.yaml
GOPATH=$(HOME)/go/1.13.5
BASE_PROTO_PATH=./pkg/core/apis
COMMON_PROTO_PATH=./pkg/core/apis/common/v1alpha
MUSIC_PROTO_PATH=./pkg/core/apis/music/v1alpha
ARTIST_PROTO_PATH=./pkg/core/apis/artist/v1alpha
ALBUM_PROTO_PATH=./pkg/core/apis/album/v1alpha
GENRE_PROTO_PATH=./pkg/core/apis/genre/v1alpha
PLAYLIST_PROTO_PATH=./pkg/core/apis/playlist/v1alpha

prepare-dev:
	sudo docker-compose -f $(DOCKER_COMPOSE_DEV_PATH) up -d core_db
up-core-dev:
	sudo docker-compose -f $(DOCKER_COMPOSE_DEV_PATH) build
	sudo docker-compose -f $(DOCKER_COMPOSE_DEV_PATH) up -d core
	sudo docker-compose -f $(DOCKER_COMPOSE_DEV_PATH) logs -f core
log-core-dev:
	sudo docker-compose -f $(DOCKER_COMPOSE_DEV_PATH) logs -f core
ps-dev:
	sudo docker-compose -f $(DOCKER_COMPOSE_DEV_PATH) ps
gen-proto:
	protoc -I $(BASE_PROTO_PATH) $(COMMON_PROTO_PATH)/common.proto --go_out=plugins=grpc:$(GOPATH)/src
	protoc -I $(BASE_PROTO_PATH) $(MUSIC_PROTO_PATH)/music.proto --go_out=plugins=grpc:$(GOPATH)/src
	protoc -I $(BASE_PROTO_PATH) $(ARTIST_PROTO_PATH)/artist.proto --go_out=plugins=grpc:$(GOPATH)/src
	protoc -I $(BASE_PROTO_PATH) $(ALBUM_PROTO_PATH)/album.proto --go_out=plugins=grpc:$(GOPATH)/src
	protoc -I $(BASE_PROTO_PATH) $(GENRE_PROTO_PATH)/genre.proto --go_out=plugins=grpc:$(GOPATH)/src
	protoc -I $(BASE_PROTO_PATH) $(PLAYLIST_PROTO_PATH)/playlist.proto --go_out=plugins=grpc:$(GOPATH)/src