package main

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/grpc"
	"groupbirthday/proto"
	"groupbirthday/groupdb"
	"log"
	"net"
	"os"
)

var db *gorm.DB

type GroupBirthdayServer struct {
	proto.UnimplementedGroupBirthdayServer
}

func (*GroupBirthdayServer) GetGroups(ctx context.Context, req *proto.GetGroupsRequest) (*proto.GetGroupsReply, error) {
	var member groupdb.Member
	db.First(&member, req.MemberId)
	var groups []groupdb.Group
	db.Model(&member).Related(&groups, "Groups")
	reply := &proto.GetGroupsReply{}
	for _, group := range groups {
		reply.Groups = append(reply.Groups, group.Name)
	}
	return reply, nil
}
func (*GroupBirthdayServer) GetMemberId(ctx context.Context, req *proto.GetMemberIdRequest) (*proto.GetMemberIdReply, error) {
	var member groupdb.Member
	db.Where("telegram_username = ?", req.TelegramUsername).Take(&member)
	reply := &proto.GetMemberIdReply{MemberId: int32(member.ID)}
	return reply, nil
}

func main() {
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("🤖 Now listening port %v\n", port)

	db, err = gorm.Open("postgres", os.Getenv("DB_DSN"))
	defer db.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("💾 Connected to database")

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	gbdServer := &GroupBirthdayServer{}
	proto.RegisterGroupBirthdayServer(grpcServer, gbdServer)
	grpcServer.Serve(lis)
}