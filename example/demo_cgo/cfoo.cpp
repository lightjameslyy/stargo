#include "foo.hpp"
#include "foo.h"

Foo FooInit(int a) {
	cxxFoo * ret = new cxxFoo(a);
	return (void*)ret;
}

void FooFree(Foo f) {
	cxxFoo * foo = (cxxFoo*)f;
	delete foo;
}

void FooBar(Foo f) {
	cxxFoo * foo = (cxxFoo*)f;
	foo->Bar();
}

