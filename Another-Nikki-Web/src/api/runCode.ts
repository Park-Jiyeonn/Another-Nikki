import HttpServer from '../http/index';
import { CodeRet } from '@/types/runCode';

const runCode = (data:{
	input: string,
    lang: string,
    code: string,
}) => HttpServer.request<typeof data,CodeRet>({
    url: `/api/runcode`,
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
}