// Formatador de JSON: formata ou minifica no lugar e mostra onde está o erro
(function () {
  var form = document.querySelector('.json-tool');
  if (!form) return;
  var textarea = form.json;
  var status = document.getElementById('json-status');

  function setStatus(message, valid) {
    status.textContent = message;
    if (valid === undefined) textarea.removeAttribute('aria-invalid');
    else textarea.setAttribute('aria-invalid', valid ? 'false' : 'true');
  }

  // O JSON.parse nem sempre diz onde está o erro (o Chrome omite a posição em
  // casos comuns, como vírgula sobrando). Este validador mínimo só serve para
  // achar a posição do primeiro erro.
  function errorPosition(text) {
    var i = 0;
    function fail() { throw i; }
    function ws() { while (/[ \t\n\r]/.test(text[i] || '')) i++; }
    function literal(word) { if (text.slice(i, i + word.length) !== word) fail(); i += word.length; }
    function string() {
      i++;
      while (text[i] !== '"') {
        if (i >= text.length || text.charCodeAt(i) < 0x20) fail();
        if (text[i] === '\\') {
          i++;
          if (text[i] === 'u') { if (!/^[0-9a-fA-F]{4}$/.test(text.slice(i + 1, i + 5))) fail(); i += 4; }
          else if ('"\\/bfnrt'.indexOf(text[i]) === -1 || i >= text.length) fail();
        }
        i++;
      }
      i++;
    }
    function number() {
      var m = text.slice(i).match(/^-?(0|[1-9]\d*)(\.\d+)?([eE][+-]?\d+)?/);
      if (!m) fail();
      i += m[0].length;
    }
    function list(close, item) {
      i++; ws();
      if (text[i] === close) { i++; return; }
      for (;;) {
        item(); ws();
        if (text[i] === close) { i++; return; }
        if (text[i] !== ',') fail();
        i++; ws();
      }
    }
    function value() {
      ws();
      var c = text[i];
      if (c === '{') list('}', function () { if (text[i] !== '"') fail(); string(); ws(); if (text[i] !== ':') fail(); i++; value(); });
      else if (c === '[') list(']', value);
      else if (c === '"') string();
      else if (c === 't') literal('true');
      else if (c === 'f') literal('false');
      else if (c === 'n') literal('null');
      else number();
    }
    try { value(); ws(); if (i < text.length) fail(); } catch (pos) { return typeof pos === 'number' ? pos : null; }
    return null;
  }

  function errorLocation(text) {
    var position = errorPosition(text);
    if (position === null) return null;
    var before = text.slice(0, position).split('\n');
    return { line: before.length, column: before[before.length - 1].length + 1 };
  }

  form.addEventListener('submit', function (event) {
    event.preventDefault();
    var text = textarea.value;
    if (!text.trim()) return setStatus('Cole um JSON para começar.');

    var data;
    try {
      data = JSON.parse(text);
    } catch (e) {
      var where = errorLocation(text);
      return setStatus('JSON inválido' + (where ? ' na linha ' + where.line + ', coluna ' + where.column : '') + ': ' + e.message, false);
    }

    var indent = form.indent.value === 'tab' ? '\t' : +form.indent.value;
    var minify = event.submitter && event.submitter.value === 'minify';
    textarea.value = minify ? JSON.stringify(data) : JSON.stringify(data, null, indent);
    setStatus(minify ? 'JSON válido e minificado.' : 'JSON válido e formatado.', true);
  });

  textarea.addEventListener('input', function () { setStatus(''); });
})();
