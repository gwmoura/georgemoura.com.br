// Conversor de timestamp Unix: epoch ⇄ data, no fuso do navegador e em UTC
(function () {
  var now = document.getElementById('agora');
  if (!now) return;
  var epoch = document.getElementById('epoch');
  var epochResult = document.getElementById('epoch-resultado');
  var date = document.getElementById('data');
  var dateResult = document.getElementById('data-resultado');
  var timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;

  document.querySelectorAll('[data-timezone]').forEach(function (el) { el.textContent = timezone; });

  var local = new Intl.DateTimeFormat('pt-BR', { dateStyle: 'full', timeStyle: 'long' });
  var utc = new Intl.DateTimeFormat('pt-BR', { dateStyle: 'full', timeStyle: 'long', timeZone: 'UTC' });

  function render(dl, rows) {
    dl.textContent = '';
    rows.forEach(function (row) {
      var dt = document.createElement('dt');
      var dd = document.createElement('dd');
      dt.textContent = row[0];
      dd.textContent = row[1];
      dl.appendChild(dt);
      dl.appendChild(dd);
    });
  }

  function tick() { now.textContent = Math.floor(Date.now() / 1000); }
  tick();
  setInterval(tick, 1000);

  epoch.addEventListener('input', function () {
    var text = epoch.value.trim();
    if (!text) return render(epochResult, []);
    if (!/^-?\d+$/.test(text)) return render(epochResult, [['Erro', 'Digite só números.']]);
    // 13 dígitos ou mais é milissegundo: 10^11 segundos só chega no ano 5138
    var ms = Math.abs(+text) >= 1e11;
    var d = new Date(ms ? +text : +text * 1000);
    if (isNaN(d)) return render(epochResult, [['Erro', 'Timestamp fora do intervalo de datas.']]);
    render(epochResult, [
      ['Unidade', ms ? 'milissegundos' : 'segundos'],
      ['Seu fuso (' + timezone + ')', local.format(d)],
      ['UTC', utc.format(d)],
      ['ISO 8601', d.toISOString()]
    ]);
  });

  date.addEventListener('input', function () {
    if (!date.value) return render(dateResult, []);
    // datetime-local não tem fuso: new Date() interpreta no fuso do navegador
    var d = new Date(date.value);
    if (isNaN(d)) return render(dateResult, []);
    render(dateResult, [
      ['Segundos', String(Math.floor(d.getTime() / 1000))],
      ['Milissegundos', String(d.getTime())],
      ['UTC', utc.format(d)],
      ['ISO 8601', d.toISOString()]
    ]);
  });
})();
