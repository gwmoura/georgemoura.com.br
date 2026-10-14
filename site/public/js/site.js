// Dark Mode Toggle (o tema inicial é aplicado por um script inline no <head>)
(function () {
  var root = document.documentElement;
  var media = window.matchMedia('(prefers-color-scheme: dark)');

  function saved() {
    try { return localStorage.getItem('theme'); } catch (e) { return null; }
  }

  document.querySelectorAll('.theme-toggle').forEach(function (el) {
    el.addEventListener('click', function (event) {
      event.preventDefault();
      var theme = root.getAttribute('data-theme') === 'dark' ? 'light' : 'dark';
      root.setAttribute('data-theme', theme);
      try { localStorage.setItem('theme', theme); } catch (e) {}
    });
  });

  // Segue a preferência do sistema enquanto o usuário não escolher um tema
  media.addEventListener('change', function (e) {
    if (!saved()) root.setAttribute('data-theme', e.matches ? 'dark' : 'light');
  });
})();

// Lightbox: imagens com class="zoom" abrem ampliadas num <dialog> compartilhado
(function () {
  var images = document.querySelectorAll('img.zoom');
  if (!images.length || !window.HTMLDialogElement) return;

  var dialog = document.createElement('dialog');
  dialog.className = 'lightbox';
  dialog.innerHTML = '<button type="button" class="lightbox-close" aria-label="Fechar">&times;</button><img alt="">';
  document.body.appendChild(dialog);
  var big = dialog.querySelector('img');

  images.forEach(function (img) {
    img.addEventListener('click', function (event) {
      event.preventDefault();
      big.src = img.currentSrc || img.src;
      big.alt = img.alt;
      dialog.showModal();
    });
  });

  // Clique em qualquer lugar (imagem, fundo ou botão) fecha; Esc é nativo do <dialog>
  dialog.addEventListener('click', function () { dialog.close(); });
  dialog.addEventListener('close', function () { big.removeAttribute('src'); });
})();
