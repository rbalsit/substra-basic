Hyperledger explorer is running on monitoring namespace so we need to copy below secrets from msd-org-1 namespace to monitoring namespace.

kubectl get secret ord-tls-rootcert  --namespace=msd-org-1 -o yaml | grep -v '^\s*namespace:\s' | kubectl apply --namespace=monitoring -f -
kubectl get secret hlf-msp-cert-admin   --namespace=msd-org-1 -o yaml | grep -v '^\s*namespace:\s' | kubectl apply --namespace=monitoring -f -
kubectl get secret hlf-msp-key-admin   --namespace=msd-org-1 -o yaml | grep -v '^\s*namespace:\s' | kubectl apply --namespace=monitoring -f -

kubectl apply -f monitoring/monitoring.yaml
kubectl apply -f monitoring/explorer.yaml
