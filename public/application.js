class AutoReload {
  constructor(id, url) {
    this.element = document.querySelector(id);
    this.url = url;
  }

  start(milliseconds = 1000) {
    if (!this.intervalId) {
      this.intervalId = setInterval(() => this.reload(), milliseconds);
    }
  }

  reload() {
    fetch(this.url)
      .then((response) => response.text())
      .then((html) => this.element.innerHTML = html);
  }

  stop() {
    if (this.intervalId) {
      clearInterval(this.intervalId);
      this.intervalId = null;
    }
  }
}

document.addEventListener('DOMContentLoaded', (event) => {
  window.sparkles = new AutoReload('#sparkles-list', "/sparkles.html")
  window.sparkles.start();
});

document.addEventListener('unload', (event) => {
  window.sparkles.stop();
});

