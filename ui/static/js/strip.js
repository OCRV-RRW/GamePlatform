import { setPreview, previewContainsVideo } from "./preview.js"

let selectedStripItem = null
let indexSelectedStripItem = null
let intervalID = null
let stripItems = null

init()

function onClickStripItem(event) {
    if (event.currentTarget === selectedStripItem)
        return
    selectedStripItem.classList.remove('selected')
    selectedStripItem = event.currentTarget
    selectedStripItem.classList.add('selected')
    indexSelectedStripItem = [...stripItems].indexOf(selectedStripItem)
    setPreview(indexSelectedStripItem)

    if (intervalID !== null)
        clearInterval(intervalID)
    if (!previewContainsVideo())
        intervalID = setInterval(switchToNextSelectedStripItem, 5000)
}

function switchToNextSelectedStripItem() {
    selectedStripItem.classList.remove('selected')
    indexSelectedStripItem = (indexSelectedStripItem + 1) % stripItems.length
    selectedStripItem = stripItems[indexSelectedStripItem]
    selectedStripItem.classList.add('selected')
    setPreview(indexSelectedStripItem)
    if (previewContainsVideo()) {
        if (intervalID !== null)
            clearInterval(intervalID)
        intervalID = null
    }
}

export function init() {
    stripItems = document.querySelectorAll(".strip_item")

    stripItems.forEach(item => {
        item.addEventListener('click', onClickStripItem)
    })

    selectedStripItem = stripItems[0]
    indexSelectedStripItem = 0
    stripItems[0].classList.add('selected')
    setPreview(0)

    if (!previewContainsVideo())
        intervalID = setInterval(switchToNextSelectedStripItem, 5000)
}