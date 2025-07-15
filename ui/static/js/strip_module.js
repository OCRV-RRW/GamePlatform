const prewiewItems = document.querySelectorAll(".preview_item");
const stripItems = document.querySelectorAll(".strip_item");
let selectedStripItem = null;

stripItems.forEach(item => {
    item.addEventListener('click', onClickStripItem);
})

selectedStripItem = stripItems[0];
stripItems[0].classList.add('selected');
prewiewItems[0].classList.add('show');  

function onClickStripItem(event) {
    if (event.currentTarget === selectedStripItem)
        return;
    selectedStripItem.classList.remove('selected');
    prewiewItems[[...stripItems].indexOf(selectedStripItem)].classList.remove('show')
    
    selectedStripItem = event.currentTarget;
    
    selectedStripItem.classList.add('selected');
    prewiewItems[[...stripItems].indexOf(selectedStripItem)].classList.add('show')
}