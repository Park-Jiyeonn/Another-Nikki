export interface LogsType {
    level: string;
    ts: string;
    service_name: string;
    trace_id: string;
    ip: string;
    platform: string;
    url: string;
    msg: string;
    args: string;
    stack: string;
    code: string;
    log_id: number;
}
