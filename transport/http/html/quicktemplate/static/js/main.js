/*document.getElementsByClassName("footer__cutout")[0].addEventListener("click", function(e) {
  alert(e.target.outerHTML);
});
document.getElementsByClassName("footer__content")[0].addEventListener("click", function(e) {
  alert(e.target.outerHTML);
});
document.getElementsByClassName("footer")[0].addEventListener("click", function(e) {
  alert(e.target.outerHTML);
});*/

// firefox sometimes triggers touch events on elements behind the footer if you touch at the very bottom of it
// probably because there is a small gab between the footer and the screen boarder
let accordions = document.getElementsByClassName("accordion");
for(let i = 0; i< accordions.length; i++) {
  accordions[i].addEventListener("click", function(e) {
    if(!isVisible(e.currentTarget)) {
      e.preventDefault();
    }
  });
}

function isVisible(elem) {
  if (!(elem instanceof Element)) throw Error('DomUtil: elem is not an element.');
  if (elem.offsetWidth + elem.offsetHeight + elem.getBoundingClientRect().height +
      elem.getBoundingClientRect().width === 0) {
      return false;
  }
  const elemCenter   = {
      x: elem.getBoundingClientRect().left + elem.offsetWidth / 2,
      y: elem.getBoundingClientRect().top + elem.offsetHeight / 2
  };
  if (elemCenter.x < 0) return false;
  if (elemCenter.x > (document.documentElement.clientWidth || window.innerWidth)) return false;
  if (elemCenter.y < 0) return false;
  if (elemCenter.y > (document.documentElement.clientHeight || window.innerHeight)) return false;
  let pointContainer = document.elementFromPoint(elemCenter.x, elemCenter.y);
  do {
      if (pointContainer === elem) return true;
  } while (pointContainer = pointContainer.parentNode);
  return false;
}
