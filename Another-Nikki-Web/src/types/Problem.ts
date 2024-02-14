export interface GetProblemByPageResp {
    problem_id:number,
    problem_title: string,
    created_time: string,
}

export interface GetProblemByIdResp {
    problem_id:number,
    problem_title: string,
    created_time: string, 
    problem_content: string,
}

export interface Problem {
    problem_id:number,
    problem_title: string,
    created_time: string,
    problem_content: string,
}