package service

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"context"
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
	err = s.dao.PostComment(ctx, &biz.PostCommentReq{
		Content:      req.Content,
		ArticleId:    req.ArticleId,
		AuthorName:   req.AuthorName,
		AuthorAvatar: req.AuthorAvatar,
		ParentId:     req.ParentId,
		RootId:       req.RootId,
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
	resp.Comments = make([]*pb.Comments, len(comments.Comments))
	for _, val := range comments.Comments {
		resp.Comments = append(resp.Comments, &pb.Comments{
			Content:      val.Content,
			AuthorName:   val.AuthorName,
			AuthorAvatar: val.AuthorAvatar,
			ParentId:     val.ParentId,
			RootId:       val.RootId,
			CreateTime:   val.CreatedTime.Format(time.DateTime),
		})
	}
	return
}
func (s *CommentService) GetLastSevenComment(ctx context.Context, req *pb.GetLastSevenCommentReq) (resp *pb.GetLastSevenCommentResp, err error) {
	resp = new(pb.GetLastSevenCommentResp)
	comments, err := s.dao.GetLastSevenComment(ctx, &biz.GetLastSevenCommentReq{
		ArticleId: req.GetArticleId(),
	})
	if err != nil {
		return nil, err
	}
	resp.Comments = make([]*pb.Comments, len(comments.Comments))
	for _, val := range comments.Comments {
		resp.Comments = append(resp.Comments, &pb.Comments{
			Content:      val.Content,
			AuthorName:   val.AuthorName,
			AuthorAvatar: val.AuthorAvatar,
			ParentId:     val.ParentId,
			RootId:       val.RootId,
			CreateTime:   val.CreatedTime.Format(time.DateTime),
		})
	}
	return
}
func (s *CommentService) GetRandomComment(ctx context.Context, req *pb.GetRandomCommentReq) (resp *pb.GetRandomCommentResp, err error) {
	resp = new(pb.GetRandomCommentResp)
	n, err := s.dao.GetCommentSum(ctx)
	if err != nil {
		return nil, err
	}
	comment, err := s.dao.GetCommentById(ctx, &biz.GetRandomCommentReq{
		ArticleId: req.ArticleId,
		CommentId: s.rx.Int63n(n),
	})
	resp.Comments = &pb.Comments{
		Content:      comment.Comments.Content,
		AuthorName:   comment.Comments.AuthorName,
		AuthorAvatar: comment.Comments.AuthorAvatar,
		ParentId:     comment.Comments.ParentId,
		RootId:       comment.Comments.RootId,
		CreateTime:   comment.Comments.CreatedTime.Format(time.DateTime),
	}
	return
}
