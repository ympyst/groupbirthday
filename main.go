package main

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/grpc"
	contract "groupbirthday/contract"
	"groupbirthday/groupdb"
	"log"
	"net"
	"os"
	"time"
)

var db *gorm.DB

type GroupBirthdayServer struct {
	contract.UnimplementedGroupBirthdayServer
}

func (*GroupBirthdayServer) GetGroups(ctx context.Context, req *contract.GetGroupsRequest) (*contract.GetGroupsReply, error) {
	fmt.Println(req)

	var member groupdb.Member
	db.First(&member, req.MemberId)
	var groups []groupdb.Group
	db.Model(&member).Related(&groups, "Groups")
	reply := &contract.GetGroupsReply{}
	for _, group := range groups {
		reply.Groups = append(reply.Groups, group.Name)
	}

	fmt.Println(reply)
	return reply, nil
}
func (*GroupBirthdayServer) GetMemberId(ctx context.Context, req *contract.GetMemberIdRequest) (*contract.GetMemberIdReply, error) {
	fmt.Println(req)

	var member groupdb.Member
	db.Where("telegram_username = ?", req.TelegramUsername).Take(&member)
	reply := &contract.GetMemberIdReply{MemberId: int32(member.ID)}

	fmt.Println(reply)
	return reply, nil
}
func (*GroupBirthdayServer) GetMemberBirthdays(ctx context.Context, req *contract.GetMemberBirthdaysRequest) (*contract.GetMemberBirthdaysReply, error) {
	fmt.Println(req)
	var group []groupdb.Group
	db.First(&group)
	var members []groupdb.Member
	db.Model(&group).Related(&members, "Members")
	reply := &contract.GetMemberBirthdaysReply{}

	for _, member := range members {
		bd, err := time.Parse(time.RFC3339, member.Birthday)
		if err != nil {
			panic(err)
		}
		month := bd.Month()
		day := bd.Day()
		reply.MemberBirthdays = append(reply.MemberBirthdays, &contract.MemberBirthday{
			FirstName:            member.FirstName,
			LastName:             member.LastName,
			Day:                  int32(day),
			Month:                int32(month),
		})
	}

	fmt.Println(reply)
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
	contract.RegisterGroupBirthdayServer(grpcServer, gbdServer)
	grpcServer.Serve(lis)
}