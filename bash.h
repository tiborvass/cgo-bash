#include <stdlib.h>
#include <string.h>

// copied from command.h
typedef struct word_desc {
  char *word;
  int flags;
} WORD_DESC;

// copied from command.h
typedef struct word_list {
  struct word_list *next;
  WORD_DESC *word;
} WORD_LIST;

// copied from general.h
typedef int sh_builtin_func_t (WORD_LIST*);

// adapted from builtins.h
typedef struct {
	char *name;
	sh_builtin_func_t *function;
	int flags;
	char * const *long_doc;
	const char *short_doc;
	char *handle;
} go_builtin;

extern go_builtin* (*external_lookup_builtin)(char*);
extern go_builtin *current_builtin;
extern int (builtin_func_wrapper)(WORD_LIST*);
extern int Main(int argc, char **argv, char **env);

static int go_builtins_sz = 0;
static go_builtin *go_builtins = 0;

static inline go_builtin* go_lookup_builtin(char *name) {
	for (go_builtin *b = go_builtins; b && b->name; b++) {
		if (!strcmp(name, b->name)) {
			return b;
		}
	}
	return 0;
}

static inline void go_add_builtin(go_builtin b) {
	go_builtins_sz++;
	go_builtins = realloc(go_builtins, sizeof(go_builtin) * (go_builtins_sz + 1));
	go_builtins[go_builtins_sz-1] = b;
	go_builtins[go_builtins_sz] = (go_builtin){};
}

static inline void go_del_builtin(char *name) {
	for (go_builtin *b = go_builtins; b && b->name; b++) {
		if (!strcmp(name, b->name)) {
			for (go_builtin *b2 = b+1; b && b->name; b2++, b++) {
				*b = *b2;
			}
			go_builtins_sz--;
			return;
		}
	}
}

static inline void go_init() {
	external_lookup_builtin = &go_lookup_builtin;
}

