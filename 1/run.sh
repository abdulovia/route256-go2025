#!/bin/bash

GO_FILE="main.go"
TEST_FOLDER="dark-room"
NUM_TESTS=1

echo "Compiling $GO_FILE..."
go build -o main $GO_FILE
if [ $? -ne 0 ]; then
    echo "Compilation failed. Exiting..."
    exit 1
fi

echo "Running tests..."
for i in $(seq 1 $NUM_TESTS); do
    INPUT_FILE="$TEST_FOLDER/$i"
    OUTPUT_FILE="$TEST_FOLDER/$i.out"
    EXPECTED_FILE="$TEST_FOLDER/$i.a"

    if [ ! -f "$INPUT_FILE" ] || [ ! -f "$EXPECTED_FILE" ]; then
        echo "Test $i: Missing input or expected output file. Skipping..."
        continue
    fi

    ./main < "$INPUT_FILE" > "$OUTPUT_FILE"

    if diff -q "$OUTPUT_FILE" "$EXPECTED_FILE" > /dev/null; then
        echo "Test $i: PASSED"
    else
        echo "Test $i: FAILED"
    fi
done

rm -f main
