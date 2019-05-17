package main

/* This is the main and single most important component of this framework. It starts your app and can also be used to gracefully shutdown the app.

   This component also helps trigger the initialization of components which can not be initialized at the time when all init () functions are called on startup.

   For instance, if a component B depends on a component A, but can't assertain if component A's init () function will be called before its own init () function or knows that component A's init () function would not be called before its own init () function, component B should have a custom initialization function (e.g. iInit_ComponentB ()) which can then be registered with this component (component aaa.aaa); then this component (aaa.aaa) would call the init function iInit_ComponentB () after all init () functions have been called.

   This way component A's init () function can be guaranteed to be executed before component B's iInit_ComponentB () function.

   Just as special case initializations are supported, special case deinit are also supported. See file "aaa.aaa.Main.InitDeinit.go", to the register custom init and deinit functions of your components.
*/

import "runtime"

func main () {
	// Triggering all special case init functions.
	for _, initFunc := range dInitFunction_Main {
		initFunc ()
	}

	// Continously check if global shutdown has been signalled: if yes, the whole app built on this framework will shutdown.
        for {
                select {
                        case _, _ = <- dShutdownChannel_Main: break
                        default: continue
                }

                runtime.Gosched ()
        }

        // Triggering all special case deinit functions.
	for _, deinitFunc := range dDeinitFunction_Main {
		deinitFunc ()
	}
}

var (
        dShutdownChannel_Main chan bool = make (chan bool, 1) // Channel that'll be used to signal global shutdown of the app.
)

func iShutdown_Main () { // To shutdown your app gracefully, this interface can be called.
        dShutdownChannel_Main <- true
}
