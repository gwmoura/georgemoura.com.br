// Gerador e validador de CPF e CNPJ. O CNPJ alfanumérico (2026) usa o mesmo
// cálculo: cada caractere vale o código ASCII menos 48 ("0" = 0, "A" = 17).
(function () {
  var generator = document.querySelector('.doc-generator');
  var validator = document.querySelector('.doc-validator');
  if (!generator || !validator) return;

  function value(c) { return c.charCodeAt(0) - 48; }

  // Dígito verificador do módulo 11, igual para CPF e CNPJ
  function checkDigit(base, weights) {
    var sum = 0;
    for (var i = 0; i < base.length; i++) sum += value(base[i]) * weights[i];
    var rest = sum % 11;
    return String(rest < 2 ? 0 : 11 - rest);
  }

  var cpfWeights = [11, 10, 9, 8, 7, 6, 5, 4, 3, 2];
  var cnpjWeights = [6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2];

  function cpfDigits(base) {
    var d1 = checkDigit(base, cpfWeights.slice(1));
    return d1 + checkDigit(base + d1, cpfWeights);
  }

  function cnpjDigits(base) {
    var d1 = checkDigit(base, cnpjWeights.slice(1));
    return d1 + checkDigit(base + d1, cnpjWeights);
  }

  function repeated(doc) { return /^(.)\1+$/.test(doc); }

  function isCPF(doc) {
    return /^\d{11}$/.test(doc) && !repeated(doc) && cpfDigits(doc.slice(0, 9)) === doc.slice(9);
  }

  function isCNPJ(doc) {
    return /^[0-9A-Z]{12}\d{2}$/.test(doc) && !repeated(doc) && cnpjDigits(doc.slice(0, 12)) === doc.slice(12);
  }

  function maskCPF(doc) { return doc.replace(/^(.{3})(.{3})(.{3})(.{2})$/, '$1.$2.$3-$4'); }
  function maskCNPJ(doc) { return doc.replace(/^(.{2})(.{3})(.{3})(.{4})(.{2})$/, '$1.$2.$3/$4-$5'); }

  function randomDigits(n) {
    var buf = new Uint8Array(n);
    crypto.getRandomValues(buf);
    return Array.prototype.map.call(buf, function (b) { return b % 10; }).join('');
  }

  var output = document.getElementById('documento-gerado');
  var copy = document.querySelector('[data-copy="#documento-gerado"]');

  generator.addEventListener('submit', function (event) {
    event.preventDefault();
    var cnpj = event.submitter && event.submitter.value === 'cnpj';
    var doc;
    do {
      // CNPJ: 8 dígitos de raiz + 0001 (matriz)
      var base = cnpj ? randomDigits(8) + '0001' : randomDigits(9);
      doc = base + (cnpj ? cnpjDigits(base) : cpfDigits(base));
    } while (repeated(doc));
    var masked = generator.mask.checked;
    output.textContent = cnpj ? (masked ? maskCNPJ(doc) : doc) : (masked ? maskCPF(doc) : doc);
    if (copy && navigator.clipboard) copy.hidden = false;
  });

  var input = validator.documento;
  var status = document.getElementById('documento-status');

  function validate() {
    var doc = input.value.toUpperCase().replace(/[.\/\-\s]/g, '');
    if (!doc) {
      input.removeAttribute('aria-invalid');
      status.textContent = 'Digite um documento, com ou sem pontuação.';
      return;
    }
    var message;
    var valid = false;
    if (doc.length === 11 && /^\d+$/.test(doc)) {
      valid = isCPF(doc);
      message = valid ? 'CPF válido: ' + maskCPF(doc) : 'CPF inválido.';
    } else if (doc.length === 14) {
      valid = isCNPJ(doc);
      message = valid ? 'CNPJ válido: ' + maskCNPJ(doc) : 'CNPJ inválido.';
    } else {
      message = 'Um CPF tem 11 dígitos e um CNPJ tem 14 caracteres (' + doc.length + ' digitados).';
    }
    input.setAttribute('aria-invalid', valid ? 'false' : 'true');
    status.textContent = message;
  }

  validator.addEventListener('submit', function (event) { event.preventDefault(); validate(); });
  input.addEventListener('input', validate);
})();
