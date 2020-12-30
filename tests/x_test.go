package x

import (
	"fmt"
	"github.com/willf/bitset"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"
	"unsafe"

	"github.com/RoaringBitmap/roaring"
	fc "github.com/coocood/freecache"
	"github.com/pubgo/x/xfunc"
)

func init1() error {

	return xfunc.WithTimeout(time.Second, func() (err error) {
		defer func() {
			err = recover().(error)
		}()

		var a = &A{}
		runtime.SetFinalizer(a, func(a *A) {
			fmt.Println("SetFinalizer")
			panic("ok")
		})

		time.Sleep(time.Second * 5)
		fmt.Println("ss")
		return nil
	})
}

type A struct {
	A string
}

func TestName(t *testing.T) {
	err := init1()
	if err != nil {
		t.Log(err)
	}
	for range time.Tick(time.Second) {
		runtime.GC()
	}
}

func printMemStats() {
	//HeapSys：从操作系统获得的堆内存大小
	//HeapAlloc：堆上目前分配的内存
	//HeapIdle：堆上目前没有使用的内存
	//HeapReleased：回收到操作系统的内存
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("HeapAlloc = %v HeapIdel= %v HeapSys = %v  HeapReleased = %v\n", m.HeapAlloc/1024, m.HeapIdle/1024, m.HeapSys/1024, m.HeapReleased/1024)
}

var m1 = make(map[interface{}]interface{})

func BenchmarkName(b *testing.B) {
	var m2 = make(map[interface{}]interface{})
	printMemStats()

	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000*100; i++ {
			m1[i] = i
		}
		for k := range m1 {
			m2[k] = m1[k]
		}
	}

	m1 = m2
	runtime.GC()
	debug.FreeOSMemory()
	printMemStats()
}

func BenchmarkMap(b *testing.B) {
	var m1 = make(map[interface{}]interface{})
	//var m1 = sync.Map{}
	//var m2 = sync.Map{}
	printMemStats()

	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000*30; i++ {
			m1[strings.Repeat("#", i)] = i
		}
		printMemStats()
		for k := range m1 {
			delete(m1, k)
		}
		printMemStats()
	}
	//m1 = nil
	//m2 = m1
	runtime.GC()
	//debug.FreeOSMemory()
	printMemStats()
}

func BenchmarkSyncMap(b *testing.B) {
	//var m1 = make(map[interface{}]interface{})
	var m1 = sync.Map{}
	//var m2 = sync.Map{}
	printMemStats()

	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000*30; i++ {
			m1.Store(strings.Repeat("#", i), i)
		}
		printMemStats()
		m1.Range(func(key, value interface{}) bool {
			//m2.Store(key, value)
			m1.Delete(key)
			return true
		})
		printMemStats()
	}
	//m2 = m1
	//m1 = sync.Map{}
	runtime.GC()
	//debug.FreeOSMemory()
	printMemStats()
}

var ss = fc.NewCache(100 * 1024 * 1024)

func BenchmarkFreeCache(b *testing.B) {
	printMemStats()

	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000*100; i++ {
			ss.Set([]byte(strings.Repeat("#", i)), []byte(fmt.Sprintf("%d", i)), 0)
		}
		printMemStats()
		ss.Clear()
		printMemStats()
	}
	//m2 = m1
	//m1 = sync.Map{}
	//debug.FreeOSMemory()
	printMemStats()
}

/*
 */

func BenchmarkRWMutes(b *testing.B) {
	printMemStats()
	var rw sync.Mutex
	var g sync.WaitGroup
	for i := 0; i < b.N*100; i++ {
		g.Add(1)
		go func() {
			rw.Lock()
			time.Sleep(time.Millisecond)
			rw.Unlock()
			g.Done()
		}()
	}
	g.Wait()
	printMemStats()
}

func BenchmarkRWPin(b *testing.B) {
	printMemStats()
	var ss = sync.Map{}
	var g sync.WaitGroup
	for i := 0; i < b.N*100; i++ {
		g.Add(1)
		go func(i int) {
			if val, ok := ss.Load(i); ok {
				val.(*sync.Mutex).Lock()
				time.Sleep(time.Millisecond)
				val.(*sync.Mutex).Unlock()
			} else {
				ss.Store(i, &sync.Mutex{})
				time.Sleep(time.Millisecond)
			}
			g.Done()
		}(i)
	}
	g.Wait()
	printMemStats()
}

var ss1 = sync.Map{}
var sschan = make(chan []byte, 10000)

func BenchmarkChan(b *testing.B) {

	for i := 0; i < b.N; i++ {
		data := []byte(strings.Repeat("#", rand.Intn(i/100)))
		if len(sschan) == 0 {
			ss1.Store(data, data)
		} else {

		}

	}
}

type A1 struct {
	A string
	B int
	b string
}

