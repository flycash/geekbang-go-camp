package concurrency

import "sync"

// PublicResource 你永远不知道你的用户拿了它会干啥
// 他即便不用 PublicResourceLock 你也毫无办法
// 如果你用这个resource，一定要用锁
var PublicResource interface{}
var PublicResourceLock sync.Mutex


// privateResource 要好一点，祈祷你的同事会来看你的注释，知道要用锁
// 很多库都是这么写的，我也写了很多类似的代码=。=
var privateResource interface{}
var privateResourceLock sync.Mutex

// safeResource 很棒，所有的期望对资源的操作都只能通过定义在上 safeResource 上的方法来进行
type safeResource struct {
	resource interface{}
	lock sync.Mutex
}

func (s *safeResource) DoSomethingToResource()  {
	s.lock.Lock()
	defer s.lock.Unlock()
}

// Registry 没有用锁，并不安全
type Registry struct {
	resources map[string]interface{}
}

func (r *Registry) Register(name string, resource interface{})  {
	r.resources[name] = resource
}

func (r *Registry) Get(name string) (interface{}, error) {
	return nil, nil
}