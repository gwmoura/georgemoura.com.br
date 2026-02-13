$( document ).ready(function(){
  $(".button-collapse").sideNav();
  $('.materialboxed').materialbox();
  $('.carousel').carousel();
  
  // Dark Mode Toggle
  const toggleDarkMode = function() {
    const currentTheme = document.documentElement.getAttribute('data-theme');
    const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
    
    document.documentElement.setAttribute('data-theme', newTheme);
    localStorage.setItem('theme', newTheme);
    
    // Atualiza ícone
    const icon = newTheme === 'dark' ? 'brightness_7' : 'brightness_4';
    $('#dark-mode-toggle i, #dark-mode-toggle-mobile i').text(icon);
  };
  
  // Event listeners para ambos os botões (desktop e mobile)
  $('#dark-mode-toggle, #dark-mode-toggle-mobile').on('click', toggleDarkMode);
  
  // Carrega preferência salva ou detecta preferência do sistema
  const savedTheme = localStorage.getItem('theme');
  const systemPrefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
  const theme = savedTheme || (systemPrefersDark ? 'dark' : 'light');
  
  if (theme === 'dark') {
    document.documentElement.setAttribute('data-theme', 'dark');
    $('#dark-mode-toggle i, #dark-mode-toggle-mobile i').text('brightness_7');
  }
  
  // Escuta mudanças na preferência do sistema
  window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', function(e) {
    if (!localStorage.getItem('theme')) {
      const newTheme = e.matches ? 'dark' : 'light';
      document.documentElement.setAttribute('data-theme', newTheme);
      const icon = newTheme === 'dark' ? 'brightness_7' : 'brightness_4';
      $('#dark-mode-toggle i, #dark-mode-toggle-mobile i').text(icon);
    }
  });
})
