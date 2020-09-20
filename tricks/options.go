package tricks

import (
	"golang.org/x/net/context"
	"net"
	"time"
)

//MyDialer 是导出的接口类型
type MyDialer interface {
	DialContext(context.Context, net.Addr) (net.Conn, error)
}

//myDialer 是未导出的接口实现
type myDialer struct {
	dialer  *net.Dialer
	timeout time.Duration
	retry   int
}

func (d *myDialer) DialContext(ctx context.Context, addr net.Addr) (net.Conn, error) {
	return d.dialer.DialContext(ctx, addr.Network(), addr.String())
}

// 我们可以在构造方法中使用
func NewMyDialer(opts ...Option) MyDialer {
	// 首先我们将默认值填上
	d := &myDialer{
		timeout: time.Second,
		retry:   2,
	}
	// 接下来用传入的 Option 修改默认值，如果不需要修改默认值，
	// 就不需要传入对应的 Option
	for _, opt := range opts {
		opt.apply(d)
	}
	// 最后再检查一下，如果 Option 没有传入自定义的必要字段，我
	// 们在这里补一下。
	if d.dialer == nil {
		d.dialer = &net.Dialer{}
	}
	return d
}

// 我们也可以提供单独的方法，并随接口导出，提供类似 Set 模式的功能。
func (d *myDialer) ApplyOptions(opts ...Option) {
	for _, opt := range opts {
		opt.apply(d)
	}
}

//Option 单方法接口
type Option interface {
	apply(*myDialer)
}

//函数转单方法接口 Option
type optFunc func(*myDialer)

func (f optFunc) apply(d *myDialer) {
	f(d)
}
func RetryOption(r int) Option {
	return optFunc(func(d *myDialer) {
		d.retry = r
	})
}
func TimeoutOption(timeout time.Duration) Option {
	return optFunc(func(d *myDialer) {
		d.timeout = timeout
	})
}
func DialerOption(dialer *net.Dialer) Option {
	return optFunc(func(d *myDialer) {
		d.dialer = dialer
	})
}

func main() {
	// 无自定义 Option，全部使用默认
	//d0 := NewMyDialer()
	//// 只修改 Retry，并且 Retry 是0次
	//d1 := NewMyDialer(RetryOption(0))
	//// 修改多个 Option
	//d2 := NewMyDialer(RetryOption(5), TimeoutOption(time.Minute), DialerOption(&net.Dialer{
	//	KeepAlive: 3 * time.Second,
	//}))
	return
}
