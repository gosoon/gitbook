

kubectl  的所有操作都是调用 kube-apisever 的 API 实现的，所以其子命令都有相应的 API，每次在调用 kubectl 时使用参数  -v=9  可以看调用的相关 API，例：
 
`$ kubectl get node -v=9` 

以下为 kubernetes 开发中常用的 API：

### Deployment

<style type="text/css">
.tg  {border-collapse:collapse;border-spacing:0;border-color:#ccc;}
.tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#fff;}
.tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#f0f0f0;}
.tg .tg-mx66{font-size:xx-small;border-color:inherit;text-align:left}
.tg .tg-tn4b{background-color:#f9f9f9;font-size:xx-small;text-align:left;vertical-align:top}
.tg .tg-kd1j{font-size:xx-small;border-color:inherit;text-align:left;vertical-align:top}
.tg .tg-hvas{font-size:xx-small;text-align:left;vertical-align:top}
</style>
<table class="tg">
  <tr>
    <th class="tg-mx66">操作</th>
    <th class="tg-mx66">api</th>
    <th class="tg-mx66">url参数</th>
    <th class="tg-kd1j">body</th>
    <th class="tg-kd1j">header</th>
    <th class="tg-kd1j">method</th>
  </tr>
  <tr>
    <td class="tg-tn4b">ListDeployments</td>
    <td class="tg-tn4b">/apis/apps/v1beta1/namespaces/<br>{namespace}/deployments</td>
    <td class="tg-tn4b">labelSelector</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">Get</td>
  </tr>
  <tr>
    <td class="tg-hvas">GetDeployment</td>
    <td class="tg-hvas">/apis/apps/v1beta1/namespaces/<br>{namespace}/deployments/{name}</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">Get</td>
  </tr>
  <tr>
    <td class="tg-tn4b">CreateDeployment</td>
    <td class="tg-tn4b">/apis/apps/v1beta1/namespaces/<br>{namespace}/deployments</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">v1.Deployment</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">Post</td>
  </tr>
  <tr>
    <td class="tg-hvas">UpdateDeployment</td>
    <td class="tg-hvas">/apis/apps/v1beta1/namespaces/<br>{namespace}/deployments/{name}</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">v1.Deployment</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">Put</td>
  </tr>
  <tr>
    <td class="tg-tn4b">PatchDeployment</td>
    <td class="tg-tn4b">/apis/apps/v1beta1/namespaces/<br>{namespace}/deployments/{name}</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">v1.Deployment</td>
    <td class="tg-tn4b">application/<br>json-<br>patch+<br>json</td>
    <td class="tg-tn4b">Patch</td>
  </tr>
  <tr>
    <td class="tg-hvas">DeleteDeployment</td>
    <td class="tg-hvas">/apis/apps/v1beta1/namespaces/<br>{namespace}/deployments/{name}</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">Delete</td>
  </tr>
  <tr>
    <td class="tg-tn4b">GetDeploymentScale</td>
    <td class="tg-tn4b">/apis/extensions/v1beta1/namespaces/<br>{namespace}/deployments/{name}/scale</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">Get</td>
  </tr>
  <tr>
    <td class="tg-hvas">PatchDeploymentScale</td>
    <td class="tg-hvas">/apis/extensions/v1beta1/namespaces/<br>{namespace}/deployments/{name}/scale</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">v1.Scale</td>
    <td class="tg-hvas">application/<br>strategic<br>-merge-patch<br>+json</td>
    <td class="tg-hvas">Patch</td>
  </tr>
  <tr>
    <td class="tg-tn4b">DeleteDeployment</td>
    <td class="tg-tn4b">/apis/apps/v1beta1/namespaces/<br>{namespace}/deployments</td>
    <td class="tg-tn4b">labelSelector</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">Delete</td>
  </tr>
  <tr>
    <td class="tg-hvas">DeleteReplicaSet</td>
    <td class="tg-hvas">/apis/extensions/v1beta1/namespaces/<br>{namespace}/replicasets</td>
    <td class="tg-hvas">labelSelector</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">Delete</td>
  </tr>
</table>

### Pod

<style type="text/css">
.tg  {border-collapse:collapse;border-spacing:0;border-color:#ccc;}
.tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#fff;}
.tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#f0f0f0;}
.tg .tg-bstp{background-color:#f9f9f9;font-size:xx-small;border-color:inherit;text-align:left}
.tg .tg-mx66{font-size:xx-small;border-color:inherit;text-align:left}
.tg .tg-tn4b{background-color:#f9f9f9;font-size:xx-small;text-align:left;vertical-align:top}
.tg .tg-r462{background-color:#f9f9f9;font-size:xx-small;border-color:inherit;text-align:left;vertical-align:top}
.tg .tg-kd1j{font-size:xx-small;border-color:inherit;text-align:left;vertical-align:top}
</style>
<table class="tg">
  <tr>
    <th class="tg-mx66">操作</th>
    <th class="tg-mx66">api</th>
    <th class="tg-mx66">url参数</th>
    <th class="tg-kd1j">body</th>
    <th class="tg-kd1j">header</th>
    <th class="tg-kd1j">method</th>
  </tr>
  <tr>
    <td class="tg-bstp">ListPods</td>
    <td class="tg-bstp">/api/v1/namespaces/{namespace}/pods<br>/api/v1/pods</td>
    <td class="tg-bstp">labelSelector</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">Get</td>
  </tr>
  <tr>
    <td class="tg-mx66">GetPod</td>
    <td class="tg-mx66">/api/v1/namespaces/{namespace}/pods/{name}</td>
    <td class="tg-mx66">-</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">Get</td>
  </tr>
  <tr>
    <td class="tg-r462">CreatePod</td>
    <td class="tg-r462">/api/v1/namespaces/{namespace}/pods</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">v1.Pod</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">Post</td>
  </tr>
  <tr>
    <td class="tg-kd1j">UpdatePod</td>
    <td class="tg-kd1j">/api/v1/namespaces/{namespace}/pods/{name}</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">v1.Pod</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">Put</td>
  </tr>
  <tr>
    <td class="tg-r462">DeletePod</td>
    <td class="tg-r462">/api/v1/namespaces/{namespace}/pods/{name}</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">Delete</td>
  </tr>
  <tr>
    <td class="tg-kd1j">WatchPodsAll</td>
    <td class="tg-kd1j">/api/v1/watch/namespaces/{namespace}/pods<br>/api/v1/watch/pods</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">"Accept",<br>"application/json, <br>*/*"</td>
    <td class="tg-kd1j">Get<br></td>
  </tr>
  <tr>
    <td class="tg-tn4b">GetPodLog</td>
    <td class="tg-tn4b">/api/v1/namespaces/{namespace}/pods/{name}/log</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">Get</td>
  </tr>
</table>

### PV

<style type="text/css">
.tg  {border-collapse:collapse;border-spacing:0;border-color:#ccc;}
.tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#fff;}
.tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#f0f0f0;}
.tg .tg-mx66{font-size:xx-small;border-color:inherit;text-align:left}
.tg .tg-tn4b{background-color:#f9f9f9;font-size:xx-small;text-align:left;vertical-align:top}
.tg .tg-r462{background-color:#f9f9f9;font-size:xx-small;border-color:inherit;text-align:left;vertical-align:top}
.tg .tg-kd1j{font-size:xx-small;border-color:inherit;text-align:left;vertical-align:top}
.tg .tg-hvas{font-size:xx-small;text-align:left;vertical-align:top}
</style>
<table class="tg">
  <tr>
    <th class="tg-mx66">操作</th>
    <th class="tg-mx66">api</th>
    <th class="tg-mx66">url参数</th>
    <th class="tg-kd1j">body</th>
    <th class="tg-kd1j">header</th>
    <th class="tg-kd1j">method</th>
  </tr>
  <tr>
    <td class="tg-r462">CreatePVC</td>
    <td class="tg-r462">/api/v1/namespaces/{namespace}/persistentvolumeclaims</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">v1.Persistent<br>VolumeClaim</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">Post</td>
  </tr>
  <tr>
    <td class="tg-kd1j">GetPVC</td>
    <td class="tg-kd1j">/api/v1/namespaces/{namespace}/persistentvolumeclaims/{pvc}</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">Get</td>
  </tr>
  <tr>
    <td class="tg-tn4b">DeletePVC</td>
    <td class="tg-tn4b">/api/v1/namespaces/{namespace}/persistentvolumeclaims/{pvc}</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">Delete</td>
  </tr>
  <tr>
    <td class="tg-hvas">DeletePV</td>
    <td class="tg-hvas">/api/v1/persistentvolumes/{pv}</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">Delete</td>
  </tr>
</table>

### Service

<style type="text/css">
.tg  {border-collapse:collapse;border-spacing:0;border-color:#ccc;}
.tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#fff;}
.tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#f0f0f0;}
.tg .tg-mx66{font-size:xx-small;border-color:inherit;text-align:left}
.tg .tg-r462{background-color:#f9f9f9;font-size:xx-small;border-color:inherit;text-align:left;vertical-align:top}
.tg .tg-kd1j{font-size:xx-small;border-color:inherit;text-align:left;vertical-align:top}
</style>
<table class="tg">
  <tr>
    <th class="tg-mx66">操作</th>
    <th class="tg-mx66">api</th>
    <th class="tg-mx66">url参数</th>
    <th class="tg-kd1j">body</th>
    <th class="tg-kd1j">header</th>
    <th class="tg-kd1j">method</th>
  </tr>
  <tr>
    <td class="tg-r462">ListServices</td>
    <td class="tg-r462">/api/v1/namespaces/{namespace}/services</td>
    <td class="tg-r462">labelSelector</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">Get</td>
  </tr>
  <tr>
    <td class="tg-kd1j">GetService</td>
    <td class="tg-kd1j">/api/v1/namespaces/{namespace}/services/{service}</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">Get</td>
  </tr>
  <tr>
    <td class="tg-r462">WatchServiceAll</td>
    <td class="tg-r462">/api/v1/watch/namespaces/{namespace}/services</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">Get</td>
  </tr>
  <tr>
    <td class="tg-kd1j">WatchService</td>
    <td class="tg-kd1j">/api/v1/watch/namespaces/{namespace}/services/{service}</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">Get</td>
  </tr>
  <tr>
    <td class="tg-r462">CreateService</td>
    <td class="tg-r462">/api/v1/namespaces/{namespace}/services</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">v1.Service</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">Post</td>
  </tr>
  <tr>
    <td class="tg-kd1j">UpdateService</td>
    <td class="tg-kd1j">/api/v1/namespaces/{namespace}/services/{service}</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">v1.Service</td>
    <td class="tg-kd1j">-</td>
    <td class="tg-kd1j">Put</td>
  </tr>
  <tr>
    <td class="tg-r462">DeleteService</td>
    <td class="tg-r462">/api/v1/namespaces/{namespace}/services/{service}</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">-</td>
    <td class="tg-r462">Delete</td>
  </tr>
</table>


### Statefulset

<style type="text/css">
.tg  {border-collapse:collapse;border-spacing:0;border-color:#ccc;}
.tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#fff;}
.tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#f0f0f0;}
.tg .tg-mx66{font-size:xx-small;border-color:inherit;text-align:left}
.tg .tg-tn4b{background-color:#f9f9f9;font-size:xx-small;text-align:left;vertical-align:top}
.tg .tg-kd1j{font-size:xx-small;border-color:inherit;text-align:left;vertical-align:top}
.tg .tg-hvas{font-size:xx-small;text-align:left;vertical-align:top}
</style>
<table class="tg">
  <tr>
    <th class="tg-mx66">操作</th>
    <th class="tg-mx66">api</th>
    <th class="tg-mx66">url参数</th>
    <th class="tg-kd1j">body</th>
    <th class="tg-kd1j">header</th>
    <th class="tg-kd1j">method</th>
  </tr>
  <tr>
    <td class="tg-tn4b">ListStatefulSet</td>
    <td class="tg-tn4b">/apis/apps/v1beta1/namespaces/<br>{namespace}/statefulsets</td>
    <td class="tg-tn4b">labelSelector</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">Get</td>
  </tr>
  <tr>
    <td class="tg-hvas">GetStatefulSet</td>
    <td class="tg-hvas">/apis/apps/v1beta1/namespaces/<br>{namespace}/statefulsets/{name}</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">Get</td>
  </tr>
  <tr>
    <td class="tg-tn4b">CreateStatefulSet</td>
    <td class="tg-tn4b">/apis/apps/v1beta1/namespaces/<br>{namespace}/statefulsets</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">v1.StatefulSet</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">Post</td>
  </tr>
  <tr>
    <td class="tg-hvas">UpdateStatefulSet</td>
    <td class="tg-hvas">/apis/apps/v1beta1/namespaces/<br>{namespace}/statefulsets/{name}</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">v1.StatefulSet</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">Put</td>
  </tr>
  <tr>
    <td class="tg-tn4b">PatchStatefulSet</td>
    <td class="tg-tn4b">/apis/apps/v1beta1/namespaces/<br>{namespace}/statefulsets/{name}</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">v1.StatefulSet</td>
    <td class="tg-tn4b">application/<br>json-patch+json</td>
    <td class="tg-tn4b">Patch</td>
  </tr>
  <tr>
    <td class="tg-hvas">PatchStatefulSetScale</td>
    <td class="tg-hvas">/apis/extensions/v1beta1/namespaces/<br>{namespace}/statefulsets/{name}/scale</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">v1.Scale</td>
    <td class="tg-hvas">application/<br>strategic-merge<br>-patch+json</td>
    <td class="tg-hvas">Patch</td>
  </tr>
  <tr>
    <td class="tg-tn4b">DeleteStatefulSet</td>
    <td class="tg-tn4b">/apis/apps/v1beta1/namespaces/<br>{namespace}/statefulsets/{name}</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">Delete</td>
  </tr>
</table>


###  StorageClass

<style type="text/css">
.tg  {border-collapse:collapse;border-spacing:0;border-color:#ccc;}
.tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#fff;}
.tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#f0f0f0;}
.tg .tg-mx66{font-size:xx-small;border-color:inherit;text-align:left}
.tg .tg-buh4{background-color:#f9f9f9;text-align:left;vertical-align:top}
.tg .tg-kd1j{font-size:xx-small;border-color:inherit;text-align:left;vertical-align:top}
</style>
<table class="tg">
  <tr>
    <th class="tg-mx66">操作</th>
    <th class="tg-mx66">api</th>
    <th class="tg-mx66">url参数</th>
    <th class="tg-kd1j">body</th>
    <th class="tg-kd1j">header</th>
    <th class="tg-kd1j">method</th>
  </tr>
  <tr>
    <td class="tg-buh4">ListStorageClass</td>
    <td class="tg-buh4">/apis/storage.k8s.io/v1/storageclasses</td>
    <td class="tg-buh4">-</td>
    <td class="tg-buh4">-</td>
    <td class="tg-buh4">-</td>
    <td class="tg-buh4">Get</td>
  </tr>
</table>

### Nodes

<style type="text/css">
.tg  {border-collapse:collapse;border-spacing:0;border-color:#ccc;}
.tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#fff;}
.tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#f0f0f0;}
.tg .tg-sg5v{font-size:100%;border-color:inherit;text-align:left;vertical-align:top}
.tg .tg-0w8i{font-size:100%;text-align:left;vertical-align:top}
.tg .tg-nlf0{font-size:100%;border-color:inherit;text-align:left}
.tg .tg-87db{background-color:#f9f9f9;font-size:100%;text-align:left;vertical-align:top}
</style>
<table class="tg">
  <tr>
    <th class="tg-nlf0">操作</th>
    <th class="tg-nlf0">api</th>
    <th class="tg-nlf0">url参数</th>
    <th class="tg-sg5v">body</th>
    <th class="tg-sg5v">header</th>
    <th class="tg-sg5v">method</th>
  </tr>
  <tr>
    <td class="tg-87db">ListNodes</td>
    <td class="tg-87db">/api/v1/nodes</td>
    <td class="tg-87db">labelSelector</td>
    <td class="tg-87db">-</td>
    <td class="tg-87db">-</td>
    <td class="tg-87db">Get</td>
  </tr>
  <tr>
    <td class="tg-0w8i">GetNode</td>
    <td class="tg-0w8i">/api/v1/nodes/{name}</td>
    <td class="tg-0w8i">-</td>
    <td class="tg-0w8i">-</td>
    <td class="tg-0w8i">-</td>
    <td class="tg-0w8i">GetNode</td>
  </tr>
  <tr>
    <td class="tg-87db">CreateNode</td>
    <td class="tg-87db">api/v1/nodes</td>
    <td class="tg-87db">-</td>
    <td class="tg-87db">v1.Node</td>
    <td class="tg-87db">-</td>
    <td class="tg-87db">Post</td>
  </tr>
  <tr>
    <td class="tg-0w8i">UpdateNode</td>
    <td class="tg-0w8i">/api/v1/nodes/{name}</td>
    <td class="tg-0w8i">-</td>
    <td class="tg-0w8i">v1.Node</td>
    <td class="tg-0w8i">-</td>
    <td class="tg-0w8i">Put</td>
  </tr>
  <tr>
    <td class="tg-87db">DeleteNode</td>
    <td class="tg-87db">/api/v1/nodes/{name}</td>
    <td class="tg-87db">-</td>
    <td class="tg-87db">-</td>
    <td class="tg-87db">-</td>
    <td class="tg-87db">Delete</td>
  </tr>
</table>


### Endpoints

<style type="text/css">
.tg  {border-collapse:collapse;border-spacing:0;border-color:#ccc;}
.tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#fff;}
.tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#f0f0f0;}
.tg .tg-mx66{font-size:xx-small;border-color:inherit;text-align:left}
.tg .tg-tn4b{background-color:#f9f9f9;font-size:xx-small;text-align:left;vertical-align:top}
.tg .tg-kd1j{font-size:xx-small;border-color:inherit;text-align:left;vertical-align:top}
.tg .tg-hvas{font-size:xx-small;text-align:left;vertical-align:top}
</style>
<table class="tg">
  <tr>
    <th class="tg-mx66">操作</th>
    <th class="tg-mx66">api</th>
    <th class="tg-mx66">url参数</th>
    <th class="tg-kd1j">body</th>
    <th class="tg-kd1j">header</th>
    <th class="tg-kd1j">method</th>
  </tr>
  <tr>
    <td class="tg-tn4b">ListEndpoints</td>
    <td class="tg-tn4b">/api/v1/namespaces/{namespace}/endpoints<br>/api/v1/endpoints</td>
    <td class="tg-tn4b">labelSelector</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">Get</td>
  </tr>
  <tr>
    <td class="tg-hvas">GetEndpoints</td>
    <td class="tg-hvas">/api/v1/namespaces/{namespace}/<br>endpoints/{service}</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">Get</td>
  </tr>
  <tr>
    <td class="tg-tn4b">ListEvents</td>
    <td class="tg-tn4b">/api/v1/namespaces/{namespace}/events<br>/api/v1/events</td>
    <td class="tg-tn4b">fieldSelector</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">Get</td>
  </tr>
  <tr>
    <td class="tg-tn4b">WatchEndpointsAll</td>
    <td class="tg-tn4b">/api/v1/watch/namespaces/{namespace}/endpoints<br>/api/v1/watch/endpoints</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">Get</td>
  </tr>
  <tr>
    <td class="tg-hvas">WatchEndpoints</td>
    <td class="tg-hvas">/api/v1/watch/namespaces/{namespace}/<br>endpoints/{service}</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">Accept, <br>application/json<br>, */*</td>
    <td class="tg-hvas">Get</td>
  </tr>
</table>

### Namespace

<style type="text/css">
.tg  {border-collapse:collapse;border-spacing:0;border-color:#ccc;}
.tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#fff;}
.tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#f0f0f0;}
.tg .tg-mx66{font-size:xx-small;border-color:inherit;text-align:left}
.tg .tg-tn4b{background-color:#f9f9f9;font-size:xx-small;text-align:left;vertical-align:top}
.tg .tg-kd1j{font-size:xx-small;border-color:inherit;text-align:left;vertical-align:top}
.tg .tg-hvas{font-size:xx-small;text-align:left;vertical-align:top}
</style>
<table class="tg">
  <tr>
    <th class="tg-mx66">操作</th>
    <th class="tg-mx66">api</th>
    <th class="tg-mx66">url参数</th>
    <th class="tg-kd1j">body</th>
    <th class="tg-kd1j">header</th>
    <th class="tg-kd1j">method</th>
  </tr>
  <tr>
    <td class="tg-hvas">ListNamespaces</td>
    <td class="tg-hvas">/api/v1/namespaces</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">-</td>
    <td class="tg-hvas">Get</td>
  </tr>
  <tr>
    <td class="tg-tn4b">WatchNamespaces</td>
    <td class="tg-tn4b">/api/v1/watch/namespaces</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">-</td>
    <td class="tg-tn4b">Accept, <br>application/json<br>, */*</td>
    <td class="tg-tn4b">Get</td>
  </tr>
</table>


