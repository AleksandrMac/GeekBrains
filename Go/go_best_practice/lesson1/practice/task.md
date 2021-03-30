# дополнтельные задания для практики из методички
- signal: улучшите код из примера, чтобы в зависимости от типа полученного сигнала запускались следующие функции  
    ```
    func gotSIGHUP() {  
        log.Print("SIGHUP handler")
    }  
    
    func gotSIGINT() {
        log.Print("SIGINT handler")
    }  
    
    func gotSIGTERM() {
        log.Print("SIGTERM handler")
    }  
    
    func gotSIGQUIT() {
        log.Print("SIGQUIT handler")
    }
    ```