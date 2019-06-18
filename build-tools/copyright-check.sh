#!/bin/bash


IFS='' read -r -d '' LICENSE <<'EOF'
################################################################################
# Copyright 2019 IBM Corp. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
################################################################################
EOF


LICENSE_NB_LINE=`echo "$LICENSE" | wc -l`

HEADER_TO_READ=$(($LICENSE_NB_LINE +5))

ERROR=0

echo "##### Copyright check #####"
for f in `find .  \( -path ./vendor -prune \) -type f  ! -iname ".*" -o  -print`; do
  if [ ! -f "$f" ] || [ "$f" = "./build-tools/copyright-check.sh" ]; then
    continue
  fi

  FILETYPE=$(basename ${f##*.})
  case "${FILETYPE}" in
  	js | sh | go | java | rb)
  		COMMENT_PREFIX=""
  		;;
  	*)
      continue
  esac

  # Read the first 10 lines
  HEADER=`head -$HEADER_TO_READ $f`
  printf "Scanning $f\n"
  if [[ "$HEADER" != *"$LICENSE"* ]]; then
      printf "Missing or incorrect license in file $f\n"
      ERROR=1
  fi
done

echo "##### Copyright check ##### ReturnCode: ${ERROR}"
exit $ERROR
