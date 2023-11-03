#!/bin/bash
kubectl delete all --all -n msd-org-1
kubectl delete all --all -n msd-org-2
kubectl delete all --all -n msd-orderer
kubectl delete all --all -n monitoring
kubectl delete sa -n msd-orderer $(kubectl get sa -n msd-orderer | awk 'NR>1{print $1}')
kubectl delete sa -n msd-org-1 $(kubectl get sa -n msd-org-1 | awk 'NR>1{print $1}')
kubectl delete sa -n msd-org-2 $(kubectl get sa -n msd-org-2 | awk 'NR>1{print $1}')
kubectl delete cm -n msd-orderer $(kubectl get cm -n msd-orderer | awk 'NR>1{print $1}')
kubectl delete cm -n msd-org-1 $(kubectl get cm -n msd-org-1 | awk 'NR>1{print $1}')
kubectl delete cm -n msd-org-2 $(kubectl get cm -n msd-org-2 | awk 'NR>1{print $1}')
kubectl delete secret -n msd-orderer $(kubectl get secrets -n msd-orderer | awk 'NR>1{print $1}')
kubectl delete secret -n msd-org-1 $(kubectl get secrets -n msd-org-1 | awk 'NR>1{print $1}')
kubectl delete secret -n msd-org-2 $(kubectl get secrets -n msd-org-2 | awk 'NR>1{print $1}')
kubectl delete secret -n monitoring $(kubectl get secrets -n monitoring | awk 'NR>1{print $1}')
kubectl delete job -n msd-orderer msd-orderer1-hlf-k8s-hook-delete-secrets
kubectl delete job -n msd-org-1 msd-org-1-peer-1-hlf-k8s-hook-chaincode-delete-secrets 
kubectl delete job -n msd-org-1 msd-org-1-peer-2-hlf-k8s-hook-chaincode-delete-secrets
kubectl delete job -n msd-org-2 msd-org-2-peer-1-hlf-k8s-hook-chaincode-delete-secrets
kubectl delete job -n msd-org-2 msd-org-2-peer-2-hlf-k8s-hook-chaincode-delete-secrets
kubectl delete job -n msd-org-2 msd-org-2-hlf-k8s-hook-delete-secrets
kubectl delete job -n msd-org-1 msd-org-1-hlf-k8s-hook-delete-secrets
kubectl delete job -n msd-orderer msd-orderer-hlf-k8s-hook-delete-secrets

kubectl delete job -n msd-org-1 msd-org-1-peer-1-hlf-k8s-hook-delete-secrets
kubectl delete job -n msd-org-2 msd-org-2-peer-1-h
lf-k8s-hook-delete-secrets
kubectl delete job -n msd-org-1 msd-org-1-peer-2-hlf-k8s-hook-delete-secrets
kubectl delete job -n msd-org-2 msd-org-2-peer-2-hlf-k8s-hook-delete-secrets

kubectl delete pvc -n msd-orderer $(kubectl get pvc -n msd-orderer | awk 'NR>1 {print $1}')
kubectl delete pvc -n msd-org-1 $(kubectl get pvc -n msd-org-1 | awk 'NR>1 {print $1}')
kubectl delete pvc -n msd-org-2 $(kubectl get pvc -n msd-org-2 | awk 'NR>1 {print $1}')
