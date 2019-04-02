```
const (
    _  = iota
    KB = 1 << (10 * iota) // 1<<(10 * 1)
    MB                    // 1<<(10 * 2)
    GB                    // 1<<(10 * 3)
    TB                    // 1<<(10 * 4)
)
```

```
// SyncPodType classifies pod updates, eg: create, update.
type SyncPodType int

const (
	// SyncPodSync is when the pod is synced to ensure desired state
	SyncPodSync SyncPodType = iota
	// SyncPodUpdate is when the pod is updated from source
	SyncPodUpdate
	// SyncPodCreate is when the pod is created from source
	SyncPodCreate
	// SyncPodKill is when the pod is killed based on a trigger internal to the kubelet for eviction.
	// If a SyncPodKill request is made to pod workers, the request is never dropped, and will always be processed.
	SyncPodKill
)
```
