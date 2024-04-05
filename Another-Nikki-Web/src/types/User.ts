export interface UserType {
    token:          string,
    user_id:        number,
    username:      string,
    avatar:         string,
    description : string,
}

export interface Commits {
    judge_id: number,
    problem_name: string,
    compile_status: string,
    judge_status: string,
    cpu_time_used: string,
    memory_used: string,
    language: string,
    created_time: string,
    problem_id :number,
    compile_log: string,
}

export interface Commit {
    problem_name: string,
    compile_status: string,
    judge_status: string,
    cpu_time_used: string,
    memory_used: string,
    language: string,
    created_time: string,
    problem_id :number,
    code: string,
    user_id:number,
    username:string,
    compile_log: string,
}