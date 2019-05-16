package main

/* This virtual component is a Zain (see meaning at "Zain" github.com/qamarian_word/some) that can be reported to whenever a critical event occurs.

   For example if a server can no more accept new connections, due to an error, this component can be notified. This is very useful if you want certain actions to be carried out whenever a critical event occurs.

   */

const iReport_vcCriticalEventZain_Amanda func (string) // An actual component that implements this virtual component, must accept a parametre, which would be a description of the critical event that occured.
