#!/bin/sh

[ -z $date ] && date=$(date +%F)

pandoc \
    --standalone \
    --from gfm --to man \
    --metadata=title:"ZET" \
    --metadata=section:"1" \
    --metadata=date:"$date" \
    --metadata=footer:"1" \
    --metadata=header:"DOCUMENTATION" \
    doc.md -o zet.1
