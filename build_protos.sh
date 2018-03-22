#! /bin/bash
set -e


if [[ "$(docker images -q protoc-bash 2> /dev/null)" == "" ]]; then
    docker build --rm --squash -t protoc-bash . -f-<<DOCKERFILE
FROM znly/protoc:latest

RUN apk update && apk add bash
DOCKERFILE
fi

function cleanTempGeneratedPaths() {
    #find's -delete won't delete things that aren't empty; this deletes the files inside the path, and the folders in the path before attemping to delete the 'gen' directories
    find . -path '*/gen/*' -delete
    find . -iname "gen" -type d -delete
}
cleanTempGeneratedPaths
mkdir -p pkg
#Delete existing generated code and swagger
find ./pkg -iname "*.go" -type f -delete
find . -iname "*.json" -type f -delete

dir=$(pwd | xargs basename)
if [[ dir != "protos" ]]; then 
    cd protos
fi
export app="weatherman"
## Build External Dependencies
externalDependencies=(
    location 
    weather
    )
for dependency in ${externalDependencies[@]} 
do
    echo "Generating go code for external dependency ${dependency}"
    mkdir -p ../pkg/external/${dependency}
    docker run --rm -v $(pwd)/external/${dependency}:/protos --entrypoint /bin/bash  -w /protos protoc-bash -c 'mkdir -p gen && protoc -I . *.proto --go_out=plugins=grpc:./gen'
    find ./external/${dependency} -iname "*.go" -exec mv {} ../pkg/external/${dependency}  \;
done
echo "Done with external dependencies"

# Svcs
services=(
    desktop_svc
    mobile_svc
)
for service in ${services[@]}
do


    mkdir -p ../pkg/weatherman/${service}
    docker run -i --rm -v $(pwd):/protos --entrypoint /bin/bash  -w /protos protoc-bash -s <<EOF 
    echo "Generating go code for service ${service}"
    mkdir -p gen 
    protoc -I/protos/${service} -I /protos/external /protos/${service}/*.proto --go_out=plugins=grpc:./gen
EOF
    find ./ -iname "*.go" -exec mv {} ../pkg/${app}/${service}  \;
done
echo "Done with services"

## Build Base Svc
echo "Building base ${app} service"
docker run --rm -v $(pwd):/protos --entrypoint /bin/bash  -w /protos protoc-bash -c 'mkdir -p gen && protoc -I. -I /protos/external weatherman_svc.proto --go_out=plugins=grpc:./gen'
find ./ -iname "*.go" -exec mv {} ../pkg/${app}/  \;

#cleanup
cleanTempGeneratedPaths
cd ..
