#include "hvm.h"
#include <stdio.h>
#define ERROR_BEAN_MISS_CONTRACT_JAR 9000
#define ERROR_BEAN_MISS_CONTRACT_ADDR 9001
#define ERROR_BEAN_MISS_INVOKE_CLASS 9002
#define ERROR_BEAN_MISS_INVOKE_BIN 9003
#define ERROR_BEAN_MISS_INVOKE_CLASS_NAME 9004
#define ERROR_BEAN_MISS_SENDER_ADDR 9005
#define ERROR_BEAN_MISS_TX_HASH 9006
#define ERROR_BEAN_MISS_ORIGIN_ADDR 9007

static unsigned long long* global_gas;


int hvm_invoke() {
    printf("23112");
    return ERROR_BEAN_MISS_ORIGIN_ADDR;
}


int hvm_deploy() {

    printf("Exce44444444444444ve\n");
    // 异常处理
    return ERROR_BEAN_MISS_TX_HASH;
}

