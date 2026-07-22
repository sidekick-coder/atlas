#!/usr/bin/env sh

# create_id.sh
# Generates an incremental zero-padded ID stored in .id_counter

COUNTER_FILE="${1:-.id_counter}"

# Create counter file if it doesn't exist
if [ ! -f "$COUNTER_FILE" ]; then
    echo "0" > "$COUNTER_FILE"
fi

# Read current value
CURRENT=$(cat "$COUNTER_FILE")

# Fallback if file is invalid
case "$CURRENT" in
    ''|*[!0-9]*) CURRENT=0 ;;
esac

# Increment
NEXT=$((CURRENT + 1))

# Save raw number
echo "$NEXT" > "$COUNTER_FILE"

# Print zero-padded ID (01, 02, ..., 99, 100, ...)
printf "%02d\n" "$NEXT"
