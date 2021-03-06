#! /bin/bash

set -e

scripts=*.mx3

mumax3 -vet $scripts


for f in $scripts; do
	mumax3 -cache /tmp -f -http "" $f;
	echo ""
done;

for g in *.go; do
	if [ "$g" != "doc.go" ]; then
		echo go run $g;
		go run $g;
	fi
done
