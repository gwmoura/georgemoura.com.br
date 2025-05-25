$( document ).ready(function(){
  $(".button-collapse").sideNav();
  $('.materialboxed').materialbox();
  $('.carousel').carousel();

  const body = $('body');

  function activateDarkMode() {
    body.addClass('dark-mode');
  }

  function activateLightMode() {
    body.removeClass('dark-mode');
  }

  const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');

  // Set initial theme
  if (mediaQuery.matches) {
    activateDarkMode();
  } else {
    activateLightMode();
  }

  // Listen for changes
  mediaQuery.addEventListener('change', (event) => {
    if (event.matches) {
      activateDarkMode();
    } else {
      activateLightMode();
    }
  });
})
