function removeChildren(id) {
    var container = document.getElementById(id);
    var child = container.lastElementChild; 
    while (child) {
        container.removeChild(child);
        child = container.lastElementChild;
    }
}
