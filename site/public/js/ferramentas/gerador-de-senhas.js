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
