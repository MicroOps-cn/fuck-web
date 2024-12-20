#!/bin/bash


awk_group_by_go_package='{
  pkgs[$2]=1;
  protos[$1]=$2
}END{
  for(pkg in pkgs){
    for(proto in protos){
      if(pkg==protos[proto]){
        printf("%s ",proto)
      }
    };
    print("")
  }
}'


if [ -z "$PROTO_DEFS" ] ;then
    echo "PROTO_DEFS not defined"
    exit 1
fi

cd $(dirname $0)/..

TMP_DIR="gogo_out"
mkdir -p "${TMP_DIR}"
trap "rm -rf '${TMP_DIR}'" SIGINT SIGQUIT EXIT
grep -HoP '(?<=option go_package = ")[^;]+' ${PROTO_DEFS} | awk -F: "${awk_group_by_go_package}" | while read line; do
    log=$(${PROTOC} ${PROTOC_OPTS} $line 2>&1)
    if [ $? -ne 0 ]; then
        echo ">> ${PROTOC} ${PROTOC_OPTS} $line"
        echo "$log"
        rm -rf "${TMP_DIR}"
        exit 1
    fi
done
find gogo_out -type f \( -name '*.pb.go' -o -name '*.pb.gw.go' \) -print | while read tmp_path; do
    sed -i ':label;N;s/\nvar E_\S\+ = gogoproto.E_\S\+\n//;b label' ${tmp_path}
    sed -i '/gogoproto "github.com\/gogo\/protobuf\/gogoproto/d' ${tmp_path}
    filename=${tmp_path#"${TMP_DIR}/${GOMODULENAME}/"}
    if [ "$filename" == "${tmp_path}" ]; then
        echo "$filename"
        echo "ERROR: incorrect path of protobuf output file: ${tmp_path}"
        exit 1;
    fi
    mv "${tmp_path}" "${filename}"
done

