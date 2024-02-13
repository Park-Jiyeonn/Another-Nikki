package service

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"Another-Nikki/pkg/log"
	"context"
	"time"

	pb "Another-Nikki/interact_hub/service/api"
)

type ProblemService struct {
	pb.UnimplementedProblemServer
	dao biz.ProblemRepo
}

func NewProblemService(dao biz.ProblemRepo) *ProblemService {
	return &ProblemService{
		dao: dao,
	}
}

func (s *ProblemService) PostProblem(ctx context.Context, req *pb.PostProblemReq) (resp *pb.PostProblemResp, err error) {
	resp = new(pb.PostProblemResp)
	if req.ProblemId != 0 {
		err = s.dao.UpdateProblem(ctx, &biz.UpdateProblemReq{
			ProblemId:          req.ProblemId,
			ProblemTitle:       req.ProblemTitle,
			ProblemDescription: req.ProblemDescription,
			ProblemContent:     req.ProblemContent,
		})
		return
	}
	err = s.dao.PostProblem(ctx, &biz.PostProblemReq{
		ProblemTitle:       req.ProblemTitle,
		ProblemDescription: req.ProblemDescription,
		ProblemContent:     req.ProblemContent,
	})
	return
}

func (s *ProblemService) GetProblemById(ctx context.Context, req *pb.GetProblemByIdReq) (*pb.GetProblemByIdResp, error) {
	resp := new(pb.GetProblemByIdResp)
	problem, err := s.dao.GetProblemById(ctx, &biz.GetProblemByIdReq{
		ProblemId: req.ProblemId,
	})
	if err != nil {
		log.Error(ctx, "GetProblemById err: %v", err)
		return nil, err
	}
	resp.ProblemTitle = problem.ProblemTitle
	resp.ProblemDescription = problem.ProblemDescription
	resp.CreateTime = problem.CreatedTime.Format(time.DateTime)
	resp.ProblemContent = problem.ProblemContent
	return resp, nil
}

func (s *ProblemService) GetProblemByPage(ctx context.Context, req *pb.GetProblemByPageReq) (resp *pb.GetProblemByPageResp, err error) {
	resp = new(pb.GetProblemByPageResp)
	problems, err := s.dao.GetProblemByPage(ctx, &biz.GetProblemByPageReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	resp.Problems = make([]*pb.ProblemPageDetail, 0, len(problems.Problems))
	for _, val := range problems.Problems {
		resp.Problems = append(resp.Problems, &pb.ProblemPageDetail{
			ProblemId:    val.ProblemId,
			ProblemTitle: val.ProblemTitle,
			CreateTime:   val.CreatedTime.Format(time.DateTime),
		})
	}
	return
}
