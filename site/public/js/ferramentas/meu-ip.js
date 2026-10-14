// Meu IP: cada caixa consulta um endpoint que só responde por IPv4 ou só por
// IPv6, para mostrar os dois endereços mesmo que o servidor tenha visto apenas um
(function () {
  document.querySelectorAll('[data-ip-lookup]').forEach(function (box) {
    var output = box.querySelector('.result');
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

