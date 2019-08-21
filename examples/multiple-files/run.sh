#!/bin/bash
pikolo render \
    --template examples/multiple-files/templates/t1 \
    --template examples/multiple-files/templates/t2 \
    --value examples/multiple-files/values/values.yaml \
    --value examples/multiple-files/values/values_2.json