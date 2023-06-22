export interface successType<T> {
    code:number,
    data:T,
    message:string,
}
export interface ErrorApi{
    code:number,
    message:string,
}