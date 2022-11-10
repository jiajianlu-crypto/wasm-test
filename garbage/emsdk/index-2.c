#include <stdio.h>
#include <stdlib.h>
#include <emscripten/emscripten.h>

// void consoleLog(int a, int b);

// EMSCRIPTEN_KEEPALIVE
int genGarbage(int a, int b) {
  int *p = (int *) malloc(10000000*sizeof(int));
  printf("%p", p);
  return 0;
}