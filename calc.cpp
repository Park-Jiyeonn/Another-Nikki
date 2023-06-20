#include <stdio.h>
#include <sys/time.h>
#include <sys/wait.h>
#include <sys/resource.h>
#include <unistd.h>

int main(int argc, char *argv[]) {
    if (argc < 2) {
        printf("Usage: %s cmd\n", argv[0]);
        printf("e.g: %s echo Hello\n", argv[0]);
        return 1;
    }

    struct timeval start, end;
    pid_t pid;
    int time_used, cpu_time_used, memory_used, signum;

    // 设置 CPU 时间限制
    struct rlimit rl;
    rl.rlim_cur = rl.rlim_max = 1;
    setrlimit(RLIMIT_CPU, &rl);

    gettimeofday(&start, NULL);  // 运行前计时
    if ((pid = fork()) == 0) {
        execvp(argv[1], &argv[1]);
    }

    int status;
    struct rusage ru;
    wait4(pid, &status, 0, &ru);

    gettimeofday(&end, NULL);  // 运行后计时
    time_used = (int) (end.tv_sec * 1000 + end.tv_usec / 1000 - start.tv_sec * 1000 -
                       start.tv_usec / 1000);
    cpu_time_used =
            ru.ru_utime.tv_sec * 1000 + ru.ru_utime.tv_usec / 1000 +
            ru.ru_stime.tv_sec * 1000 + ru.ru_stime.tv_usec / 1000;

    memory_used = ru.ru_maxrss;
    if (WIFEXITED(status)) {
        signum = WEXITSTATUS(status);
    } else if (WIFSIGNALED(status)) {
        signum = WTERMSIG(status);
    } else if (WIFSTOPPED(status)) {
        signum = WSTOPSIG(status);
    } else {
        signum = 0;
    }

    printf("\n");
    printf("%dms\n", cpu_time_used);
    printf("%dkb\n", memory_used);
    printf("%d", signum);
    return 0;
}