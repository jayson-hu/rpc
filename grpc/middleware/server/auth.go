package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const  (
	ClientHeaderAccessKey = "client-id"
	ClientSecretAccessKey = "client-secret"
)

func NewClientCredential(ak, sk string) metadata.MD {
	return metadata.MD{
		ClientHeaderAccessKey:[]string{ak},
		ClientSecretAccessKey:[]string{sk},
	}
}


func NewAuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return (&GrpcAuther{}).UnaryServerInterceptor
}

func NewAuthStreamUnaryServerInterceptor() grpc.StreamServerInterceptor {
	return (&GrpcAuther{}).StreamServerInterceptor
}

type GrpcAuther struct {
	
}

//stream interceptor
func (a *GrpcAuther) StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error  {
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return   fmt.Errorf("ctx is not a contenx")

	}
	clientId, clientSecret := a.GetClientCredentialsFromMeta(md)
	if err := a.validateServiceCredential(clientId,clientSecret); err != nil{
		return err
	}
	return handler(srv,ss)
}

//request response 拦截器
func (a *GrpcAuther) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//读取凭证， 凭证放在meta信息 [http2 header]
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return  nil, fmt.Errorf("ctx is not a contenx")
	}
	//从客户端传递过来的凭证
	clientId, clientSecret := a.GetClientCredentialsFromMeta(md)
	if err := a.validateServiceCredential(clientId,clientSecret); err != nil{
		return nil, err
	}


	return handler(ctx, req)
}

func (a *GrpcAuther) GetClientCredentialsFromMeta(md metadata.MD) (clientId, clientSecret string){
	cakList := md[ClientHeaderAccessKey]
	chkList := md[ClientSecretAccessKey]
	if len(cakList) >0 {
		clientId = cakList[0]
	}
	if len(chkList) >0 {
		clientSecret = chkList[0]
	}
	return
}

func (a *GrpcAuther) validateServiceCredential(clientId, clientSecret string) error {
	if !(clientId == "admin" && clientSecret == "123456"){
		//返回一个认证错误，返回错误
		return status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	return nil
}























