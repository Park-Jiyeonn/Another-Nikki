import { ElMessage } from "element-plus"

const open_warning = (message: string) => {
    ElMessage({
        showClose: true,
        message: message,
        type: 'warning',
    })
}

const open_success = (message: string) => {
    ElMessage({
        showClose: true,
        message: message,
        type: 'success',
    })
}

export function send_warning(s : string) {
    open_warning(s)
}

export function send_success(s: string) {
    open_success(s)
}