滑动窗口常用来求子串，虽然有两个for循环，但算法时间复杂度为O(N)。因为每个元素只会进入窗口一次，然后被移出窗口一次，不会存在某些元素重复进入与移出窗口。





几种case：

（1）最小窗口固定，子串匹配问题；

（2）最小窗口不固定，略有复杂；



典型解题思路：

（1）滑动窗口

（2）滑动窗口+有序集合



