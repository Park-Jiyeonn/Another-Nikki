import HttpServer from '../http/index';
import { CodeRet } from '@/types/runCode';

const runCode = (data:{
	input: string,
    lang: string,
    code: string,
}) => HttpServer.request<typeof data,CodeRet>({
    url: `/api/runcode/run`,
    method: "POST",
    data
})

const judgeCode = (data:{
    user_id: string,
    user_name : string,
    problem_id: string,
    problem_name:string,
    language: string,
    code: string, 
}) => HttpServer.request<typeof data,CodeRet>({
    url: `/api/code/post`,
    method: "POST",
    data
})

const codeInit = (data:{
    lang: string,
}) => HttpServer.request<typeof data,{code:string, input:string}>({
    url: `/api/runcode/default_code`,
    method: "GET",
    data
})


export {
	runCode,
    codeInit,
    judgeCode,
}