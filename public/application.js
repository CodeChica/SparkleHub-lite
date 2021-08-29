class ContentLoaderController {
  constructor(elementId, url, refreshInterval = 10000) {
    this.element = document.querySelector(elementId);
    this.url = url;
    this.start(refreshInterval);
  }

  start(refreshInterval) {
    if (this.hasStarted())
      return

    this.intervalId = setInterval(() => this.reload(), refreshInterval);
  }

  hasStarted() {
    return this.intervalId;
  }

  reload() {
    fetch(this.url)
      .then((response) => response.text())
      .then((html) => this.element.innerHTML = html);
  }

  stop() {
    if (!this.hasStarted())
      return

    clearInterval(this.intervalId);
    this.intervalId = null;
  }
}

document.addEventListener('DOMContentLoaded', (event) => {
  window.sparkles = new ContentLoaderController('#sparkles-list', "/sparkles.html")
});

document.addEventListener('unload', (event) => {
  window.sparkles.stop();
});
