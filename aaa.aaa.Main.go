package main

/* This is the main and single most important component of this framework. It starts your app and can also be used to gracefully shutdown the app.

   This component also helps to trigger the initialization of components which can not be initialized at the time when all init () functions are called on startup. For instance if component B needs component A, but doesn't know if component A will be initialized first, it would be better for component B's initialization function to be registered with this component. This way component A's initialiation can be guaranteed to occur before component B's initialization.

   Just as special case initializations are supported, special case deinit are also supported. See file "aaa.aaa.Main.InitDeinit.go", to the register init and deinit functions of components.
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
                        case _, _ = <- dShutdownChannel_Main: return
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
