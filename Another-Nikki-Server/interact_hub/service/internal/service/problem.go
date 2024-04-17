package service

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"Another-Nikki/pkg/log"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	pb "Another-Nikki/interact_hub/service/api"
)

const OnlineJudgePath = "../../onlineJudge/problemData/"

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
	problemId := req.ProblemId
	if req.ProblemId != 0 {
		err = s.dao.UpdateProblem(ctx, &biz.UpdateProblemReq{
			ProblemId:          req.ProblemId,
			ProblemTitle:       req.ProblemTitle,
			ProblemDescription: req.ProblemDescription,
			ProblemContent:     req.ProblemContent,
		})
		if err != nil {
			return
		}
	} else {
		var ret *biz.PostProblemResp
		ret, err = s.dao.PostProblem(ctx, &biz.PostProblemReq{
			ProblemTitle:       req.ProblemTitle,
			ProblemDescription: req.ProblemDescription,
			ProblemContent:     req.ProblemContent,
		})
		if err != nil {
			return
		}
		problemId = ret.ProblemId
	}
	req.DataIn = strings.ReplaceAll(req.DataIn, "data:application/octet-stream;base64,", "")
	req.DataOut = strings.ReplaceAll(req.DataOut, "data:application/octet-stream;base64,", "")
	if req.DataIn != "" {
		var dataInBytes []byte
		if dataInBytes, err = base64.StdEncoding.DecodeString(req.DataIn); err != nil {
			return
		}
		if err = writeContentInFile(problemId, req.ProblemTitle, string(dataInBytes), "input.in"); err != nil {
			return
		}
	}
	if req.DataOut != "" {
		var dataOutBytes []byte
		if dataOutBytes, err = base64.StdEncoding.DecodeString(req.DataOut); err != nil {
			return
		}
		if err = writeContentInFile(problemId, req.ProblemTitle, string(dataOutBytes), "output.out"); err != nil {
			return
		}
	}
	return
}

// 如果内容是空的，按这里的需求，是不会继续创建文件下去的
func writeContentInFile(ID int64, problemName, content, filename string) (err error) {
	if content == "" {
		return
	}
	var f *os.File
	dirPath := fmt.Sprintf(OnlineJudgePath+"%d.%s", ID, problemName)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return
	}
	path := dirPath + "/" + filename
	f, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	if _, err = io.WriteString(f, content); err != nil {
		return
	}
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
	resp.CreatedTime = problem.CreatedTime.Format(time.DateTime)
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
			CreatedTime:  val.CreatedTime.Format(time.DateTime),
		})
	}
	return
}
