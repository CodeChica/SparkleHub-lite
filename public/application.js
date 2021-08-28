class Sparkles {
  constructor(id, url) {
    this.element = document.querySelector(id);
    this.url = url;
  }

  refresh() {
    fetch(this.url)
      .then((response) => response.text())
      .then((html) => this.element.innerHTML = html);
  }

  startPolling() {
    if (!this.intervalId) {
      this.intervalId = setInterval(() => this.refresh(), 1000);
    }
  }

  stopPolling() {
    if (this.intervalId) {
      clearInterval(this.intervalId);
    }
  }
}

document.addEventListener('DOMContentLoaded', (event) => {
  window.Sparkles = new Sparkles('#sparkles-list', "/sparkles.html")
  window.Sparkles.startPolling();
});

document.addEventListener('unload', (event) => {
  window.Sparkles.stopPolling();
});
