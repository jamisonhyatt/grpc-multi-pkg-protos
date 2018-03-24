#! /bin/bash
set -e

SwiftGen=""
if [[ "$1" == "swift" ]]; then
    SwiftGen="--swift_out=./gen "
    rm -rf swift
    mkdir -p swift
    echo "Generating Swift code"
fi

JavaNanoGen=""
if [[ "$1" == "java" ]]; then
    JavaNanoGen="--javanano_out=ignore_services=true:./gen "
    rm -rf java
    mkdir -p java
    echo "Generating Java code"
fi


typeOverrides="Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api"
includes="-I /go/src/github.com/grpc-ecosystem/grpc-gateway/ -I /go/src -I /usr/local/include/ -I /usr/include/"

function cleanTempGeneratedPaths() {
    #find's -delete won't delete things that aren't empty; this deletes the files inside the path, and the folders in the path before attemping to delete the 'gen' directories
    find . -path '*/gen/*' -delete
    find . -iname "gen" -type d -delete
}
cleanTempGeneratedPaths
mkdir -p pkg
#Delete existing generated code and swagger
find ./pkg -iname "*.go" -type f -delete
find . -iname "*.json" ! -iname "launch.json" -type f -delete
find . -iname "*.protoset" -type f -delete

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
    docker run --rm -i -v $(pwd)/external/${dependency}:/protos --entrypoint /bin/bash  -w /protos namely-protoc:latest -s <<EOF
    mkdir -p gen && protoc -I. ${includes} *.proto \
        --gogo_out=plugins=grpc,${typeOverrides}:./gen ${SwiftGen} ${JavaNanoGen}
        
EOF
    find ./external/${dependency} -iname "*.go" -exec mv {} ../pkg/external/${dependency}  \;
    find ./external/${dependency} -iname "*.swift" -exec mv {} ../swift  \;
    find ./external/${dependency} -iname "*.java" -exec mv {} ../java  \;
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
    docker run -i --rm -v $(pwd):/protos --entrypoint /bin/bash  -w /protos namely-protoc:latest -s <<EOF 
    echo "Generating go code for service ${service}"
    mkdir -p gen 
    protoc -I/protos/${service} -I /protos/external ${includes} /protos/${service}/*.proto --grpc-gateway_out=logtostderr=true,${typeOverrides}:. \
    --gogo_out=plugins=grpc,${typeOverrides}:./gen --descriptor_set_out=${service}_svc.protoset --include_imports ${SwiftGen} ${JavaNanoGen}
EOF
    find ./ -iname "*.go" -exec mv {} ../pkg/${app}/${service}  \;
    find ./ -iname "*.protoset" -exec mv {} ../protosets/ \;
    find ./ -iname "*.swift" -exec mv {} ../swift  \;
    find ./ -iname "*.java" -exec mv {} ../java  \;
done
echo "Done with services"

## Build Base Svc
docker run -i --rm -v $(pwd):/protos --entrypoint /bin/bash  -w /protos namely-protoc:latest -s <<EOF
  echo "Building base ${app} service"
  mkdir -p gen
  protoc -I. -I /protos/external ${app}_svc.proto ${includes} --grpc-gateway_out=logtostderr=true,${typeOverrides}:. \
  --gogo_out=plugins=grpc,${typeOverrides}:./gen --descriptor_set_out=${app}_svc.protoset --include_imports ${SwiftGen} ${JavaNanoGen}
EOF
find ./ -iname "*.go" -exec mv {} ../pkg/${app}/  \;
find ./ -iname "*.protoset" -exec mv {} ../protosets/ \;
find ./ -iname "*.swift" -exec mv {} ../swift  \;
find ./ -iname "*.java" -exec mv {} ../java  \;

## Fix types empty
find ../pkg -iname "*.pb.gw.go" -exec sed -i '' 's/empty.Empty/types.Empty/g' {} \;

#cleanup
cleanTempGeneratedPaths
