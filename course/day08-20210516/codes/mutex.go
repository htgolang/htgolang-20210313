package main

func main() {
	// Mutex 互斥锁
	// RWMutex 读写锁
	// obj => 写write: g1 write g2 write 互锁锁
	// obj => 读write read: g1 write g2 read 互斥锁
	// obj => 读 read: g1 read g2 read 读写锁
}
