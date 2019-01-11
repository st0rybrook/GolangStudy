package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	_ "google.golang.org/grpc/peer"
)

const KEY_SESSIONID = "session-id"
const KEY_ORIGIN_IP = "origin-ip"
const KEY_ORIGIN_PORT = "origin-port"
const KEY_ORIGIN_SERVICE = "origin-service"

var GRPC_TIMEOUT = 10 * time.Second

// NewContext 创建context
func NewContext(ctx context.Context, timeout int) context.Context {
	if ctx == nil {
		ctx = context.Background()
		if timeout == 0 {
			ctx, _ = context.WithTimeout(ctx, GRPC_TIMEOUT)
		} else if timeout > 0 {
			ctx, _ = context.WithTimeout(ctx, time.Second*time.Duration(timeout))
		}
	}
	return ctx
}

// NewContextWithAddr 创建带有ip、port的context
func NewContextWithAddr(ctx context.Context, ip string, port int, serviceName string, timeout int) context.Context {
	ctx = NewContext(ctx, timeout)
	sessionID := "tangshenzheng"
	md := metadata.Pairs(KEY_SESSIONID, sessionID, KEY_ORIGIN_IP, ip, KEY_ORIGIN_PORT, fmt.Sprintf("%d", port), KEY_ORIGIN_SERVICE, serviceName)
	//兼容目前ctx，同时将MD带入普通ctx传输
	ctx = metadata.NewOutgoingContext(ctx, md)
	ctx = context.WithValue(ctx, "MD", md)
	return ctx
}

// NewContextWithAddr 创建带有ip、port的context
func NewContextWithSessionID(sessionID string, ip string, port int, serviceName string, timeout int) context.Context {
	if sessionID == "" {
		sessionID = "tangshenzheng"
	}
	ctx := NewContext(nil, timeout)
	md := metadata.Pairs(KEY_SESSIONID, sessionID, KEY_ORIGIN_IP, ip, KEY_ORIGIN_PORT, fmt.Sprintf("%d", port), KEY_ORIGIN_SERVICE, serviceName)
	//兼容目前ctx，同时将MD带入普通ctx传输
	ctx = metadata.NewOutgoingContext(ctx, md)
	ctx = context.WithValue(ctx, "MD", md)
	return ctx
}

// DumpContext 获取context信息
func DumpContext(ctx context.Context) string {
	md := ExtractMDFromCtx(ctx)
	res := fmt.Sprintf("context: [%s]%s, [%s]%s, [%s]%s, [%s]%s", KEY_SESSIONID, md[KEY_SESSIONID], KEY_ORIGIN_IP, md[KEY_ORIGIN_IP], KEY_ORIGIN_PORT, md[KEY_ORIGIN_PORT], KEY_ORIGIN_SERVICE, md[KEY_ORIGIN_SERVICE])
	return res
}
func ExtractMDFromCtx(ctx context.Context) metadata.MD {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md, ok = metadata.FromIncomingContext(ctx)
	}
	//Value returns the value associated with this context for key, or nil, if no value is associated with key
	if !ok {
		//interface类型转换获取MD
		md, _ = ctx.Value("MD").(metadata.MD)
	}
	return md
}

func main() {
	ip := "192.168.153.94"
	port := 8080
	sessionID := "tangshenzheng"
	serviceName := "foroutgoing"
	ctx := NewContextWithSessionID(sessionID, ip, port, serviceName, 10)
	//ctx:=NewContext()
	md := ctx.Value("MD").(metadata.MD)
	fmt.Println(md)
	//Value returns the value associated with this context for key, or nil
	md1, _ := metadata.FromIncomingContext(ctx)
	mp := make(map[string][]string)
	fmt.Println(mp == nil)

	fmt.Println(md1 == nil)
	fmt.Println(md1)

	fmt.Println(ctx)
	ctx1 := context.Background()
	ctx1 = context.WithValue(ctx1, "session-id", sessionID)
	fmt.Println(ctx1)
	fmt.Println(ctx1.Value("session-id"))
	fmt.Println(md1.Get("tang"))
	if len(md1.Get("tang")) > 0 {

	}
	fmt.Println(DumpContext(context.WithValue(context.Background(), "MD", md)))
}
