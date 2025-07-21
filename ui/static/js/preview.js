import { show as showScreenshotModal } from "./screenshotModal.js"

let previewItemsScreenshotSources = []
let lastPreview = null
let lastIndex = null
let lastIndexImage = null

const prewiewItems = document.querySelectorAll(".preview_item")

prewiewItems.forEach(item => {
    const src = getSourceImage(item)
    if (src !== null) {
        previewItemsScreenshotSources.push(src)
        item.addEventListener('click', onClickPreview)
    }
})

export function setPreview(index) {
    if (lastPreview !== null)
        lastPreview.classList.remove('show')

    prewiewItems[index].classList.add('show')
    lastIndexImage = previewItemsScreenshotSources.findIndex(item => item === getSourceImage(prewiewItems[index]))
    lastIndex = index
    lastPreview = prewiewItems[lastIndex]
}

export function previewContainsVideo(child=null) {
    let element = child
    if (child === null)
        element = lastPreview
    
    if (element.tagName === 'VIDEO') return true
    for (const child of element.children) {
        if (previewContainsVideo(child)) return true
    }
    return false
}

function getSourceImage(element) {
    let img = element.querySelector('img')
    if (img === null)
        return null
    return img.src
}

function onClickPreview(event) {
    showScreenshotModal(getSourceImage(event.currentTarget), lastIndexImage, previewItemsScreenshotSources)
}