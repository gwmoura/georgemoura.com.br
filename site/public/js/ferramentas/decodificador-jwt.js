// Decodificador de JWT: só decodifica header e payload (base64url), sem
// verificar a assinatura. Nada sai do navegador.
(function () {
  var input = document.getElementById('jwt');
  if (!input) return;
  var status = document.getElementById('jwt-status');
  var parts = document.querySelector('.jwt-parts');
  var header = document.getElementById('jwt-header');
  var payload = document.getElementById('jwt-payload');
  var dates = document.getElementById('jwt-datas');
  var format = new Intl.DateTimeFormat('pt-BR', { dateStyle: 'short', timeStyle: 'medium' });
  var claims = { iat: 'Emitido em (iat)', nbf: 'Válido a partir de (nbf)', exp: 'Expira em (exp)' };

  function decode(part) {
    var base64 = part.replace(/-/g, '+').replace(/_/g, '/');
    base64 += '===='.slice(0, (4 - base64.length % 4) % 4);
    var binary = atob(base64);
    var bytes = new Uint8Array(binary.length);
    for (var i = 0; i < binary.length; i++) bytes[i] = binary.charCodeAt(i);
    return JSON.parse(new TextDecoder().decode(bytes));
  }

  function fail(message) {
    parts.hidden = true;
    status.textContent = message;
    input.setAttribute('aria-invalid', 'true');
  }

  function renderDates(data) {
    dates.textContent = '';
    var now = Date.now() / 1000;
    Object.keys(claims).forEach(function (claim) {
      if (typeof data[claim] !== 'number') return;
      var dt = document.createElement('dt');
      var dd = document.createElement('dd');
      dt.textContent = claims[claim];
      dd.textContent = format.format(new Date(data[claim] * 1000));
      if (claim === 'exp') dd.textContent += data.exp < now ? ' (expirado)' : ' (ainda válido)';
      if (claim === 'nbf' && data.nbf > now) dd.textContent += ' (ainda não vale)';
      dates.appendChild(dt);
      dates.appendChild(dd);
    });
  }

  input.addEventListener('input', function () {
    var token = input.value.trim().replace(/^Bearer\s+/i, '');
    if (!token) {
      parts.hidden = true;
      status.textContent = '';
      input.removeAttribute('aria-invalid');
      return;
    }
    var pieces = token.split('.');
    if (pieces.length !== 3) return fail('Um JWT tem 3 partes separadas por ponto (header.payload.assinatura); este tem ' + pieces.length + '.');
    var h, p;
    try { h = decode(pieces[0]); } catch (e) { return fail('O header não é um base64url com JSON válido.'); }
    try { p = decode(pieces[1]); } catch (e) { return fail('O payload não é um base64url com JSON válido.'); }

    header.textContent = JSON.stringify(h, null, 2);
    payload.textContent = JSON.stringify(p, null, 2);
    renderDates(p);
    parts.hidden = false;
    input.setAttribute('aria-invalid', 'false');
    status.textContent = 'Token decodificado' + (h.alg ? ' (algoritmo ' + h.alg + ')' : '') + '. A assinatura não foi verificada.';
  });
})();
