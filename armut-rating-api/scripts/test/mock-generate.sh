#!/bin/bash

echo "STARTING TO GENERATE MOCKS"

files=$(find ../../internal -name "*.go")

for file in $files; do
  if ! grep -E -q "([A-z]*(test|model|response|.receiver|handler|\.pb).go)" <<< "$file"; then
    filtered_files+=("${file}")
  fi
done

for file in "${filtered_files[@]}"; do
  mock_file=$(echo "${file}" | sed -e "s/\.go/_mock.go/")
  package_name=$(basename $(dirname $mock_file))

  mockgen -source="$file" -destination "$mock_file" -package "$package_name"
  echo "${file} MOCKED WITH THIS PATH: ${mock_file}"
done

echo "END"
