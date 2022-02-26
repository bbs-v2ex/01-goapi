package hanlder

// 名称冲突问题
const HelloServiceName = "hanlder/HelloService"

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = request + *reply
	return nil
}
