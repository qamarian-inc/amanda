package main

/* This is the main and single most important component of this framework.
        It starts your app and can also be used to gracefully shutdown the app. */

import "runtime"

func main () {
        // All initialization functions registered with this component will be invoked.
        sInvokeInitFuncs ()

        // Continously check if global shutdown has been signalled:
        // if yes, the whole app built on this framework will shutdown.
        for {
                select {
                        case _, _ = <- dShutdownChannel_AAAAAA: return
                        default: continue
                }

                runtime.Gosched ()
        }

        // All deinitialization functions registered with this component will be invoked.
        sInvokeDnitFuncs ()
}

func iRegisterInitFunc (compID string, initFunc func (), precedentComp []string) {
}

func iRegisterDnitFunc (compID string, dnitFunc func (), precedentComp []string) {
}

func iShutdown_AAAAAA () { // To shutdown your app gracefully, this interface can be called.
        dShutdownChannel_AAAAAA <- true
}

func sInvokeInitFuncs () {}

func sInvokeDnitFuncs () {}

var dShutdownChannel_AAAAAA chan bool = make (chan bool, 1)
