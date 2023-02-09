```
// Evicted
for i in `kubectl get pod -o wide | grep -i Evicted | awk '{print $1}'`;do kubectl delete pod $i --force ;done

// Unknown
for i in `kubectl get pod -o wide | grep -i Unknown | awk '{print $1}'`;do kubectl delete pod $i --force ;done

// UnexpectedAdmissionError
for i in `kubectl get pod -o wide | grep -i UnexpectedAdmissionError | awk '{print $1}'`;do kubectl delete pod $i --force ;done

// Failed
for i in `kubectl get pod -o wide | grep -i Failed | awk '{print $1}'`;do kubectl delete pod $i --force ;done

// Error
for i in `kubectl get pod -o wide | grep -i Error | awk '{print $1}'`;do kubectl delete pod $i --force ;done
```

