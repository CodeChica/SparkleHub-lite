document.addEventListener('DOMContentLoaded', (event) => {
  window.app = new Vue({
    el: '#app',
    data: {
      intervalId: null,
      sparkle: "",
      sparkles: []
    },
    created: function() {
      this.reload();
      this.intervalId = setInterval(() => this.reload(), 5000);
    },
    destroyed: function() {
      if (this.intervalId)
        clearInterval(this.intervalId);
      this.intervalId = null;
    },
    computed: {
      heading: function() {
        return this.sparkles.length == 0 ? "No Sparkles Yet" : "Recent Sparkles";
      },
      recentSparkles: function() {
        return this.sparkles.reverse();
      }
    },
    methods: {
      reload: function() {
        fetch("/sparkles.json")
          .then((response) => response.json())
          .then((data) => this.sparkles = data.sparkles);
      },
      submitSparkle: function() {
        fetch("/sparkles.json", {
          method: "POST",
          mode: "cors",
          cache: "no-cache",
          headers: { "Content-Type": "application/json" },
          redirect: "follow",
          body: JSON.stringify({ body: this.sparkle })
        })
        .then((response) => response.json())
        .then((json) => {
          this.sparkles.push(json);
          this.sparkle = "";
        })
      }
    }
  })
});
