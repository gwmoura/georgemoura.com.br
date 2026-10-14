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

// Aviso de cookies: guarda a escolha, atualiza o Consent Mode do GA4 e só
// carrega o AdSense depois que o visitante aceita
(function () {
  var banner = document.querySelector('.cookie-banner');
  if (!banner) return;

  function saved() {
    try { return localStorage.getItem('cookie-consent'); } catch (e) { return null; }
  }

  var adsLoaded = false;
  function loadAds() {
    var ad = document.querySelector('.anuncio');
    if (!ad || adsLoaded) return;
    adsLoaded = true;
    var script = document.createElement('script');
    script.async = true;
    script.src = 'https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-7444396631758040';
    script.crossOrigin = 'anonymous';
    document.head.appendChild(script);
    ad.hidden = false;
    (window.adsbygoogle = window.adsbygoogle || []).push({});
  }

  banner.querySelectorAll('[data-consent]').forEach(function (button) {
    button.addEventListener('click', function () {
      var choice = button.getAttribute('data-consent');
      try { localStorage.setItem('cookie-consent', choice); } catch (e) {}
      if (window.gtag) {
        gtag('consent', 'update', {
          analytics_storage: choice,
          ad_storage: choice,
          ad_user_data: choice,
          ad_personalization: choice
        });
      }
      if (choice === 'granted') loadAds();
      banner.hidden = true;
    });
  });

  document.querySelectorAll('.cookie-settings').forEach(function (link) {
    link.addEventListener('click', function (event) {
      event.preventDefault();
      banner.hidden = false;
    });
  });

  if (saved() === 'granted') loadAds();
  else if (!saved()) banner.hidden = false;
})();

// Botões com data-copy="#seletor" copiam o texto do elemento indicado
(function () {
  document.querySelectorAll('[data-copy]').forEach(function (button) {
    var target = document.querySelector(button.getAttribute('data-copy'));
    if (!target || !navigator.clipboard) { button.hidden = true; return; }
    var label = button.textContent;
    button.addEventListener('click', function () {
      navigator.clipboard.writeText(target.textContent.trim()).then(function () {
        button.textContent = 'Copiado!';
        setTimeout(function () { button.textContent = label; }, 2000);
      });
    });
  });
})();

// Meu IP: cada caixa consulta um endpoint que só responde por IPv4 ou só por
// IPv6, para mostrar os dois endereços mesmo que o servidor tenha visto apenas um
(function () {
  document.querySelectorAll('[data-ip-lookup]').forEach(function (box) {
    var output = box.querySelector('.ip');
    var button = box.querySelector('[data-copy]');

    function unavailable() {
      if (output.textContent.trim() !== 'Buscando…') return;
      output.textContent = 'Não disponível';
      output.classList.add('unavailable');
      if (button) button.hidden = true;
    }

    if (!window.fetch) return unavailable();
    fetch(box.getAttribute('data-ip-lookup'), { cache: 'no-store' })
      .then(function (res) { return res.ok ? res.json() : Promise.reject(); })
      .then(function (data) {
        output.textContent = data.ip;
        output.classList.remove('unavailable');
        if (button && navigator.clipboard) button.hidden = false;
      })
      .catch(unavailable);
  });
})();

// Gerador de senhas: usa crypto.getRandomValues, tudo acontece no navegador.
// Sempre tem minúscula, maiúscula e número (e símbolo, se marcado).
(function () {
  var form = document.querySelector('.password-generator');
  if (!form || !window.crypto || !crypto.getRandomValues) return;

  var groups = {
    lower: 'abcdefghijklmnopqrstuvwxyz',
    upper: 'ABCDEFGHIJKLMNOPQRSTUVWXYZ',
    digits: '0123456789',
    symbols: '!@#$%&*()-_=+[]{};:,.?/'
  };
  var output = document.getElementById('senha');
  var copy = document.querySelector('[data-copy="#senha"]');
  var lengthValue = form.querySelector('[data-length-value]');

  // Número aleatório em [0, max) sem viés de módulo
  function randomInt(max) {
    var limit = Math.floor(0x100000000 / max) * max;
    var buf = new Uint32Array(1);
    do { crypto.getRandomValues(buf); } while (buf[0] >= limit);
    return buf[0] % max;
  }

  function generate(length, sets) {
    var chars = sets.join('');
    var password;
    // Sorteia de novo até ter ao menos um caractere de cada grupo escolhido
    do {
      password = '';
      for (var i = 0; i < length; i++) password += chars[randomInt(chars.length)];
    } while (!sets.every(function (set) {
      return password.split('').some(function (c) { return set.indexOf(c) !== -1; });
    }));
    return password;
  }

  form.tamanho.addEventListener('input', function () {
    lengthValue.textContent = form.tamanho.value;
  });

  form.addEventListener('submit', function (event) {
    event.preventDefault();
    var sets = [groups.lower, groups.upper, groups.digits];
    if (form.symbols.checked) sets.push(groups.symbols);
    output.textContent = generate(parseInt(form.tamanho.value, 10), sets);
    if (copy && navigator.clipboard) copy.hidden = false;
  });
})();
