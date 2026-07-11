#!/bin/sh

input=$(cat)

basename=$(printf '%s' "$input" | jq -r '.entry_info.basename')

id=${basename%%_*}

if [ "$basename" = "$id" ]; then
    slug=""
else
    slug=${basename#*_}
fi

jq -n \
    --arg id "$id" \
    --arg slug "$slug" \
    '{
        metas: {
            id: $id,
            slug: $slug
        }
    }'
