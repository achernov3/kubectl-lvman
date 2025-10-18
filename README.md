# kubectl-lvman

A kubectl plugin for working with topolvm logicalvolumes in k8s.

## **Installation**

### Using `krew`

1. Install the [Krew](https://krew.sigs.k8s.io/docs/user-guide/setup/) plugin manager if you haven’t already.
2. Run the following command:

```bash
kubectl krew install lvman
```

## Examples

```shell
kubectl lvman show df data-opensearch-cluster-data-rumsk1-0 data-opensearch-cluster-data-rumsk2-0
                 PVC                                      PV                     STATUS       NODE                     VOLUME ID                CAPACITY  USAGE 
data-opensearch-cluster-data-rumsk1-0  pvc-03a879ce-6e15-47e2-b4c7-3190c9710475  Bound   afd1fea-log-058  57985969-bdd5-4c20-8e11-74aa757a1678  20Gi      850M  
data-opensearch-cluster-data-rumsk2-0  pvc-70aa4c17-be88-4168-9f6c-6c83c46d98e3  Bound   afd1fea-log-054  6b138144-d5b1-42fa-b289-3bd8f56d86a5  20Gi      1.7G  
```

```shell
kubectl lvman show orphan
             LOGICAL VOLUME                    NODE                     VOLUME ID               
pvc-02853c06-93fc-42fa-b37b-38ebb1e32a7f  afd1fea-w-053    650b5e08-e16f-48c5-ad3b-3b1037f7d04c 
pvc-0301da0c-052e-4f11-a6ed-9c675b58c7d9  afd1fea-w-053    ec30bf27-713c-4be1-bfe1-994bc98b1cd4 
pvc-048c5228-ccf5-4eaf-9df2-1dfabe5d629a  afd1fea-w-053    c2681c40-7da2-471d-8e92-40837730def6 
```

## TODO:
  - [x] bump urfave/cli to v3
  - [ ] add validate LV version