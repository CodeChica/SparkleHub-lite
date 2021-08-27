fetch("/sparkles.html")
  .then((response) => response.text())
  .then((html) => {
    document.querySelector('#sparkles-list').innerHTML = html
  });
