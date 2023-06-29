#!/bin/bash
kubectl delete all --all -n org-1
kubectl delete all --all -n org-2
kubectl delete all --all -n orderer
kubectl delete sa -n orderer $(kubectl get sa -n orderer | awk 'NR>1{print $1}')
kubectl delete sa -n org-1 $(kubectl get sa -n org-1 | awk 'NR>1{print $1}')
kubectl delete sa -n org-2 $(kubectl get sa -n org-2 | awk 'NR>1{print $1}')
kubectl delete cm -n orderer $(kubectl get cm -n orderer | awk 'NR>1{print $1}')
kubectl delete cm -n org-1 $(kubectl get cm -n org-1 | awk 'NR>1{print $1}')
kubectl delete cm -n org-2 $(kubectl get cm -n org-2 | awk 'NR>1{print $1}')
kubectl delete secret -n orderer $(kubectl get secrets -n orderer | awk 'NR>1{print $1}')
kubectl delete secret -n org-1 $(kubectl get secrets -n org-1 | awk 'NR>1{print $1}')
kubectl delete secret -n org-2 $(kubectl get secrets -n org-2 | awk 'NR>1{print $1}')
kubectl delete job -n orderer orderer1-hlf-k8s-hook-delete-secrets
kubectl delete job -n org-1 org-1-peer-1-hlf-k8s-hook-chaincode-delete-secrets 
kubectl delete job -n org-1 org-1-peer-2-hlf-k8s-hook-chaincode-delete-secrets
kubectl delete job -n org-2 org-2-peer-1-hlf-k8s-hook-chaincode-delete-secrets
kubectl delete job -n org-2 org-2-peer-2-hlf-k8s-hook-chaincode-delete-secrets
kubectl delete job -n org-2 org-2-hlf-k8s-hook-delete-secrets
kubectl delete job -n org-1 org-1-hlf-k8s-hook-delete-secrets
kubectl delete job -n orderer orderer-hlf-k8s-hook-delete-secrets

kubectl delete job -n org-1 org-1-peer-1-hlf-k8s-hook-delete-secrets
kubectl delete job -n org-2 org-2-peer-1-h
lf-k8s-hook-delete-secrets
kubectl delete job -n org-1 org-1-peer-2-hlf-k8s-hook-delete-secrets
kubectl delete job -n org-2 org-2-peer-2-hlf-k8s-hook-delete-secrets

kubectl delete pvc -n orderer $(kubectl get pvc -n orderer | awk 'NR>1 {print $1}')
kubectl delete pvc -n org-1 $(kubectl get pvc -n org-1 | awk 'NR>1 {print $1}')
kubectl delete pvc -n org-2 $(kubectl get pvc -n org-2 | awk 'NR>1 {print $1}')
