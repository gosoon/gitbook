

````
func (jm *Controller) deleteActivePods(job *batch.Job, pods []*v1.Pod) (int32, error) {
    errCh := make(chan error, len(pods))
    successfulDeletes := int32(len(pods))
    wg := sync.WaitGroup{}
    wg.Add(len(pods))
    for i := range pods {
        go func(pod *v1.Pod) {
            defer wg.Done()
            if err := jm.podControl.DeletePod(job.Namespace, pod.Name, job); err != nil && !apierrors.IsNotFound(err) {
                atomic.AddInt32(&successfulDeletes, -1)
                errCh <- err
                utilruntime.HandleError(err)
            }
        }(pods[i])
    }
    wg.Wait()
    return successfulDeletes, errorFromChannel(errCh)
}
```

error handle:
```
func errorFromChannel(errCh <-chan error) error {
    select {
    case err := <-errCh:
        return err
    default:
    }
    return nil
}
```

used:
```
func ConvertErrsToErr(errs []error) error {
    if len(errs) == 0 {
        return nil
    }

    var str string
    for _, err := range errs {
        str += err.Error() + "\n "
    }
    return errors.New(str)
}
```