func TestName2(t *testing.T) {
	var hdrBuf [100]byte
	hdr := (*A1)(unsafe.Pointer(&hdrBuf[0]))
	hdr.A = "11"
	hdr.b = "ddddd"
	fmt.Println((*A1)(unsafe.Pointer(&hdrBuf[0])).b)
	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.GC()
	fmt.Println(hdr, len(hdrBuf), cap(hdrBuf))
}

func BenchmarkBit(b *testing.B) {
	var rb1 = roaring.NewBitmap()
	printMemStats()
	for i := 0; i < b.N; i++ {
		rb1.AddInt(i)
		if !rb1.ContainsInt(i) {
			panic("")
		}
		rb1.Remove(uint32(i))
	}
	printMemStats()
}

var mm []interface{}

func BenchmarkBitset(b *testing.B) {
	printMemStats()
	b.ResetTimer()

	var b1 bitset.BitSet
	for i := 0; i < b.N; i++ {
		aa := rand.Intn(2 << 30)
		b1.Set(uint(aa))
		if !b1.Test(uint(aa)) {
			panic("")
		}
		//b1.Clear(uint(aa))
		//if b1.Test(uint(aa)) {
		//	panic("")
		//}
	}

	b.StopTimer()
	//runtime.GC()
	//debug.FreeOSMemory()
	printMemStats()
}

var sss = os.Stderr

func TestNamecc(t *testing.T) {

	f2, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}

	//int(f.Fd())

	//f.Write([]byte("ddd"))
	//os.Stderr = f
	//syscall.Write(2,[]byte("ddd"))
	//fmt.Println("ss1", f2.Fd())

	f, err := os.OpenFile("/dev/stderr", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	fmt.Println("ss", f.Fd())

	f.Write([]byte("ddd"))
	f.Write([]byte("ddd"))
	f.Write([]byte("ddd"))
	f.Write([]byte("ddd"))
	f.Write([]byte("ddd"))
	f.Write([]byte("ddd"))

	var ss []byte
	fmt.Println(ss[100])

	os.Stderr = sss
	fmt.Println("ss2", f2.Fd())

}

func TestNamew(t *testing.T) {
	var err error
	//os.Stdout, err = os.OpenFile("/dev/stdout", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	//if err != nil {
	//	panic(err)
	//}

	f1, err := os.OpenFile("./test1.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}

	f2, err := os.OpenFile("./test2.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}

	fmt.Println(os.Getpid())

	if er := syscall.Dup2(int(f1.Fd()), int(f2.Fd())); er != nil {
		panic(er)
	}

	f3 := os.NewFile(uintptr(f2.Fd()), "./test2.txt")
	f3.Write([]byte("hh"))
}

func TestName1(t *testing.T) {
	printMemStats()
	printMemStats()
}

func TestHttpClient(t *testing.T) {
	var g sync.WaitGroup
	g.Add(10000)
	var c http.Client
	for i := 0; i < 10000; i++ {
		go func() {
			req, _ := http.NewRequest("GET", "http://127.0.0.1:8080", nil)
			req.Close = true
			res, err := c.Do(req)
			if err != nil {
				panic(err)
			}

			io.Copy(ioutil.Discard, res.Body)
			_ = res.Body.Close()
			g.Done()
			//fmt.Println(res)
		}()

		//fmt.Println(shutil.Execute(fmt.Sprintf("gops %d| wc -l", os.Getpid())))
		//fmt.Println(os.Getpid())
		printMemStats()
		time.Sleep(time.Microsecond * 10)
	}
	g.Wait()
}
func TestNIl(t *testing.T) {
	var ni interface{}
	nn := *&ni

	if nn == nil {
		fmt.Println("nn is nil")
	}

	var nf func() error
	if nf == nil {
		fmt.Println("nf is nil")
	}

	mirror := func(i interface{}) interface{} { return i }

	if mirror(nf) == mirror(nn) {
		fmt.Println("nil is nil")
	}

	m := sync.Map{}
	k := "key"

	type aa struct {
	}

	m.Store(k, (*aa)(nil))

	val, ok := m.Load(k)
	if val == nil {
		fmt.Println("This line will not print !!!!!")
	}

	fmt.Println(reflect.ValueOf(val).IsValid())
	fmt.Println(reflect.ValueOf(val).IsNil())
	fmt.Println(reflect.ValueOf(val).IsZero())
	fmt.Println(reflect.ValueOf(val))
	fmt.Println(int(reflect.ValueOf(val).Pointer()))

	fmt.Println(reflect.TypeOf(val), val, ok)

	//v2, ok := val.(*aa)
	//fmt.Println(reflect.ValueOf(v2), ok)
	init11 := func(a interface{}) bool {
		fmt.Printf("%#v\n\n", a)
		return a == nil
	}
	fmt.Println((*aa)(nil) == val, (*aa)(nil) == nil, val == nil, init11(val))
}

func TestNil(t *testing.T) {
	fn := func(a interface{}) bool {
		return a == nil
	}
	t.Log(fn((*struct{})(nil)))
}
