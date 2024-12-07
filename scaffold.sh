#!/bin/bash

read -p "Enter day number: " DAY

DAY="day$DAY"

mkdir -p $DAY solutions/$DAY
touch inputs/$DAY.txt

cat <<EOF > $DAY/$DAY.go
package $DAY

EOF

cat <<EOF > "$DAY/${DAY}_test.go"
package ${DAY}_test

EOF

cat <<EOF > solutions/$DAY/main.go
package main

func main() {

}
EOF
