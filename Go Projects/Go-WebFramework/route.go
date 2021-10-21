package main

import "Go-WebFramework/framework"

func registerRouter(core *framework.Core)  {
	core.Get("foo",FooControllerHandler)

}
