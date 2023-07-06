#include <stdio.h>

#include "parse.h"
#include "run.h"

void print_result(struct result *);

int main(int argc, char *argv[]) {
    struct config conf;
    struct result res;

    init_config(&conf);
    // 解析命令行参数
    parse_args(argc, argv, &conf);
    // 执行
    run(&conf, &res);

    // 输出结果
    print_result(&res);
    return 0;
}

void print_result(struct result *res) {
    printf("\n%d\n", res->cpu_time_used);
    printf("%d\n", res->real_time_used);
    printf("%d\n", res->memory_used);
    printf("%d\n", res->signum);
    printf("%s", res->result);
}