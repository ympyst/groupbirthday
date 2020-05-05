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
)

var db *gorm.DB

type GroupBirthdayServer struct {
	contract.UnimplementedGroupBirthdayServer
}

func (*GroupBirthdayServer) GetGroups(ctx context.Context, req *contract.GetGroupsRequest) (*contract.GetGroupsReply, error) {
	fmt.Printf("➡️ GetGroupsRequest: %v\n", req)

	var member groupdb.Member
	db.First(&member, req.MemberId)
	var groups []groupdb.Group
	db.Model(&member).Related(&groups, "Groups")
	reply := &contract.GetGroupsReply{}
	for _, group := range groups {
		reply.Groups = append(reply.Groups, group.Name)
	}

	fmt.Printf("⬅️ GetGroupsReply: %v\n", reply)
	return reply, nil
}
func (*GroupBirthdayServer) GetMemberId(ctx context.Context, req *contract.GetMemberIdRequest) (*contract.GetMemberIdReply, error) {
	fmt.Printf("➡️ GetMemberIdRequest: %v\n", req)

	var member groupdb.Member
	db.Where("telegram_username = ?", req.TelegramUsername).Take(&member)
	reply := &contract.GetMemberIdReply{MemberId: int32(member.ID)}

	fmt.Printf("⬅️ GetMemberIdReply: %v\n", reply)
	return reply, nil
}
func (*GroupBirthdayServer) GetMemberBirthdays(ctx context.Context, req *contract.GetMemberBirthdaysRequest) (*contract.GetMemberBirthdaysReply, error) {
	fmt.Printf("➡️ GetMemberBirthdaysRequest: %v\n", req)
	var group []groupdb.Group
	db.Where("name = ?", req.GroupName).Take(&group)
	var members []groupdb.Member
	db.Model(&group).Related(&members, "Members")
	reply := &contract.GetMemberBirthdaysReply{}

	for _, member := range members {
		if member.Birthday == nil {
			continue
		}
		month := member.Birthday.Month()
		day := member.Birthday.Day()
		reply.MemberBirthdays = append(reply.MemberBirthdays, &contract.MemberBirthday{
			FirstName:            member.FirstName,
			LastName:             member.LastName,
			Day:                  int32(day),
			Month:                int32(month),
		})
	}

	fmt.Printf("⬅️ GetMemberBirthdaysReply: %v\n", reply)
	return reply, nil
}

func (*GroupBirthdayServer) SaveMember(ctx context.Context, req *contract.SaveMemberRequest) (*contract.SaveMemberReply, error) {
	fmt.Printf("➡️ SaveMemberRequest: %v\n", req)
	var member groupdb.Member
	db.Where("telegram_user_id = ?", req.TelegramUserId).Take(&member)

	if member.ID == 0 {
		member.TelegramUserId = req.TelegramUserId
		member.TelegramUsername = req.TelegramUsername
		member.FirstName = req.FirstName
		member.LastName = req.LastName
		db.Create(&member)
	}

	var group groupdb.Group
	db.Where("telegram_chat_id = ?", req.TelegramChatId).Take(&group)

	if group.ID == 0 {
		group.Name = req.TelegramChatName
		group.TelegramChatId = req.TelegramChatId
		db.Create(&group)
	}

	db.Model(&group).Association("Members").Append(member)

	reply := &contract.SaveMemberReply{MemberId: int32(member.ID)}
	fmt.Printf("⬅️ SaveMemberReply: %v\n", reply)
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