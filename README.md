# kubectl-lvman

A kubectl plugin for managing TopoLVM logical volumes in Kubernetes clusters.

## Features

- **Disk usage monitoring** - View PVC disk usage via SSH on nodes
- **Orphan LV detection** - Find logical volumes without bound PVs
- **Prune orphan LVs** - Remove all orphan logical volumes at once
- **Full chain removal** - Delete PVC → PV → LogicalVolume chain

## Requirements

- Kubernetes cluster with TopoLVM installed
- Configured `kubectl` with cluster access
- SSH access to nodes with private key (default: `~/.ssh/id_rsa`)
- SSH user (default: `ops`)

## Installation

### Via krew (recommended)

```bash
kubectl krew install lvman
```

### From source

```bash
git clone https://github.com/yourrepo/kubectl-lvman.git
cd kubectl-lvman
make build
sudo mv kubectl-lvman /usr/local/bin/
kubectl lvman --help
```

## Usage

### Help

```bash
kubectl lvman --help
```

```
kubectl plugin for managing logical volumes in a kubernetes cluster with topolvm as storageClass

Usage:
  lvman [flags] [command]

Commands:
  show    shows disk usage by provided pvc names and oprhaned logical volumes in cluster
  prune   prune all oprhaned LV (which hasn't binded PV)
  remove  remove chain pvc -> pv -> lv by provided pvc
```

---

### 1. Show Disk Usage (`show df`)

Displays information about PVC, PV, node, and disk usage on the node via SSH.

```bash
kubectl lvman show df <pvc-name> [pvc-name ...]
```

**Example:**

```bash
kubectl lvman show df data-opensearch-cluster-data-rumsk1-0 data-opensearch-cluster-data-rumsk2-0
```

**Output:**

```
                  PVC                                      PV                     STATUS       NODE                     VOLUME ID                CAPACITY  USAGE 
data-opensearch-cluster-data-rumsk1-0  pvc-03a879ce-6e15-47e2-b4c7-3190c9710475  Bound   afd1fea-log-058  57985969-bdd5-4c20-8e11-74aa757a1678  20Gi      850M  
data-opensearch-cluster-data-rumsk2-0  pvc-70aa4c17-be88-4168-9f6c-6c83c46d98e3  Bound   afd1fea-log-054  6b138144-d5b1-42fa-b289-3bd8f56d86a5  20Gi      1.7G  
```

**Options:**

| Flag | Description | Default |
|------|-------------|---------|
| `--kubeconfig` | Path to kubeconfig | `~/.kube/config` |
| `--context` | kubectl context | current |
| `-n, --namespace` | Namespace | current |
| `--id_rsa` | Path to SSH private key | `~/.ssh/id_rsa` |
| `--username` | SSH username | `ops` |
| `--port` | SSH port | `22` |

---

### 2. Show Orphan Volumes (`show orphan`)

Displays TopoLVM logical volumes that exist in the cluster but have no bound PersistentVolume.

```bash
kubectl lvman show orphan
```

**Example:**

```bash
kubectl lvman show orphan
```

**Output:**

```
              LOGICAL VOLUME                    NODE                     VOLUME ID               
pvc-02853c06-93fc37b-38ebb1e32-42fa-ba7f  afd1fea-w-053    650b5e08-e16f-48c5-ad3b-3b1037f7d04c 
pvc-0301da0c-052e-4f11-a6ed-9c675b58c7d9  afd1fea-w-053    ec30bf27-713c-4be1-bfe1-994bc98b1cd4 
pvc-048c5228-ccf5-4eaf-9df2-1dfabe5d629a  afd1fea-w-053    c2681c40-7da2-471d-8e92-40837730def6 
```

If there are no orphan volumes:

```
There's no oprhaned logical volumes
```

**Options:**

| Flag | Description | Default |
|------|-------------|---------|
| `--kubeconfig` | Path to kubeconfig | `~/.kube/config` |
| `--context` | kubectl context | current |
| `-n, --namespace` | Namespace | current |

---

### 3. Prune Orphan LVs (`prune`)

Removes all TopoLVM logical volumes that have no bound PersistentVolume.

> **Warning:** This is a destructive operation. Make sure you understand which volumes will be deleted.

```bash
kubectl lvman prune
```

**Example:**

```bash
kubectl lvman prune
```

**Options:**

| Flag | Description | Default |
|------|-------------|---------|
| `--kubeconfig` | Path to kubeconfig | `~/.kube/config` |
| `--context` | kubectl context | current |
| `-n, --namespace` | Namespace | current |

---

### 4. Remove PVC Chain (`remove`)

Fully removes specified PVCs along with associated PV and LogicalVolume.

> **Warning:** This is a destructive operation. All data on the PVC will be lost.

```bash
kubectl lvman remove <pvc-name> [pvc-name ...]
```

**Example:**

```bash
kubectl lvman remove data-opensearch-cluster-data-rumsk1-0
```

**Options:**

| Flag | Description | Default |
|------|-------------|---------|
| `--kubeconfig` | Path to kubeconfig | `~/.kube/config` |
| `--context` | kubectl context | current |
| `-n, --namespace` | Namespace | current |

---

## Common Use Cases

### Cleanup after pod failure

1. Check for orphan volumes:
   ```bash
   kubectl lvman show orphan
   ```

2. Remove them:
   ```bash
   kubectl lvman prune
   ```

### Monitor disk usage

```bash
kubectl lvman show df -n production my-app-data-0 my-app-data-1
```

### Delete test PVC

```bash
kubectl lvman remove -n default test-pvc-001
```
