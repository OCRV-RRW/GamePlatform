const modalBackground = document.querySelector('.modalBackground')

let stackModalWindows = []
let activated = false
modalBackground.addEventListener('click', onClickModalWindow)

export function pushModalWindow(modalWindow, hideHandler, autoShow=true) {
    stackModalWindows.push({modalWindow, hideHandler})
    if (autoShow)
        modalWindow.style.display = 'block'
    if (stackModalWindows.length >= 1 && activated === false)
        activate()
}

export function popModalWindow(modalWindow=null) {
    let hideHandler = null
    let _ = null
    if (modalWindow === null)
        ({_, hideHandler} = stackModalWindows.pop())
    else 
        ({_, hideHandler} = stackModalWindows
            .pop(stackModalWindows
                .findIndex(item => item.modalWindow === modalWindow)))
    hideHandler()
    if (stackModalWindows.length === 0 && activated === true)
        deactivate()
}

function onClickModalWindow() {
    popModalWindow()
}

function activate() {
    modalBackground.style.display = 'block' 
    activated = true
}

function deactivate() {
    modalBackground.style.display = 'none' 
    activated = false
}