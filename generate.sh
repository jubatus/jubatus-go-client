#!/bin/bash -ue

JUBATUS_DIR="jubatus-generate"
JUBATUS_BRANCH="master"
CLIENT_DIR="$(cd $(dirname ${0}) && echo ${PWD})"

[ $# -eq 0 ] || JUBATUS_BRANCH="${1}"

rm -rf "${JUBATUS_DIR}"
git clone https://github.com/jubatus/jubatus.git "${JUBATUS_DIR}"
pushd "${JUBATUS_DIR}"
git checkout "${JUBATUS_BRANCH}"
popd

# Go

capitalize() {
	echo "$(echo ${1:0:1} | tr 'a-z' 'A-Z')${1:1}"
}

for DIR in "${CLIENT_DIR}/lib/"*; do
	if [ -d "${DIR}" ] && [ "$(basename "${DIR}")" != "common" ]; then
		rm -rf $DIR
	fi
done
pushd "${JUBATUS_DIR}/jubatus/server/server"
for IDL in *.idl; do
	IDL_HASH=`git log -1 --format=%H -- ${IDL}`
	IDL_VER=`git describe ${IDL_HASH}`
	NAMESPACE="$(capitalize $(basename "${IDL}" ".idl"))"
	jenerator -l go "${IDL}" -o "${CLIENT_DIR}/lib" --idl-version $IDL_VER
done
popd

rm -rf "${JUBATUS_DIR}"

find .| grep "\.go$" | xargs gofmt -w
find .| grep "\.go$" | xargs goimports -w
