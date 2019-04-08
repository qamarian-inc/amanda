# Amanda

Amanda is a framework for developing cloud applications. This framework is based on Golang's package "net/http"; in other words, it's an HTTP-based framework. This framework comes with components like logger, alert raiser (a component that helps send 'distress notification' emails to the admin), and DBMS connection pool manager; in other words, all you possibly need to do, is to add more services.



## Using Amanda

Using this framework is as simple as downloading its source code, adding more services to the source code, then compiling the code, to produce your cloud application.

> In case you don't understand what is meant by a service, a service is a function that serves an HTTP request. See shutdown\_Manager (in file cx\_Shutdown\_Manager.go), to see how services are designed.

After adding your desired services, remember to modify the constants in file "cz\_Global.go", as appropriate, then register those new services, as appropriate.

> Registering a service is done by adding it to either variable software\_Service or admin\_Service, both of which can be found in file "cz\_Global.go". See file "cz\_Global.go", for an example.

