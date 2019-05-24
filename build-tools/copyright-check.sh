#!/bin/bash

# Licensed Materials - Property of IBM
# (c) Copyright IBM Corporation 2017, 2019. All Rights Reserved.
# Note to U.S. Government Users Restricted Rights:
# Use, duplication or disclosure restricted by GSA ADP Schedule
# Contract with IBM Corp.

YEAR="2017, 2019"

#LINE1="${COMMENT_PREFIX}Licensed Materials - Property of IBM"
CHECK1="# Licensed Materials - Property of IBM"
#LINE2="${COMMENT_PREFIX}(c) Copyright IBM Corporation ${YEAR}. All Rights Reserved."
CHECK2=" Copyright IBM Corporation ${YEAR}. All Rights Reserved."
#LINE3="${COMMENT_PREFIX}Note to U.S. Government Users Restricted Rights:"
CHECK3="# U.S. Government Users Restricted Rights -"
#LINE4="${COMMENT_PREFIX}Use, duplication or disclosure restricted by GSA ADP Schedule"
CHECK4=" Use, duplication or disclosure restricted by GSA ADP"
#LINE5="${COMMENT_PREFIX}Contract with IBM Corp."
CHECK5="#  IBM Corporation - initial API and implementation"

#LIC_ARY to scan for
LIC_ARY=("$CHECK1" "$CHECK2" "$CHECK3" "$CHECK4" "$CHECK5")


ERROR=0

echo "##### Copyright check #####"
for f in `find .  \( -path ./vendor -prune \) -o \( -path ./api/i18n/i18nBinData -prune \)  -o \( -path ./pkg/errors -prune \)-type f  ! -iname ".*" -o  -print`; do
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
  HEADER=`head -10 $f`
  printf "Scanning $f\n"
  for i in `seq 0 4`; do
    if [[ "$HEADER" != *"${LIC_ARY[$i]}"* ]]; then
      printf "Missing [${LIC_ARY[$i]}] file $f\n"
      ERROR=1
      break
    fi
  done
done

echo "##### Copyright check ##### ReturnCode: ${ERROR}"
exit $ERROR
