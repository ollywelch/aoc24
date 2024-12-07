#!/bin/bash

read -p "Enter day number: " DAY

DAY="day$DAY"

mkdir -p $DAY solutions/$DAY
touch inputs/$DAY.txt

cat <<EOF > $DAY/$DAY.go
package $DAY

EOF

cat <<EOF > $DAY/$DAY_test.go
package $DAY_test

EOF

cat <<EOF > solutions/$DAY/main.go
package main

func main() {

}
EOF
