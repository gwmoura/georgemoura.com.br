// Gerador de UUID v4 e v7 (RFC 9562) com crypto.getRandomValues
(function () {
  var form = document.querySelector('.uuid-generator');
  if (!form || !window.crypto || !crypto.getRandomValues) return;
  var output = document.getElementById('uuids');
  var copy = form.querySelector('[data-copy]');

  function hex(bytes) {
    var s = Array.prototype.map.call(bytes, function (b) { return (b + 0x100).toString(16).slice(1); }).join('');
    return s.slice(0, 8) + '-' + s.slice(8, 12) + '-' + s.slice(12, 16) + '-' + s.slice(16, 20) + '-' + s.slice(20);
  }

  function uuid(version) {
    var bytes = new Uint8Array(16);
    crypto.getRandomValues(bytes);
    if (version === 7) {
      // Os 48 primeiros bits são o horário Unix em milissegundos (big-endian)
      var ts = Date.now();
      for (var i = 5; i >= 0; i--) {
        bytes[i] = ts % 256;
        ts = Math.floor(ts / 256);
      }
    }
    bytes[6] = (bytes[6] & 0x0f) | (version << 4); // versão
    bytes[8] = (bytes[8] & 0x3f) | 0x80; // variante RFC 9562
    return hex(bytes);
  }

  form.addEventListener('submit', function (event) {
    event.preventDefault();
    var quantity = Math.min(Math.max(parseInt(form.quantity.value, 10) || 1, 1), 500);
    var version = +form.version.value;
    var list = [];
    for (var i = 0; i < quantity; i++) {
      var id = uuid(version);
      if (form.nohyphens.checked) id = id.replace(/-/g, '');
      if (form.uppercase.checked) id = id.toUpperCase();
      list.push(id);
    }
    output.value = list.join('\n');
    output.rows = Math.min(Math.max(quantity, 2), 12);
    output.hidden = false;
    if (copy && navigator.clipboard) copy.hidden = false;
  });
})();
