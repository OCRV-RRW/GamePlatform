import { pushModalWindow, popModalWindow } from "./modalBackgroundModule.js";

const screenshotModalImgContainer = document.querySelector('.screenshotModalImgContainer')
const screenshotModalImgElement = screenshotModalImgContainer.querySelector('img')
const screenshotModal = document.querySelector('.screenshotModalWindow')
const screenshotModalHeader = document.querySelector('.screenshotModalHeader')
const closeModalButton = screenshotModalHeader.querySelector('.closeModalButton')
const previousScreenshotButton = document.querySelector('.screenshotModalButtonChangeScreen.previous')
const nextScreenshotButton = document.querySelector('.screenshotModalButtonChangeScreen.next')
const countScreenshots = document.getElementById('countScreenshots')
const handlerImageIndexProxy = {
    set(target, prop, val) {
        target[prop] = val
        if (prop === 'value') {
            if (val >= sources.length - 1)
                nextScreenshotButton.style.display = 'none'
            else 
                nextScreenshotButton.style.display = 'flex'

            if (val <= 0) 
                previousScreenshotButton.style.display = 'none'
            else 
                previousScreenshotButton.style.display = 'flex'
            countScreenshots.innerHTML = `${val + 1} из ${sources.length} изображений`
        }
        return true;
    }
}
let currentIndex = { value: null }
let sources = null
let currentImageIndexProxy = null

export function show(currentSource, index, allSources) {
    sources = allSources
    currentImageIndexProxy = new Proxy(currentIndex, handlerImageIndexProxy)
    currentImageIndexProxy.value = index
    screenshotModalImgElement.src = currentSource 
    pushModalWindow(screenshotModal, hide)
    subscribe()
}

export function hide() {
    screenshotModal.style.display = 'none'
    unsubscribe()
}

function nextScreenshot() {
    if (currentIndex.value >= sources.length - 1)
        return
    currentImageIndexProxy.value++
    screenshotModalImgElement.src = sources[currentIndex.value]
}

function previousScreenshot() {
    if (currentIndex.value <= 0)
        return
    currentImageIndexProxy.value--
    screenshotModalImgElement.src = sources[currentIndex.value]
}

function onClose() {
    popModalWindow(screenshotModal)
}

function onKeyDown(event) {
    switch (event.key) {
        case 'ArrowLeft':
            previousScreenshot()
            break
        case 'ArrowRight':
            nextScreenshot()
            break
        case 'Escape':
            onClose()
            break
    }
};

function subscribe() {
    closeModalButton.addEventListener('click', onClose)
    nextScreenshotButton.addEventListener('click', nextScreenshot)
    previousScreenshotButton.addEventListener('click', previousScreenshot)
    document.addEventListener('keydown', onKeyDown)
}

function unsubscribe() {
    closeModalButton.removeEventListener('click', onClose)
    nextScreenshotButton.removeEventListener('click', nextScreenshot)
    previousScreenshotButton.removeEventListener('click', previousScreenshot)
    document.removeEventListener('keydown', onKeyDown)
}