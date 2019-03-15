#include <stdio.h>
#include "foo.hpp"

void cxxFoo::Bar(void) {
    printf("foo[%d]: bar...\n", this->a);
}

