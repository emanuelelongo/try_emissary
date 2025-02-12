#!/usr/bin/bash
host=test.lcl:8080

for bucket in ops-2522-test-0-{0..4}
do
    if ! tmux has-session -t $bucket 2>/dev/null; then
        tmux new-session -d -s $bucket
    fi
    tmux send-keys -t $bucket "
    	warp mixed --host=$host \
        	--access-key=minioadmin \
         	--secret-key=minioadmin \
         	--objects=10 \
         	--obj.size=1MiB \
         	--concurrent=1 \
         	--duration=8s \
         	--bucket \"$bucket\" \
         	--header \"x-cbt-bucket:$bucket\"" C-m
done
