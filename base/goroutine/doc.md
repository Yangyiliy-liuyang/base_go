## 使用sync.Mutex
- goroutine使用时间过长，main函数已经早早结束了，会获得一个不完整的预期goroutine操作
- 预计到操作时间比较久，人为设置休眠


## 使用Channel（线程安全）
- 