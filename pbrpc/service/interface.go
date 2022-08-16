package service

const  (
	SERVICE_NAME = "HelloService"
)


type HelloService interface {
	Hello(request *service.Request, response *service.Response) error

}
