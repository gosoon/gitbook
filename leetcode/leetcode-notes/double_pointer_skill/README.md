### 双指针问题总结以及刷题技巧：

#### 只要数组有序，就应该想到双指针的技巧。

通常使用两个指针，一个在后一个在前，前面的指针用于获取最终结果所需要的元素，比如元素去重或者删除指针的元素时，前面的指针需要停留在非重复元素或者要保留的元素处，然后慢指针记录该元素的值，慢指针所在的最终位置通常为最终所需要数组的结尾处。
