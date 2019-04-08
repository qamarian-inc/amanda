# Amanda

Amanda is a framework for developing cloud applications, and it is HTTPS-based. This framework takes a form quite different from many other frameworks. Rather than existing as a package that is imported, the framework is a software on its own. It can be turned into a custom cloud application, by expanding its capabilites and/or adding services to it.

The framework comes with useful components like logger, and alert raiser. The alert raiser is a component which can be used to send 'distress emails' to the admin of a cloud application. For instances, if something happens at runtime and requires immediate attention of the admin, the alert raiser can be used to send an email to the admin. The framework furthermore supports the use of configuration file.

Before doing anything with this framework, ensure all constants in file "cx30.Global.go" have been modified as appropriate.

