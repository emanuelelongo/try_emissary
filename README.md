# Experimenting Ambassador Emissary

```bash
# install emissary
./install-emissary.sh

# install the example resources
helm install example example

# add test.lcl to /etc/hosts
echo "127.0.0.1\ttest.lcl" >> /etc/hosts

# port-forward emissary
kubectl port-forward -n emissary svc/emissary-ingress 8080:80

# run the test (using patched warp)
warp mixed --host=test.lcl:8080 \
    --access-key=minioadmin \
    --secret-key=minioadmin \
    --objects=10 \
    --obj.size=1MiB \
    --concurrent=1 \
    --duration=4s
    --header "x-cbt-bucket:bucket1"

# extract reports
unzstd *.zst

# take only errors and ops
find . -name "*.json" -exec sh -c 'cat {} | jq "{errors: .total.total_errors, ops: .total.throughput.segmented.median_ops}" > ${1%.json}_skim.json' sh {} \;

# compute error sum and ops average
jq -s 'reduce .[] as $item ({"errors": 0, "ops": 0, "count": 0}; .errors += $item.errors | .ops += $item.ops | .count += 1) | {errors: .errors, ops: (.ops / .count)}' *_skim.json
