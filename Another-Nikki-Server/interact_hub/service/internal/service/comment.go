package service

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"context"
	"fmt"
	"math/rand"
	"time"

	pb "Another-Nikki/interact_hub/service/api"
)

type CommentService struct {
	pb.UnimplementedCommentServer
	dao biz.CommentRepo
	rx  *rand.Rand
}

func NewCommentService(dao biz.CommentRepo) *CommentService {
	return &CommentService{
		dao: dao,
		rx:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *CommentService) PostComment(ctx context.Context, req *pb.PostCommentReq) (resp *pb.PostCommentResp, err error) {
	resp = new(pb.PostCommentResp)
	if len(req.Content) == 0 {
		return nil, fmt.Errorf("发布的留言或评论不可以空空~")
	}
	if len(req.Content) > 400 {
		return nil, fmt.Errorf("发布的留言或评论太长啦~")
	}
	err = s.dao.PostComment(ctx, &biz.PostCommentReq{
		Content:    req.Content,
		ArticleId:  req.ArticleId,
		Username:   req.Username,
		UserAvatar: req.UserAvatar,
		ParentId:   req.ParentId,
		RootId:     req.RootId,
		ParentName: req.ParentName,
	})
	return
}
func (s *CommentService) GetCommentsByArticleId(ctx context.Context, req *pb.GetCommentsByArticleIdReq) (resp *pb.GetCommentsByArticleIdResp, err error) {
	resp = new(pb.GetCommentsByArticleIdResp)
	comments, err := s.dao.GetCommentsByArticleId(ctx, &biz.GetCommentsByArticleIdReq{
		ArticleId: req.GetArticleId(),
	})
	if err != nil {
		return nil, err
	}
	resp.Comments = make([]*pb.CommentDetail, len(comments.Comments))
	for _, val := range comments.Comments {
		comment := &pb.CommentDetail{
			CommentId:   val.CommentId,
			Content:     val.Content,
			Username:    val.Username,
			UserAvatar:  val.UserAvatar,
			ParentId:    val.ParentId,
			RootId:      val.RootId,
			CreatedTime: val.CreatedTime.Format(time.DateTime),
		}
		for _, child := range val.Children {
			comment.Children = append(comment.Children, &pb.CommentDetail{
				CommentId:   child.CommentId,
				Content:     child.Content,
				Username:    child.Username,
				UserAvatar:  child.UserAvatar,
				ParentId:    child.ParentId,
				RootId:      child.RootId,
				CreatedTime: child.CreatedTime.Format(time.DateTime),
				ParentName:  child.ParentName,
			})
		}
		resp.Comments = append(resp.Comments, comment)
	}
	return
}
func (s *CommentService) GetLastSevenComment(ctx context.Context, req *pb.GetLastSevenCommentReq) (resp *pb.GetLastSevenCommentResp, err error) {
	resp = new(pb.GetLastSevenCommentResp)
	comments, err := s.dao.GetLastSevenComment(ctx, &biz.GetLastSevenCommentReq{
		ArticleId: req.GetArticleId(),
		NumLimit:  req.Num,
	})
	if err != nil {
		return nil, err
	}
	resp.Comments = make([]*pb.CommentDetail, 0, len(comments.Comments))
	for _, val := range comments.Comments {
		comment := &pb.CommentDetail{
			CommentId:   val.CommentId,
			Content:     val.Content,
			Username:    val.Username,
			UserAvatar:  val.UserAvatar,
			ParentId:    val.ParentId,
			RootId:      val.RootId,
			CreatedTime: val.CreatedTime.Format(time.DateTime),
		}
		for _, child := range val.Children {
			comment.Children = append(comment.Children, &pb.CommentDetail{
				CommentId:   child.CommentId,
				Content:     child.Content,
				Username:    child.Username,
				UserAvatar:  child.UserAvatar,
				ParentId:    child.ParentId,
				RootId:      child.RootId,
				CreatedTime: child.CreatedTime.Format(time.DateTime),
				ParentName:  child.ParentName,
			})
		}
		resp.Comments = append(resp.Comments, comment)
	}
	return
}
func (s *CommentService) GetRandomComment(ctx context.Context, req *pb.GetRandomCommentReq) (resp *pb.GetRandomCommentResp, err error) {
	resp = new(pb.GetRandomCommentResp)
	n, err := s.dao.GetCommentSum(ctx)
	if err != nil {
		return nil, err
	}
	commentId := s.rx.Int63n(n) + 1
	comment, err := s.dao.GetCommentById(ctx, &biz.GetRandomCommentReq{
		ArticleId: req.ArticleId,
		CommentId: commentId,
	})
	time.Sleep(time.Millisecond * 100)
	resp.Comment = &pb.CommentDetail{
		CommentId:   commentId,
		Content:     comment.Comments.Content,
		Username:    comment.Comments.Username,
		UserAvatar:  comment.Comments.UserAvatar,
		ParentId:    comment.Comments.ParentId,
		RootId:      comment.Comments.RootId,
		CreatedTime: comment.Comments.CreatedTime.Format(time.DateTime),
	}
	return
}
