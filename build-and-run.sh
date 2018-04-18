#!/usr/bin/env bash

echo "Building Project"

/usr/local/go/bin/go build -o "app"

echo "Starting Code Execution"

echo "2
3 3 2 1
1 2
3 1
2 3
6 6 2 5
1 3
3 4
2 4
1 2
2 3
5 6
" | ./app

echo "3
6 4 2 1
1 2
2 3
4 5
5 6
7 6 2 1
1 2
2 3
3 4
2 4
5 6
6 7
9 7 3 4
1 2
1 5
2 4
2 3
3 9
4 3
6 8" | ./app

echo "Done. Deleting executable."

rm app

