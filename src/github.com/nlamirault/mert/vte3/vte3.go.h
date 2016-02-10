#ifndef GO_VTE3_H
#define GO_VTE3_H

#include <vte/vte.h>
#include <gtk/gtk.h>
#include <stdlib.h>

static inline VteTerminal* toVTerminal(GtkWidget* w) {
  return VTE_TERMINAL(w);
}

static inline char** make_strings(int count) {
        return (char**)malloc(sizeof(char*) * count);
}

static inline void set_string(char** strings, int n, char* str) {
        strings[n] = str;
}

static inline char** argv() {
	char **argv = malloc(sizeof(char*) * 2);
	argv[0] = "/bin/bash";
	argv[1] = NULL;
	return argv;
}


#endif
