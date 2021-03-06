package commands

import (
	"github.com/Summer2021-Project/cli/flag"
	"github.com/Summer2021-Project/cli/process"
	"github.com/Summer2021-Project/dotenv"
	"github.com/Summer2021-Project/grpc-skeleton/di"
	pb "github.com/Summer2021-Project/grpc-skeleton/protos"
	"github.com/Summer2021-Project/grpc-skeleton/services"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var netListener net.Listener

type GrpcServerCommand struct {
}

func (t *GrpcServerCommand) Main() {
	if flag.Match("d", "daemon").Bool() {
		process.Daemon()
	}

	addr := dotenv.Getenv("GIN_ADDR").String(":8080")
	logger := di.Logrus()

	// listen
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	netListener = listener

	// signal
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-ch
		logger.Info("Server shutdown")
		if err := listener.Close(); err != nil {
			panic(err)
		}
	}()

	// server
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &services.UserService{})

	// run
	welcome()
	logger.Infof("Server run %s", addr)
	if err := s.Serve(listener); err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
		panic(err)
	}
}
