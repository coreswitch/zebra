package main

import (
	"context"
	"fmt"
	"net"
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
				fmt.Println("XXX interface stream.Recv error", err)
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

	// req = &pb.InterfaceRequest{
	// 	Op:    pb.Op_InterfaceUnsubscribe,
	// 	VrfId: vrfId,
	// }
	// err = stream.Send(req)
	// if err != nil {
	// 	return err
	// }

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

func (c *zebraClient) RouteIPv4Add(r *pb.Route) error {
	if c.routeIPv4Stream == nil {
		err := c.RouteIPv4Service()
		if err != nil {
			return err
		}
	}
	r.Op = pb.Op_RouteAdd
	return c.routeIPv4Stream.Send(r)
}

func (c *zebraClient) RouteIPv4Delete(r *pb.Route) error {
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
	fmt.Println("XXX IPv6 service", stream)

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

func (c *zebraClient) RouteIPv6Add(r *pb.Route) error {
	if c.routeIPv6Stream == nil {
		err := c.RouteIPv6Service()
		if err != nil {
			return err
		}
	}
	r.Op = pb.Op_RouteAdd
	return c.routeIPv6Stream.Send(r)
}

func (c *zebraClient) RouteIPv6Delete(r *pb.Route) error {
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

	// Dispatch function.
	done := make(chan interface{})
	go func() {
		for {
			select {
			case res := <-c.dispatCh:
				switch res.(type) {
				case *pb.InterfaceUpdate:
					mes := res.(*pb.InterfaceUpdate)
					fmt.Println("IfUpdate:", mes.Op, mes.Name, mes.Index, mes.Metric, mes.Mtu)
					for _, addr := range mes.AddrIpv4 {
						p := &netutil.Prefix{}
						p.IP = addr.Addr.Addr
						p.Length = int(addr.Addr.Length)
						fmt.Println("  Addr:", p)
					}
					for _, addr := range mes.AddrIpv6 {
						p := &netutil.Prefix{}
						p.IP = addr.Addr.Addr
						p.Length = int(addr.Addr.Length)
						fmt.Println("  Addr:", p)
					}
				case *pb.RouterIdUpdate:
					mes := res.(*pb.RouterIdUpdate)
					routerId := net.IP{}
					routerId = mes.RouterId
					fmt.Println("RouterId:", routerId)
				case *pb.Route:
					fmt.Println("")
				}
			case <-done:
				return
			}
		}
	}()

	fmt.Println("goroutine", runtime.NumGoroutine())

	// Subscribe to interface service.
	err = c.InterfaceSubscribe(DEFAULT_VRF)
	if err != nil {
		c.Stop()
	}
	// Subscribe to router id service.
	err = c.RouterIdSubscribe(DEFAULT_VRF)
	if err != nil {
		c.Stop()
	}

	// fmt.Println("-- sleep start --")
	// time.Sleep(time.Second * 3)
	// fmt.Println("-- sleep end --")

	// IPv4 route add.
	p, _ := netutil.ParsePrefix("10.0.0.0/24")
	nhop := netutil.ParseIPv4("10.211.55.1")
	r := &pb.Route{
		Type: pb.RIB_BGP,
		Prefix: &pb.Prefix{
			Addr:   p.IP,
			Length: uint32(p.Length),
		},
	}
	r.Nexthops = append(r.Nexthops, &pb.Nexthop{
		Addr:    nhop,
		Ifindex: 0,
	})
	c.RouteIPv4Add(r)
	//c.RouteIPv4Delete(r)

	// IPv6 route add.
	p6, _ := netutil.ParsePrefix("::1/128")
	r6 := &pb.Route{
		Type: pb.RIB_BGP,
		Prefix: &pb.Prefix{
			Addr:   p6.IP,
			Length: uint32(p6.Length),
		},
	}
	nhop6 := net.ParseIP("a::1")
	r6.Nexthops = append(r6.Nexthops, &pb.Nexthop{
		Addr:    nhop6,
		Ifindex: 0,
	})
	c.RouteIPv6Add(r6)

	// Close interafce stream -> Invoke all client EOF.
	//c.interfaceStream.CloseSend()

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
