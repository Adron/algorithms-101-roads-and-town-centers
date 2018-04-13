#!/usr/bin/env bash

echo "Build Project\n"

/usr/local/go/bin/go build -o "app"

echo "Starting Code Execution\n"

echo "2 \
3 3 2 1 \
1 2 \
3 1 \
2 3 \
6 6 2 5 \
1 3 \
3 4 \
2 4 \
1 2 \
2 3 \
5 6 \
" | ./app

echo "Results Correct? They should be 4, and then subsequently 12."