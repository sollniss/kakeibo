/*!
 * swiped-events.js - v@version@
 * Pure JavaScript swipe events
 * https://github.com/john-doherty/swiped-events
 * @inspiration https://stackoverflow.com/questions/16348031/disable-scrolling-when-touch-moving-certain-element
 * @author John Doherty <www.johndoherty.info>
 * @license MIT
 */
(function (document) {

  'use strict';
  document.addEventListener('touchstart', handleTouchStart, false);
  document.addEventListener('touchmove', handleTouchMove, false);
  document.addEventListener('touchend', handleTouchEnd, false);

  const swipeThreshold = 20; // default 20px
  const swipeTimeout = 500;    // default 500ms

  let xDown = null;
  let yDown = null;
  let xDiff = null;
  let yDiff = null;
  let timeDown = null;
  let startEl = null;

  /**
   * Fires swiped event if swipe detected on touchend
   * @param {object} e - browser event object
   * @returns {void}
   */
  function handleTouchEnd(e) {

      // if the user released on a different target, cancel!
      if (startEl !== e.target) return;
      let timeDiff = Date.now() - timeDown;


      if (Math.abs(xDiff) > Math.abs(yDiff)) { // most significant
          if (Math.abs(xDiff) > swipeThreshold && timeDiff < swipeTimeout) {
              if (xDiff > 0) { // swiped-left
                navNext();
              }
              else { // swiped-right
                navPrev();
              }
          }
      }

      // reset values
      xDown = null;
      yDown = null;
      timeDown = null;
  }

  /**
   * Records current location on touchstart event
   * @param {object} e - browser event object
   * @returns {void}
   */
  function handleTouchStart(e) {

      startEl = e.target;

      timeDown = Date.now();
      xDown = e.touches[0].clientX;
      yDown = e.touches[0].clientY;
      xDiff = 0;
      yDiff = 0;
  }

  /**
   * Records location diff in px on touchmove event
   * @param {object} e - browser event object
   * @returns {void}
   */
  function handleTouchMove(e) {

      if (!xDown || !yDown) return;

      let xUp = e.touches[0].clientX;
      let yUp = e.touches[0].clientY;

      xDiff = xDown - xUp;
      yDiff = yDown - yUp;
  }

  function navPrev() {
    location.href = document.getElementById("nav-prev").href;
  }

  function navNext() {
    location.href = document.getElementById("nav-next").href;
  }

}(document));