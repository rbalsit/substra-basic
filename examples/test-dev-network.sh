#!/bin/bash

# (Local development script)
# This script verifies that the network is functional by invoking a smart
# contract on each node.
#
# Usage:
#   ./test-dev-network.sh N
#   where N is the number of nodes on your network
#
# Example:
#  ./test-dev-network.sh 2
#

# NUM_ORGS=${1:-2}

# for i in `seq $NUM_ORGS`; do
#     echo org-$i
#     kubectl exec -it -n org-$i `kubectl get pods -n org-$i | grep toolbox | cut -d' ' -f1` -- \
#     bash -c "peer chaincode invoke \
#         -C mychannel \
#         -n mycc \
#         --tls \
#         --clientauth \
#         --cafile /var/hyperledger/tls/ord/cert/cacert.pem \
#         --certfile /var/hyperledger/tls/server/pair/tls.crt \
#         --keyfile /var/hyperledger/tls/server/pair/tls.key \
#         -o network-orderer-hlf-ord.orderer:7050 \
#         -c '{\"Args\":[\"queryTraintuples\"]}'"
#     echo '-----------'
# done


NUM_ORGS=${1:-2}
NUM_PEERS=${2:-2}

for i in `seq $NUM_ORGS`; do
    echo msd-org-$i
    for k in `seq $NUM_PEERS`; do
        echo peer-${k}
        winpty kubectl exec -it -n org-$i `kubectl get pods -n org-$i | grep peer-${k}-hlf-k8s-toolbox | cut -d' ' -f1` -- \
        bash -c "peer chaincode invoke \
            -C mychannel \
            -n basic \
            --tls \
            --clientauth \
            --cafile /var/hyperledger/tls/ord/cert/cacert.pem \
            --certfile /var/hyperledger/tls/server/pair/tls.crt \
            --keyfile /var/hyperledger/tls/server/pair/tls.key \
            -o msd-orderer3-hlf-ord.msd-orderer.svc.cluster.local:7050 \
            -c '{\"Args\":[\"GetAllAssets\"]}'"
        echo '-----------'
    done
done
## Execute an invoke on 1 peer to demonstrate distributed data on network
r=$(openssl rand -hex 4)
echo peer-1 msd-org-1
winpty kubectl exec -it -n org-1 `kubectl get pods -n org-1 | grep peer-1-hlf-k8s-toolbox | cut -d' ' -f1` -- \
bash -c "peer chaincode invoke \
    -C mychannel \
    -n basic \
    --tls \
    --clientauth \
    --cafile /var/hyperledger/tls/ord/cert/cacert.pem \
    --certfile /var/hyperledger/tls/server/pair/tls.crt \
    --keyfile /var/hyperledger/tls/server/pair/tls.key \
    -o msd-orderer3-hlf-ord.msd-orderer.svc.cluster.local:7050 \
    -c '{\"Args\":[\"CreateAsset\", \"msd-${r}\", \"msdcolor\", \"13\", \"huy\",  \"1000000\"]}'"

echo "created asset msd-${r} on chain -- query again to see this"
echo '-----------'
#     -c '{\"Args\":[\"CreateAsset\", \"{\"id\": \"msd3\", \"color\": \"msdcolor\", \"size\": 13, \"owner\": \"huy\", \"appraisedValue\": 1000000}\"]}'"
#for i in `seq $NUM_ORGS`; do
#    echo msd-org-$i
#    for k in `seq $NUM_PEERS`; do
#        echo peer-${k}
#        winpty kubectl exec -it -n msd-org-$i `kubectl get pods -n msd-org-$i | grep peer-${k}-hlf-k8s-toolbox | cut -d' ' -f1` -- \
#        bash -c 'cat $FABRIC_CFG_PATH/core.yaml | grep "address:"'
#        echo '-----------'
#    done
#done
