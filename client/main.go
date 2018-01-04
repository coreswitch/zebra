package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/coreswitch/netutil"
	pb "github.com/coreswitch/zebra/proto"
	"google.golang.org/grpc"
)

type zebraClient struct {
	serv             pb.ZebraClient
	dispatCh         chan interface{}
	interfaceStream  pb.Zebra_InterfaceServiceClient
	routerIdStream   pb.Zebra_RouterIdServiceClient
	redistIPv4Stream pb.Zebra_RedistributeIPv4ServiceClient
	redistIPv6Stream pb.Zebra_RedistributeIPv6ServiceClient
	routeIPv4Stream  pb.Zebra_RouteIPv4ServiceClient
	routeIPv6Stream  pb.Zebra_RouteIPv6ServiceClient
	wg               *sync.WaitGroup
}

func NewZebraClient() *zebraClient {
	client := &zebraClient{
		wg:       &sync.WaitGroup{},
		dispatCh: make(chan interface{}, 4096),
	}
	return client
}

func (c *zebraClient) Stop() {
}

const (
	DEFAULT_VRF = 0
)

func (c *zebraClient) InterfaceSubscribe(vrfId uint32) error {
	stream, err := c.serv.InterfaceService(context.Background())
	if err != nil {
		return err
	}
	c.interfaceStream = stream

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			res, err := stream.Recv()
			if err != nil {
				return
			}
			c.dispatCh <- res
		}
	}()

	req := &pb.InterfaceRequest{
		Op:    pb.Op_InterfaceSubscribe,
		VrfId: vrfId,
	}
	err = stream.Send(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *zebraClient) RouterIdSubscribe(vrfId uint32) error {
	stream, err := c.serv.RouterIdService(context.Background())
	if err != nil {
		return err
	}
	c.routerIdStream = stream

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			res, err := stream.Recv()
			if err != nil {
				return
			}
			c.dispatCh <- res
		}
	}()

	req := &pb.RouterIdRequest{
		Op:    pb.Op_RouterIdSubscribe,
		VrfId: vrfId,
	}
	err = stream.Send(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *zebraClient) RouteIPv4Service() error {
	stream, err := c.serv.RouteIPv4Service(context.Background())
	if err != nil {
		return err
	}
	c.routeIPv4Stream = stream

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			res, err := stream.Recv()
			if err != nil {
				return
			}
			c.dispatCh <- res
		}
	}()

	return nil
}

func (c *zebraClient) RouteIPv4Add(r *pb.RouteIPv4) error {
	if c.routeIPv4Stream == nil {
		err := c.RouteIPv4Service()
		if err != nil {
			return err
		}
	}
	r.Op = pb.Op_RouteAdd
	return c.routeIPv4Stream.Send(r)
}

func (c *zebraClient) RouteIPv4Delete(r *pb.RouteIPv4) error {
	if c.routeIPv4Stream == nil {
		err := c.RouteIPv4Service()
		if err != nil {
			return err
		}
	}
	r.Op = pb.Op_RouteDelete
	return c.routeIPv4Stream.Send(r)
}

func (c *zebraClient) RouteIPv6Service() error {
	stream, err := c.serv.RouteIPv6Service(context.Background())
	if err != nil {
		return err
	}
	c.routeIPv6Stream = stream

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			res, err := stream.Recv()
			if err != nil {
				return
			}
			c.dispatCh <- res
		}
	}()

	return nil
}

func (c *zebraClient) RouteIPv6Add(r *pb.RouteIPv6) error {
	if c.routeIPv6Stream == nil {
		err := c.RouteIPv6Service()
		if err != nil {
			return err
		}
	}
	r.Op = pb.Op_RouteAdd
	return c.routeIPv6Stream.Send(r)
}

func (c *zebraClient) RouteIPv6Delete(r *pb.RouteIPv6) error {
	if c.routeIPv6Stream == nil {
		err := c.RouteIPv6Service()
		if err != nil {
			return err
		}
	}
	r.Op = pb.Op_RouteDelete
	return c.routeIPv6Stream.Send(r)
}

func main() {
	fmt.Println("goroutine", runtime.NumGoroutine())

	// Create Client.
	c := NewZebraClient()

	// Dial.
	conn, err := grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial fail", err)
	}
	fmt.Println("goroutine", runtime.NumGoroutine())

	// Get server client.
	c.serv = pb.NewZebraClient(conn)
	fmt.Println("goroutine", runtime.NumGoroutine())

	// Add services.
	err = c.InterfaceSubscribe(DEFAULT_VRF)
	if err != nil {
		c.Stop()
	}
	err = c.RouterIdSubscribe(DEFAULT_VRF)
	if err != nil {
		c.Stop()
	}

	// Dispatch function.
	done := make(chan interface{})
	go func() {
		select {
		case res := <-c.dispatCh:
			switch res.(type) {
			case *pb.InterfaceUpdate:
				fmt.Println("IfUpdate res is processing")
			case *pb.RouterIdUpdate:
				fmt.Println("RouterId res is processing")
			case *pb.RouteIPv4:
				fmt.Println("")
			case *pb.RouteIPv6:
				fmt.Println("")
			}
		case <-done:
			return
		}
	}()

	fmt.Println("goroutine", runtime.NumGoroutine())

	fmt.Println("-- sleep start --")
	time.Sleep(time.Second * 3)
	fmt.Println("-- sleep end --")

	// Route Add.
	p, _ := netutil.ParsePrefix("10.0.0.0/24")
	r := &pb.RouteIPv4{
		Type: pb.RIB_BGP,
		Prefix: &pb.PrefixIPv4{
			Addr:   p.IP,
			Length: uint32(p.Length),
		},
	}
	c.RouteIPv4Add(r)

	p6, _ := netutil.ParsePrefix("::1/128")
	r6 := &pb.RouteIPv6{
		Type: pb.RIB_BGP,
		Prefix: &pb.PrefixIPv6{
			Addr:   p6.IP,
			Length: uint32(p6.Length),
		},
	}
	c.RouteIPv6Add(r6)

	// Close interafce stream -> Invoke all client EOF.
	c.interfaceStream.CloseSend()

	for {
		fmt.Println("goroutine", runtime.NumGoroutine())
		time.Sleep(time.Second)
	}

	// stream2.CloseSend()
	//conn.Close()
	//cancel()
	//	fmt.Println("CloseSend err", err)

	// for {
	// 	fmt.Println("goroutine", runtime.NumGoroutine())
	// 	time.Sleep(time.Second)
	// }

	fmt.Println("Before wait group")
	c.wg.Wait()
	close(done)
	fmt.Println("After wait group")

	for {
		fmt.Println("goroutine", runtime.NumGoroutine())
		time.Sleep(time.Second)
	}
}
